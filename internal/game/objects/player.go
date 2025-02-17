package objects

func NewPlayer(diceCount int) Player {
	return Player{
		Hand: NewHand(diceCount),
	}
}

func (player *Player) GetDiceTotalValue() int {
	return player.totalValue
}

func (player *Player) GetDiceValues() []int {
	var values []int
	for _, die := range player.Dice {
		values = append(values, die.Value)
	}
	return values
}

type Player struct {
	Hand
	Score  int
	Health int
}
