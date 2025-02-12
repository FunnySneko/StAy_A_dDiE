package objects

func NewPlayer(diceCount int) Player {
	return Player{
		Hand: NewHand(diceCount),
	}
}

func(player *Player) GetDiceValues() []int {
	var values []int
	for _, die := range player.dice {
		values = append(values, die.value)
	}
	return values
}

type Player struct {
	Hand
	Score int
}