package centrifugo

import (
	"github.com/centrifugal/gocent/v3"
	"os/exec"
	"time"
)

var GoCentClient = gocent.New(gocent.Config{
	Addr: CENTRIFUGO_API_ENDPOINT,
	Key:  CENTRIFUGO_API_KEY,
})

func RunCentrifugo() {
	var err error

	//centrifugo --config /home/khoujani/Desktop/flutter_web3_games/flutter_web3_games_back/gocode/config.json

	err = exec.Command("centrifugo", "--config", "/home/khoujani/Desktop/flutter_web3_games/flutter_web3_games_back/gocode/config.json").Start()
	time.Sleep(3 * 1000)
	if err != nil {
		panic(0)
	}
}
