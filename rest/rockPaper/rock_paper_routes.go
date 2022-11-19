package rockPaper

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"web3game/centrifugo/realtime_api"
	"web3game/database"
	"web3game/games/rockPaperScissors"
	"web3game/helper"
	"web3game/models/entity"
	"web3game/models/request"
	"web3game/models/response"
	"web3game/test/web3"
)

func AddRockPaperRoutes(group *gin.RouterGroup) {
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
		tx, isPending, err := web3.EthClient.TransactionByHash(context.Background(), common.HexToHash(rq.TxHash))
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
			//if everything is done publish game to rock paper channel
			if (realtime_api.PublishRockPaperGame(&response.GameEvent{
				GameId:          gameEntity.GameId,
				PlayerAddresses: gameEntity.Players,
				Status:          1,
			})) {
				helper.SendSuccessResponse(c, response.CreateGameRp{GameId: gameEntity.GameId}, nil)
			} else {
				helper.SendErrorResponse(c, helper.UNKNOWN_ERROR)
			}
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
		tx, isPending, err := web3.EthClient.TransactionByHash(context.Background(), common.HexToHash(rq.TxHash))
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
			realtime_api.PublishRockPaperGame(&response.GameEvent{
				GameId:          rq.GameId,
				PlayerAddresses: result.Players + "," + rq.PlayerAddress,
				Status:          2,
			})
			helper.SendSuccessResponse(c, response.JoinGameRp{GameId: rq.GameId}, nil)
		}

	})

	rockPaperGroup.POST("/open_reward", func(c *gin.Context) {
		var err error

		//check request params
		var rq request.OpenRewardRq
		err = c.ShouldBindJSON(&rq)

		if err != nil {
			err = errors.New(helper.REQURED_PARAMETERS_IS_NOT_SET)
			helper.SendErrorResponse(c, err.Error())
			return
		}
		rockPaperHistoryRp := realtime_api.GetRockPaperHistory(&rq)
		if rockPaperScissors.CalculatorRockPaperGameWinner(rockPaperHistoryRp) == rq.PlayerAddress {
			helper.SendSuccessResponse(c, response.OpenRewardRp{
				Winner: rockPaperScissors.CalculatorRockPaperGameWinner(rockPaperHistoryRp),
			}, nil)
		} else {
			helper.SendErrorResponse(c, helper.UNKNOWN_ERROR)
		}

	})
}
