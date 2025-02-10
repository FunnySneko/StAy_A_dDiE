package objects

func NewPlayer(diceCount int) Player {
	return Player{
		Hand: NewHand(diceCount),
	}
}

type Player struct {
	Hand
	Score int
}