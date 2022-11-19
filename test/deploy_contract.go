package test

//deploy contract to use
import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"
	"web3game/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var EthClient, err = ethclient.Dial("http://127.0.0.1:8545")

func DeployContract(privateKeyParam string) *contracts.GamePool {
	//client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")

	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyParam)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := EthClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	accountBalance, err := EthClient.BalanceAt(context.Background(), fromAddress, nil)
	println("Contract balance:" + accountBalance.String())

	auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(900000) // in units
	auth.GasPrice = big.NewInt(20000000000)
	//auth.Signer = func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
	//	signature, err := crypto.Sign(types.HomesteadSigner{}.Hash(tx).Bytes(), privateKey)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return tx.WithSignature(types.HomesteadSigner{}, signature)
	//}

	address, tx, gamePool, err := contracts.DeployGamePool(auth, EthClient)
	time.Sleep(2 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deployed Contract Address: " + address.Hex())  // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println("Deployed Contract TxHash: " + tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0
	return gamePool
}
