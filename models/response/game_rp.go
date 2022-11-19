package response

type GameEvent struct {
	GameId          string `json:"gameId"`
	PlayerAddresses string `json:"playerAddresses"`
	Status          int    `json:"status"` //1 open, 2 running, 3 close
}
