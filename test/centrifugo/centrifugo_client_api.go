package centrifugo

import (
	"context"
	"log"
	"web3game/centrifugo"
)

func Fire() {

	ch := "rock_paper_channel:1"
	ctx := context.Background()

	// How to get history.
	historyResult, err := centrifugo.GoCentClient.History(ctx, ch)
	if err != nil {
		log.Fatalf("Error calling history: %v", err)
	}
	log.Printf("History for channel %s, %d messages", ch, len(historyResult.Publications))

	// How to get channels.
	channelsResult, err := centrifugo.GoCentClient.Channels(ctx)
	if err != nil {
		log.Fatalf("Error calling channels: %v", err)
	}
	log.Printf("Channels: %#v", channelsResult.Channels)

	// Get info about nodes.
	info, err := centrifugo.GoCentClient.Info(ctx)
	if err != nil {
		log.Fatalf("Error calling info: %v", err)
	}
	log.Printf("Info: %d Centrifugo nodes running", len(info.Nodes))

	// How to broadcast the same data into 3 different channels in one request.
	chs := []string{"chat_1", "chat_2", "chat_3"}
	broadcastResult, err := centrifugo.GoCentClient.Broadcast(ctx, chs, []byte(`{"input": "test"}`))
	if err != nil {
		log.Fatalf("Error calling broadcast: %v", err)
	}
	var hasPublishError bool
	for i, resp := range broadcastResult.Responses {
		if resp.Error != nil {
			hasPublishError = true
			log.Printf("error broadcasting to %s: %v", chs[i], resp.Error)
		}
	}
	if !hasPublishError {
		log.Printf("Broadcast to %d channels is successful", len(chs))
	}

	// How to remove history.
	err = centrifugo.GoCentClient.HistoryRemove(ctx, ch)
	if err != nil {
		log.Fatalf("Error calling history remove: %v", err)
	}
	log.Print("History for channel removed")

	// How to send 3 commands in one request.
	pipe := centrifugo.GoCentClient.Pipe()
	_ = pipe.AddPublish(ch, []byte(`{"input": "test1"}`))
	_ = pipe.AddPublish(ch, []byte(`{"input": "test2"}`))
	_ = pipe.AddPublish(ch, []byte(`{"input": "test3"}`))
	replies, err := centrifugo.GoCentClient.SendPipe(ctx, pipe)
	if err != nil {
		log.Fatalf("Error sending pipe: %v", err)
	}
	for _, reply := range replies {
		if reply.Error != nil {
			log.Fatalf("Error in pipe reply: %v", err)
		}
	}
	log.Printf("Sent %d publish commands in one HTTP request ", len(replies))

}
