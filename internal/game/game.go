package game

import (
	"APP/internal/game/objects"
)

type turn int
const (
	EnemyTurn turn = iota
	PlayerTurn
)

func NewGame(diceCount int) Game {
	return Game{
		Player: objects.NewPlayer(diceCount),
		Enemy: objects.NewPlayer(diceCount),
		isRunning: true,
		turn: PlayerTurn,
	}
}

func(game *Game) GetTurn() turn {
	return game.turn
}

func(game *Game) Update() {
	if(game.turn == EnemyTurn) {
		game.turn = PlayerTurn
	} else {
		game.turn = EnemyTurn
	}
}

type Game struct {
	Player objects.Player
	Enemy objects.Player
	isRunning bool
	turn turn
}