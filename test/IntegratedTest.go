package test

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"time"
	"web3game/contracts"
	"web3game/contracts/model"
	"web3game/database"
	"web3game/models/events/input"
	"web3game/models/request"
)

func Fire() {

	// delete existed test record
	database.GetTableGames().Where("game_id = ?", 1).Delete(nil)

	// provide ganache accounts "addresses" / "private keys"
	accounts := ProvideGanacheAccounts()

	// deploy contract to `localhost:8545` network with a private key
	gamePool := DeployContract(accounts.PrivateKeys[0])

	println("\n")
	time.Sleep(1 * time.Second)
	GenesisPlayerTest(gamePool, &accounts)

	println("\n")
	time.Sleep(1 * time.Second)
	OtherPlayerTest(gamePool, &accounts)

	//println("\n")
	//time.Sleep(1 * time.Second)
	//WithdrawToWinTest(gamePool, &accounts)

	println("\n")
	time.Sleep(1 * time.Second)
	ReturnAGameTest(gamePool, &accounts)

	println("\n")
}

func GenesisPlayerTest(gamePool *contracts.GamePool, accounts *model.GanacheAccounts) {
	// player 1 deposit
	accountPlayer1 := accounts.PrivateKeys[1]
	depositValue := big.NewInt(10000000000)
	accountBalance, err := EthClient.BalanceAt(context.Background(), GetAddressFromPrivateKey(accountPlayer1), nil)
	println("before deposit player1 balance:" + accountBalance.String())

	tx, err := gamePool.GenesisPlayer(&bind.TransactOpts{
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			privateKey, err := crypto.HexToECDSA(accountPlayer1)
			signature, err := crypto.Sign(types.HomesteadSigner{}.Hash(tx).Bytes(), privateKey)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(types.HomesteadSigner{}, signature)
		},
		Context:  context.Background(),
		GasLimit: uint64(900000), // in units
		GasPrice: big.NewInt(20000000000),
		From:     GetAddressFromPrivateKey(accountPlayer1),
		Value:    depositValue,
	}, "1")

	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 2)
		GenesisPlayerTest(gamePool, accounts)
	} else {

		//print player createGameRq
		rq := request.CreateGameRq{
			PlayerAddress: GetAddressFromPrivateKey(accountPlayer1).String(),
			TxHash:        tx.Hash().String(),
			Value:         depositValue.String(),
		}
		rqBytes, _ := json.MarshalIndent(rq, "", " ")
		println(string(rqBytes))

		//print player GameAction
		rq2 := input.GameActionEvent{
			PlayerAddress: GetAddressFromPrivateKey(accountPlayer1).String(),
			GameId:        "1",
			Action:        "rock",
		}
		rqBytes, _ = json.MarshalIndent(rq2, "", " ")
		println(string(rqBytes))

		accountBalance, err = EthClient.BalanceAt(context.Background(), GetAddressFromPrivateKey(accountPlayer1), nil)
		println("after deposit player1 balance:" + accountBalance.String())
	}
}

func OtherPlayerTest(gamePool *contracts.GamePool, accounts *model.GanacheAccounts) {
	//player 2 deposit
	accountPlayer2 := accounts.PrivateKeys[2]
	depositValue := big.NewInt(10000000000)
	accountBalance, err := EthClient.BalanceAt(context.Background(), GetAddressFromPrivateKey(accounts.PrivateKeys[0]), nil)
	println("before deposit player2 balance:" + accountBalance.String())

	tx, err := gamePool.OtherPlayer(&bind.TransactOpts{
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			privateKey, err := crypto.HexToECDSA(accounts.PrivateKeys[0])
			signature, err := crypto.Sign(types.HomesteadSigner{}.Hash(tx).Bytes(), privateKey)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(types.HomesteadSigner{}, signature)
		},
		Context:  context.Background(),
		GasLimit: uint64(9000000), // in units
		GasPrice: big.NewInt(20000000000),
		From:     GetAddressFromPrivateKey(accountPlayer2),
		Value:    depositValue,
	}, "1")
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 2)
		OtherPlayerTest(gamePool, accounts)
	} else {

		//print player createGameRq
		rq := request.JoinGameRq{
			PlayerAddress: GetAddressFromPrivateKey(accountPlayer2).String(),
			TxHash:        tx.Hash().String(),
			Value:         depositValue.String(),
			GameId:        "1",
		}
		rqBytes, _ := json.MarshalIndent(rq, "", " ")
		println(string(rqBytes))

		//print player GameAction
		rq2 := input.GameActionEvent{
			PlayerAddress: GetAddressFromPrivateKey(accountPlayer2).String(),
			GameId:        "1",
			Action:        "rock",
		}
		rqBytes, _ = json.MarshalIndent(rq2, "", " ")
		println(string(rqBytes))

		accountBalance, err = EthClient.BalanceAt(context.Background(), GetAddressFromPrivateKey(accountPlayer2), nil)
		println("after deposit player2 balance:" + accountBalance.String())
	}
}

func WithdrawToWinTest(gamePool *contracts.GamePool, accounts *model.GanacheAccounts) {
	//Withdraw to Winner
	tx, err := gamePool.WithdrawToWin(&bind.TransactOpts{
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			privateKey, err := crypto.HexToECDSA(accounts.PrivateKeys[2])
			signature, err := crypto.Sign(types.HomesteadSigner{}.Hash(tx).Bytes(), privateKey)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(types.HomesteadSigner{}, signature)
		},
		Context:  context.Background(),
		GasLimit: uint64(9000000), // in units
		GasPrice: big.NewInt(20000000000),
		From:     GetAddressFromPrivateKey(accounts.PrivateKeys[0]),
	}, GetAddressFromPrivateKey(accounts.PrivateKeys[1]), "1")
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 2)
		WithdrawToWinTest(gamePool, accounts)
	} else {
		println("txHash: " + tx.Hash().String())
		println("winner player1 withdraw success")
		accountBalance, _ := EthClient.BalanceAt(context.Background(), GetAddressFromPrivateKey(accounts.PrivateKeys[1]), nil)
		println("player1 balance:" + accountBalance.String())
		accountBalance, _ = EthClient.BalanceAt(context.Background(), GetAddressFromPrivateKey(accounts.PrivateKeys[2]), nil)
		println("player2 balance:" + accountBalance.String())
	}
}

func ReturnAGameTest(gamePool *contracts.GamePool, accounts *model.GanacheAccounts) {
	//header, err := helper.EthClient.BlockNumber(context.Background())
	game, err := gamePool.ReturnAGame(&bind.CallOpts{
		Pending: true,
		//BlockNumber: big.NewInt(int64(header)),
		Context: context.Background(),
		From:    common.HexToAddress(accounts.Addresses[3]),
	}, "1")
	if err != nil {
		fmt.Println(err)
	}
	print(game[:])
}

func GetAddressFromPrivateKey(privateKeyParam string) common.Address {
	privateKey, _ := crypto.HexToECDSA(privateKeyParam)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress
}
