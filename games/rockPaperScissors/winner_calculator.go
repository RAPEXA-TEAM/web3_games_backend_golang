package rockPaperScissors

import "web3game/models/entity"

func CalculaterRockPaperRoundWinner(user1 *entity.RockPaperAction, user2 *entity.RockPaperAction) (winner string) {
	if user2.Action == "rock" {
		if user1.Action == "rock" {
			winner = "Tie"
		}

		if user1.Action == "paper" {
			winner = user1.PlayerAddress
		}

		if user1.Action == "scissors" {
			winner = user2.PlayerAddress
		}
	}

	if user2.Action == "paper" {
		if user1.Action == "rock" {
			winner = user2.PlayerAddress
		}

		if user1.Action == "paper" {
			winner = "Tie"
		}

		if user1.Action == "scissors" {
			winner = user1.PlayerAddress
		}
	}

	if user2.Action == "scissors" {
		if user1.Action == "rock" {
			winner = user1.PlayerAddress
		}

		if user1.Action == "paper" {
			winner = user2.PlayerAddress
		}

		if user1.Action == "scissors" {
			winner = "Tie"
		}
	}

	return winner
}
