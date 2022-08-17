package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"log"
	"time"
	"web3game/database"
	"web3game/database/inRamDb"
	"web3game/games/rockPaperScissors"
	"web3game/helper"
	"web3game/models/entity"
	"web3game/models/events/input"
	"web3game/models/events/output"
	"web3game/models/request"
	"web3game/models/response"
	"web3game/test"
)

var (
	socketIsConnected = false
	server            *socketio.Server
)

func ServeSocket(group *gin.RouterGroup, router *gin.Engine) *socketio.Server {

	server = socketio.NewServer(nil)

	socketServerMainRoutes(group, server)

	router.LoadHTMLFiles("index.html")
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	go server.Serve()

	return server
}

func socketServerMainRoutes(group *gin.RouterGroup, server *socketio.Server) {

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		socketIsConnected = true

		rockPaperGroup := group.Group("/rockpaper")

		rockPaperGroup.GET("/getAllRockPaper", func(c *gin.Context) {
			var result entity.Games
			query := database.GetTableGames().Find(&result)
			if query.Error != nil {
				helper.SendErrorResponse(c, query.Error.Error())
			}
			helper.SendSuccessResponse(c, result, nil)
		})
		rockPaperGroup.POST("/createGame", func(c *gin.Context) {
			var err error

			//check request params
			var rq request.CreateGameRq
			err = c.ShouldBindJSON(&rq)

			if err != nil {
				err = errors.New(helper.REQURED_PARAMETERS_IS_NOT_SET)
				helper.SendErrorResponse(c, err.Error())
				return
			}

			//check player transaction
			tx, isPending, err := test.EthClient.TransactionByHash(context.Background(), common.HexToHash(rq.TxHash))
			if isPending {
				err = errors.New(helper.TRANSACTION_IS_PENDING)
				helper.SendErrorResponse(c, err.Error())
				return
			}
			if tx.Hash().String() != rq.TxHash && tx.Value().String() != rq.Value {
				err = errors.New(helper.TRANSACTION_IS_NOT_VALID)
				helper.SendErrorResponse(c, err.Error())
				return
			}
			//create Game in database
			var gameEntity = entity.Game{}
			gameEntity.GameId = "1" /*guuid.NewString()*/
			gameEntity.Players = rq.PlayerAddress
			gameEntity.Status = 1
			query := database.GetTableGames().Create(&gameEntity)
			err = query.Error

			if err != nil {
				helper.SendErrorResponse(c, err.Error())
			} else {
				s.Emit("gameEvent", output.GameEvent{
					GameId:          gameEntity.GameId,
					PlayerAddresses: gameEntity.Players,
					Status:          1,
				})
				s.Join(gameEntity.GameId)
				helper.SendSuccessResponse(c, response.CreateGameRp{Success: true}, nil)
			}

		})
		rockPaperGroup.POST("/joinGame", func(c *gin.Context) {
			var err error

			//check request params
			var rq request.JoinGameRq
			err = c.ShouldBindJSON(&rq)

			if err != nil {
				err = errors.New(helper.REQURED_PARAMETERS_IS_NOT_SET)
				helper.SendErrorResponse(c, err.Error())
				return
			}

			//check player transaction
			tx, isPending, err := test.EthClient.TransactionByHash(context.Background(), common.HexToHash(rq.TxHash))
			if isPending {
				err = errors.New(helper.TRANSACTION_IS_PENDING)
				helper.SendErrorResponse(c, err.Error())
				return
			}
			if tx.Hash().String() != rq.TxHash && tx.Value().String() != rq.Value {
				err = errors.New(helper.TRANSACTION_IS_NOT_VALID)
				helper.SendErrorResponse(c, err.Error())
				return
			}

			var result *entity.Game
			database.GetTableGames().Where("gameId = LIKE ?", rq.GameId).Find(&result)

			query := database.GetTableGames().
				Where("game_id LIKE ?", rq.GameId).
				Update("players", result.Players+","+rq.PlayerAddress).
				Update("status", 2)
			err = query.Error

			if err != nil {
				helper.SendErrorResponse(c, err.Error())
			} else {
				s.Emit("gameEvent", output.GameEvent{
					GameId:          rq.GameId,
					PlayerAddresses: result.Players + "," + rq.PlayerAddress,
					Status:          2,
				})
				s.Join(rq.GameId)
				helper.SendSuccessResponse(c, response.JoinGameRp{Success: true}, nil)
			}

		})

		return nil
	})

	server.OnEvent("/", "echo", func(s socketio.Conn, msg string) {
		print(msg)
	})
	server.OnEvent("/", "gameEmitter", func(s socketio.Conn, msg string) {
		var err error

		//check request params
		var rq input.GameActionEvent
		err = json.Unmarshal([]byte(msg), &rq)
		if err != nil {
			err = errors.New(helper.REQURED_PARAMETERS_IS_NOT_SET)
			return
		}

		allRounds, err := inRamDb.ReadAllRockPaperRounds()
		thisRound := len(allRounds) / 2

		if thisRound < 3 {
			//add player action to game actions

			if allRounds != nil {
				//player try to send multiple request we prevent it
				if allRounds[len(allRounds)-1].PlayerAddress == rq.PlayerAddress {
					if len(allRounds)%2 != 0 {
						println("player try to send multiple request we prevent it")
						return
					}
				}
			}

			err = inRamDb.InsertRockPaperAction(&entity.RockPaperAction{
				Id:            time.Now().String(),
				GameId:        rq.GameId,
				PlayerAddress: rq.PlayerAddress,
				Action:        rq.Action,
			})

			println(rq.PlayerAddress+" action: ", rq.Action)

			//reinit all round because in ram db changed
			allRounds, err = inRamDb.ReadAllRockPaperRounds()

			thisRound := len(allRounds) / 2
			if thisRound == 3 {
				//this round is last round of this game, and game should end and send to all result
				player1Address := allRounds[0].PlayerAddress
				player1Wins := 0
				player2Address := allRounds[1].PlayerAddress
				player2Wins := 0
				count := 0
				for range allRounds {
					if count+2 <= len(allRounds) {
						winner := rockPaperScissors.CalculaterRockPaperRoundWinner(
							allRounds[count], allRounds[count+1],
						)
						if winner == player1Address {
							player1Wins++
						} else {
							player2Wins++
						}
					} else {
						if player1Wins > player2Wins {
							server.BroadcastToRoom("/", rq.GameId, "gameEndResult", output.GameEndResultEvent{Winner: player1Address})
							println("victory for: " + player1Address)
						} else {
							server.BroadcastToRoom("/", rq.GameId, "gameEndResult", output.GameEndResultEvent{Winner: player2Address})
							println("victory for: " + player2Address)
						}
					}
					count++
				}
			} else {
				if len(allRounds)%2 == 0 && len(allRounds) > 0 {
					//should send to all this round result

					winner := rockPaperScissors.CalculaterRockPaperRoundWinner(
						allRounds[len(allRounds)-1], allRounds[len(allRounds)-2],
					)
					// if game come to tie should delete last round
					if winner == "Tie" {
						err = inRamDb.DeleteRockPaper(allRounds[len(allRounds)-1])
						err = inRamDb.DeleteRockPaper(allRounds[len(allRounds)-2])
					}
					println("winner is: " + winner)
					s.Emit(
						"gameEndResult",
						output.GameRoundResultEvent{
							Winner: winner,
						},
					)
				} else {
					//waiting for other player to send round action
					println("waiting for other player")
				}
			}
		} else {
			//player try to send ended game
			println("player try to send ended game")
		}
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
		socketIsConnected = false
		server.Close()
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
		log.Println("count", server.Count())
		socketIsConnected = false
		server.Close()
	})

}
