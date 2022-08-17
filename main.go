package main

import (
	"crypto/rand"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"math/big"
	"os"
	helper2 "web3game/contracts"
	"web3game/database"
	"web3game/database/inRamDb"
	"web3game/helper"
	"web3game/models/entity"
	"web3game/models/request"
	"web3game/models/response"
	"web3game/test"
)

var (
	router = gin.Default()
)

func terminateCode() {
	e := os.Remove("/home/khoujani/keys.json")
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	defer terminateCode()

	test.Fire()
	database.GetDB()
	inRamDb.GetInRamDb()

	_ = router.SetTrustedProxies([]string{"localhost,192.168.0.1"})
	v1 := router.Group("/v1")

	AddLoginRoutes(v1)

	defer func(socket *socketio.Server) {
		err := socket.Close()
		if err != nil {
			print("Socket close error: ", err)
		}
	}(ServeSocket(v1, router))

	_ = router.Run(":5000")

}

func AddLoginRoutes(group *gin.RouterGroup) {

	//localhost:5000/v1/auth/requestNonce
	authGroup := group.Group("/auth")

	authGroup.POST("/requestNonce", func(c *gin.Context) {
		var err error

		//validate required parameters
		var rq request.RequestNonceRq
		if err = c.ShouldBindJSON(&rq); err != nil {
			helper.SendErrorResponse(c, helper.REQURED_PARAMETERS_IS_NOT_SET)
			return
		}

		//generate nonce
		max := new(big.Int)
		max.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(max, big.NewInt(1))
		n, err := rand.Int(rand.Reader, max)
		nonce := n.Text(32)

		//save nonce to database for requested wallet address
		var nonceEntity = entity.Nonce{}
		nonceEntity.Nonce = helper.LoginToWeb3 + nonce
		nonceEntity.WalletAddress = rq.WalletAddress
		query := database.GetTableNonce().Create(&nonceEntity)
		err = query.Error

		if err != nil {
			helper.SendErrorResponse(c, err.Error())
		} else {
			helper.SendSuccessResponse(c, response.RequestNonceRp{Nonce: helper.LoginToWeb3 + nonce}, nil)
		}

	})

	authGroup.POST("/validateSign", func(c *gin.Context) {
		var err error

		//validate required parameters
		var rq request.ValidateSignRq
		if err = c.ShouldBindJSON(&rq); err != nil {
			helper.SendErrorResponse(c, helper.REQURED_PARAMETERS_IS_NOT_SET)
		}

		//get nonce from database for requested wallet address
		var nonceEntity = entity.Nonce{}
		query := database.GetTableNonce().Where("wallet_address = ?", rq.WalletAddress).First(&nonceEntity)
		err = query.Error

		//delete nonce from database
		query = database.GetTableNonce().Where("wallet_address = ?", rq.WalletAddress).Delete(&nonceEntity)

		if helper2.DecodePersonalSign(nonceEntity.Nonce, rq.Sign, rq.WalletAddress) {
			helper.SendSuccessResponse(c, response.ValidateSignRp{Success: true}, nil)
		} else {
			helper.SendErrorResponse(c, "Invalid signature")
		}

	})

}
