package centrifugo

//import (
//	"context"
//	"encoding/json"
//	"errors"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/gin-gonic/gin"
//	socketio "github.com/googollee/go-socket.io"
//	"log"
//	"time"
//	"web3game/database"
//	"web3game/database/inRamDb"
//	"web3game/games/rockPaperScissors"
//	"web3game/helper"
//	"web3game/models/entity"
//	"web3game/models/events/input"
//	"web3game/models/events/output"
//	"web3game/models/request"
//	"web3game/test"
//)
//
//var (
//	socketIsConnected = false
//	SocketServer      *socketio.Server
//)
//
//func ServeSocket(router *gin.Engine) *socketio.Server {
//
//	SocketServer = socketio.NewServer(nil)
//
//	socketServerMainRoutes(SocketServer)
//
//	router.LoadHTMLFiles("index.html")
//	router.GET("/socket.io/*any", gin.WrapH(SocketServer))
//	router.POST("/socket.io/*any", gin.WrapH(SocketServer))
//	go SocketServer.Serve()
//
//	return SocketServer
//}
//
//func socketServerMainRoutes(server *socketio.Server) {
//
//	server.OnConnect("/", func(s socketio.Conn) error {
//		s.SetContext("")
//		log.Println("connected:", s.ID())
//		socketIsConnected = true
//		return nil
//	})
//
//	server.OnEvent("/", "createGame", func(s socketio.Conn, msg string) {
//		var err error
//
//		//check request params
//		var rq request.CreateGameRq
//		err = json.Unmarshal([]byte(msg), &rq)
//
//		if err != nil {
//			err = errors.New(helper.REQURED_PARAMETERS_IS_NOT_SET)
//			return
//		}
//
//		//check player transaction
//		tx, isPending, err := test.EthClient.TransactionByHash(context.Background(), common.HexToHash(rq.TxHash))
//		if isPending {
//			err = errors.New(helper.TRANSACTION_IS_PENDING)
//			return
//		}
//		if tx.Hash().String() != rq.TxHash && tx.Value().String() != rq.Value {
//			err = errors.New(helper.TRANSACTION_IS_NOT_VALID)
//			return
//		}
//		//create Game in database
//		var gameEntity = entity.Game{}
//		gameEntity.GameId = "1" /*guuid.NewString()*/
//		gameEntity.Players = rq.PlayerAddress
//		gameEntity.Status = 1
//		query := database.GetTableGames().Create(&gameEntity)
//		err = query.Error
//
//		if err == nil {
//			s.Emit("gameEvent", output.GameEvent{
//				GameId:          gameEntity.GameId,
//				PlayerAddresses: gameEntity.Players,
//				Status:          1,
//			})
//			s.Join(gameEntity.GameId)
//		}
//
//	})
//	server.OnEvent("/", "joinGame", func(s socketio.Conn, msg string) {
//		var err error
//
//		//check request params
//		var rq request.JoinGameRq
//		err = json.Unmarshal([]byte(msg), &rq)
//
//		if err != nil {
//			err = errors.New(helper.REQURED_PARAMETERS_IS_NOT_SET)
//			return
//		}
//
//		//check player transaction
//		tx, isPending, err := test.EthClient.TransactionByHash(context.Background(), common.HexToHash(rq.TxHash))
//		if isPending {
//			err = errors.New(helper.TRANSACTION_IS_PENDING)
//			return
//		}
//		if tx.Hash().String() != rq.TxHash && tx.Value().String() != rq.Value {
//			err = errors.New(helper.TRANSACTION_IS_NOT_VALID)
//			return
//		}
//
//		var result *entity.Game
//		database.GetTableGames().Where("gameId = LIKE ?", rq.GameId).Find(&result)
//
//		query := database.GetTableGames().
//			Where("game_id LIKE ?", rq.GameId).
//			Update("players", result.Players+","+rq.PlayerAddress).
//			Update("status", 2)
//		err = query.Error
//
//		if err == nil {
//			s.Emit("gameEvent", output.GameEvent{
//				GameId:          rq.GameId,
//				PlayerAddresses: result.Players + "," + rq.PlayerAddress,
//				Status:          2,
//			})
//			s.Join(rq.GameId)
//		}
//
//	})
//
//	server.OnEvent("/", "gameEmitter", func(s socketio.Conn, msg string) {
//		var err error
//
//		//check request params
//		var rq input.GameActionEvent
//		err = json.Unmarshal([]byte(msg), &rq)
//		if err != nil {
//			err = errors.New(helper.REQURED_PARAMETERS_IS_NOT_SET)
//			return
//		}
//
//		allRounds, err := inRamDb.ReadAllRockPaperRounds()
//		thisRound := len(allRounds) / 2
//
//		if thisRound < 3 {
//			//add player action to game actions
//
//			if allRounds != nil {
//				//player try to send multiple request we prevent it
//				if allRounds[len(allRounds)-1].PlayerAddress == rq.PlayerAddress {
//					if len(allRounds)%2 != 0 {
//						println("player try to send multiple request we prevent it")
//						return
//					}
//				}
//			}
//
//			err = inRamDb.InsertRockPaperAction(&entity.RockPaperAction{
//				Id:            time.Now().String(),
//				GameId:        rq.GameId,
//				PlayerAddress: rq.PlayerAddress,
//				Action:        rq.Action,
//			})
//
//			println(rq.PlayerAddress+" action: ", rq.Action)
//
//			//reinit all round because in ram db changed
//			allRounds, err = inRamDb.ReadAllRockPaperRounds()
//
//			thisRound := len(allRounds) / 2
//			if thisRound == 3 {
//				//this round is last round of this game, and game should end and send to all result
//				player1Address := allRounds[0].PlayerAddress
//				player1Wins := 0
//				player2Address := allRounds[1].PlayerAddress
//				player2Wins := 0
//				count := 0
//				for range allRounds {
//					if count+2 <= len(allRounds) {
//						winner := rockPaperScissors.CalculaterRockPaperRoundWinner(
//							allRounds[count], allRounds[count+1],
//						)
//						if winner == player1Address {
//							player1Wins++
//						} else {
//							player2Wins++
//						}
//					} else {
//						if player1Wins > player2Wins {
//							server.BroadcastToRoom("/", rq.GameId, "gameEndResult", output.GameEndResultEvent{Winner: player1Address})
//							println("victory for: " + player1Address)
//						} else {
//							server.BroadcastToRoom("/", rq.GameId, "gameEndResult", output.GameEndResultEvent{Winner: player2Address})
//							println("victory for: " + player2Address)
//						}
//					}
//					count++
//				}
//			} else {
//				if len(allRounds)%2 == 0 && len(allRounds) > 0 {
//					//should send to all this round result
//
//					winner := rockPaperScissors.CalculaterRockPaperRoundWinner(
//						allRounds[len(allRounds)-1], allRounds[len(allRounds)-2],
//					)
//					// if game come to tie should delete last round
//					if winner == "Tie" {
//						err = inRamDb.DeleteRockPaper(allRounds[len(allRounds)-1])
//						err = inRamDb.DeleteRockPaper(allRounds[len(allRounds)-2])
//					}
//					println("winner is: " + winner)
//					s.Emit(
//						"gameRoundResult",
//						output.GameRoundResultEvent{
//							Winner: winner,
//						},
//					)
//				} else {
//					//waiting for other player to send round action
//					println("waiting for other player")
//				}
//			}
//		} else {
//			//player try to send ended game
//			println("player try to send ended game")
//		}
//	})
//
//	server.OnError("/", func(s socketio.Conn, e error) {
//		log.Println("meet error:", e)
//		socketIsConnected = false
//	})
//
//	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
//		log.Println("closed", reason)
//		log.Println("count", server.Count())
//		socketIsConnected = false
//	})
//
//	server.OnEvent("/", "join", func(s socketio.Conn, msg string) {
//		println("JOIN")
//		s.Join("1")
//		println(s.Rooms())
//		s.Emit("1", "player joined 1")
//		server.BroadcastToRoom("/", "1", "1", "player joined 2")
//	})
//
//}
