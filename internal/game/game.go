package game

import (
	"APP/internal/game/objects"
)

type turn int
const (
	Player turn = iota
	Enemy
)

func NewGame(diceCount int) Game {
	return Game{
		Player: objects.NewPlayer(diceCount),
		Enemy: objects.NewPlayer(diceCount),
		isRunning: true,
		turn: Player,
	}
}

func(game *Game) GetTurn() string {
	if game.turn == Player {
		return "PLAYER"
	} else {
		return "ENEMY"
	}
}

func(game *Game) Update() {

}

type Game struct {
	Player objects.Player
	Enemy objects.Player
	isRunning bool
	turn turn
}