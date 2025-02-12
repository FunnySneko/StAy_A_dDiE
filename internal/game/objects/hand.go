package objects

func NewHand(diceCount int) Hand {
	return Hand{
		dice:make([]Die, diceCount),
	}
}

func(hand *Hand) Reroll() {
	for i := range hand.dice {
		hand.dice[i].Roll()
	}
}

func(hand *Hand) RollDie(index int) {
	hand.dice[index].Roll()
}

type Hand struct {
	dice []Die
}