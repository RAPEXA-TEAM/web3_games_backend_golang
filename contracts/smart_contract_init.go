package contracts

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

//var dbLockSC = &sync.Mutex{}
//var gamePoolSC *GamePool
//var EthClient *ethclient.EthClient
//
//func GetGamePoolSC() *GamePool {
//
//	var err error
//	if gamePoolSC == nil {
//		dbLockSC.Lock()
//		defer dbLockSC.Unlock()
//		EthClient, err = ethclient.Dial("http://localhost:8545")
//		if err != nil {
//			log.Fatal(err)
//		}
//		gamePoolSC, err = NewGamePool(
//			common.HexToAddress("0xde3D9aCEB0aDF78Dd374622b094eAB1c71E2FA6c"),
//			EthClient)
//		if err != nil {
//			panic(fmt.Errorf("smart contract fail to connect: %w", err))
//		}
//	}
//	return gamePoolSC
//}

func AddBlockChainListeners() {
	//start ganache
	//deploy contract to `development` network with truffle
	//copy the port and contract address and put the values
	//client, err := ethclient.Dial("centrifugo://localhost:8545")
	client, err := ethclient.Dial("wss://rpc-mumbai.matic.today")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xE029D63dD0eC859ADCeD490Cee1010Cf80caFF36")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
