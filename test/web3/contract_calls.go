package web3

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"log"
	"math/big"
	"strings"
	"time"
	"web3game/contracts"
	"web3game/contracts/model"
)

func GenesisPlayerTest(gamePool *contracts.GamePool, accounts *model.GanacheAccounts) bool {
	// player 1 deposit
	accountPlayer1 := accounts.PrivateKeys[1]
	depositValue := big.NewInt(10000000000)
	_, err = gamePool.GenesisPlayer(&bind.TransactOpts{
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
		return false
	} else {

		//println player createGameRq
		//rq := request.CreateGameRq{
		//	PlayerAddress: GetAddressFromPrivateKey(accountPlayer1).Hex(),
		//	TxHash:        tx.Hash().String(),
		//	Value:         depositValue.String(),
		//}
		//rqBytes, _ := json.MarshalIndent(rq, "", " ")
		//println(string(rqBytes))

		//println player GameAction
		//rq2 := request.RockPaperActionRq{
		//	PlayerAddress: GetAddressFromPrivateKey(accountPlayer1).String(),
		//	GameId:        "1",
		//	Action:        "rock",
		//}
		//rqBytes, _ = json.MarshalIndent(rq2, "", " ")
		//println(string(rqBytes))
		return true
	}
}

func OtherPlayerTest(gamePool *contracts.GamePool, accounts *model.GanacheAccounts) bool {
	//player 2 deposit
	accountPlayer2 := accounts.PrivateKeys[2]
	depositValue := big.NewInt(10000000000)
	_, err = gamePool.OtherPlayer(&bind.TransactOpts{
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			privateKey, err := crypto.HexToECDSA(accountPlayer2)
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
		return false
	} else {

		//println player createGameRq
		//rq := request.JoinGameRq{
		//	PlayerAddress: GetAddressFromPrivateKey(accountPlayer2).Hex(),
		//	TxHash:        tx.Hash().String(),
		//	Value:         depositValue.String(),
		//	GameId:        "1",
		//}
		//rqBytes, _ := json.MarshalIndent(rq, "", " ")
		//println(string(rqBytes))

		//println player GameAction
		//rq2 := request.RockPaperActionRq{
		//	PlayerAddress: GetAddressFromPrivateKey(accountPlayer2).String(),
		//	GameId:        "1",
		//	Action:        "rock",
		//}
		//rqBytes, _ = json.MarshalIndent(rq2, "", " ")
		//println(string(rqBytes))
		return true
	}
}

func WithdrawToWinTest(gamePool *contracts.GamePool, winner string, ownerPrivateKey string) bool {
	//Withdraw to Winner
	_, err := gamePool.WithdrawToWin(&bind.TransactOpts{
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			privateKeyGen, err := crypto.HexToECDSA(ownerPrivateKey)
			signature, err := crypto.Sign(types.HomesteadSigner{}.Hash(tx).Bytes(), privateKeyGen)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(types.HomesteadSigner{}, signature)
		},
		Context:  context.Background(),
		GasLimit: uint64(9000000), // in units
		GasPrice: big.NewInt(20000000000),
	}, common.HexToAddress(strings.ToLower(winner)), "1")
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 2)
		WithdrawToWinTest(gamePool, winner, ownerPrivateKey)
		return false
	} else {
		return true
	}
}

func ReturnAGameTest(gamePool *contracts.GamePool, accounts *model.GanacheAccounts) {
	//header, err := helper.EthClient.BlockNumber(context.Background())
	game, err := gamePool.ReturnAGame(&bind.CallOpts{
		Pending: true,
		//BlockNumber: big.NewInt(int64(header)),
		Context: context.Background(),
		From:    common.HexToAddress(accounts.Addresses[4]),
	}, "1")
	if err != nil {
		fmt.Println(err)
	}
	println(game[:])
}

func PrintContractBalance(contractAddress common.Address) {
	println()
	accountBalance, _ := EthClient.BalanceAt(context.Background(), contractAddress, nil)
	println("Contract " + contractAddress.String() + " is balance:" + accountBalance.String())
}

func PrintNormalBalance(contractAddress common.Address) {
	accountBalance, _ := EthClient.BalanceAt(context.Background(), contractAddress, nil)
	println(contractAddress.String() + " is balance:" + accountBalance.String())
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

func Public(privateKey string) (publicKey string) {
	var e ecdsa.PrivateKey
	e.D, _ = new(big.Int).SetString(privateKey, 16)
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	return fmt.Sprintf("%x", elliptic.Marshal(secp256k1.S256(), e.X, e.Y))
}
