package game

import (
	"APP/internal/game/objects"
	"math/rand"
)

type turn int

const (
	EnemyTurn turn = iota
	PlayerTurn
)

type Event int

const (
	Fight Event = iota
	Heal
)

func (game *Game) NextTurn() {
	if game.Turn == PlayerTurn {
		game.Turn = EnemyTurn
	} else {
		game.Turn = PlayerTurn
	}
}

func (game *Game) NextStage() Event {
	game.Stage++
	if game.Stage == 1 {
		game.NewFight(2, 1, 5)
		return Fight
	}
	return Fight
}

func (game *Game) NewFight(diceCount, aggressiveness, difficulty int) {
	game.Enemy = objects.NewPlayer(diceCount)
	for i := range game.Enemy.Dice {
		game.Enemy.SetDie(i, difficulty+rand.Intn(2))
	}
	game.enemyAggressiveness = aggressiveness
	game.Enemy.Health = 5 * difficulty
	game.Turn = PlayerTurn
}

func NewGame() Game {
	var game Game
	game.Player = objects.NewPlayer(3)
	for i := range game.Player.Hand.Dice {
		game.Player.SetDie(i, 4+rand.Intn(3))
	}
	game.Player.Health = 25
	game.Stage = 0
	return game
}

func (game *Game) Update() {
	if game.Turn == EnemyTurn {
		game.Turn = PlayerTurn
	} else {
		game.Turn = EnemyTurn
	}
}

type Game struct {
	Player              objects.Player
	Enemy               objects.Player
	enemyAggressiveness int
	Turn                turn
	Stage               int
}
