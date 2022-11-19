package realtime_api

import (
	"encoding/json"
	"github.com/centrifugal/gocent/v3"
	"golang.org/x/net/context"
	"web3game/centrifugo"
	"web3game/models/request"
	"web3game/models/response"
)

func PublishRockPaperGame(data *response.GameEvent) bool {
	bytes, err := json.Marshal(data)
	_, err = centrifugo.GoCentClient.Publish(
		context.Background(),
		centrifugo.ROCK_PAPER_CHANNEL+":games",
		bytes,
	)
	if err == nil {
		return true
	} else {
		return false
	}
}

func PublishRockPaperGameAction(data *request.RockPaperActionRq) bool {
	bytes, err := json.Marshal(data)
	_, err = centrifugo.GoCentClient.Publish(
		context.Background(),
		centrifugo.ROCK_PAPER_CHANNEL+":"+data.GameId,
		bytes,
	)
	if err == nil {
		return true
	} else {
		return false
	}
}

func GetRockPaperHistory(gameId string) []*response.RockPaperActionRp {
	result, err := centrifugo.GoCentClient.History(
		context.Background(),
		centrifugo.ROCK_PAPER_CHANNEL+":"+gameId,
		gocent.WithLimit(-1),
	)
	var crp []*response.RockPaperActionRp
	for _, pub := range result.Publications {
		item := &response.RockPaperActionRp{}
		err = json.Unmarshal(pub.Data, item)
		if err == nil {
			crp = append(crp, item)
		}
	}
	if err == nil {
		return crp
	} else {
		return nil
	}
}

func RemoveAHistory(channel string) bool {
	err := centrifugo.GoCentClient.HistoryRemove(
		context.Background(),
		channel,
	)
	if err == nil {
		return true
	} else {
		return false
	}
}
