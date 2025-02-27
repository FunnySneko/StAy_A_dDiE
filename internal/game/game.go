package game

import (
	"APP/internal/game/objects"
	"math"
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

type GameState int
const (
	Running GameState = iota
	Won
	Lost
)

func (game *Game) DealDamage() {
	var victim *objects.Player
	damage := math.Abs(float64(game.Enemy.GetDiceTotalValue() - game.Player.GetDiceTotalValue()))
	if game.Enemy.GetDiceTotalValue() > game.Player.GetDiceTotalValue() {
		victim = &game.Player
	} else {
		victim = &game.Enemy
	}
	victim.Health -= int(damage)
	if game.Enemy.Health <= 0 {
		game.GameState = Won
	}
}

func (game *Game) EnemyMove() {
	var playersMaxValueDie int
	_ = playersMaxValueDie
	empty := true
	for i := range game.Player.Dice {
		if game.Player.RollOpportunities[i] == 1 {
			if empty {
				playersMaxValueDie = i
				empty = false
			} else {
				if game.Player.Dice[i].Value > game.Player.Dice[playersMaxValueDie].Value {
					playersMaxValueDie = i
				}
			}
		}
	}
	if game.Enemy.GetDiceTotalValue() > game.Player.GetDiceTotalValue() {
		if !empty {
			game.Player.RollDie(playersMaxValueDie)
		} else {
			game.Enemy.Reroll()
		}
	} else {
		game.Enemy.Reroll()
	}
	game.DealDamage()
}

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
		game.NewFight(2, 1, 2)
		return Fight
	} else if game.Stage == 2 {
		game.NewFight(4, 1, 2)
		return Fight
	} else {
		game.NewFight(6, 1, 5)
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
	game.GameState = Running
}

func NewGame() Game {
	var game Game
	game.Player = objects.NewPlayer(3)
	for i := range game.Player.Hand.Dice {
		game.Player.SetDie(i, 4+rand.Intn(3))
	}
	game.Player.Health = 50
	game.Stage = 0
	return game
}

type Game struct {
	Player              objects.Player
	Enemy               objects.Player
	enemyAggressiveness int
	Turn                turn
	Stage               int
	GameState			GameState
}