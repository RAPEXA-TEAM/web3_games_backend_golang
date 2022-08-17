package entity

type Games []Game

type Game struct {
	Id      int    `json:"id"`
	GameId  string `json:"game_id"`
	Players string `json:"players"`
	Status  int    `json:"status"`
}
