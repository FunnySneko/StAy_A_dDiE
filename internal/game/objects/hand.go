package objects

func NewHand(diceCount int) Hand {
	return Hand{
		dice:make([]Die, diceCount),
	}
}

func(hand *Hand) CalculateValue() {
	hand.totalValue = 0
	for _, die := range hand.dice {
		hand.totalValue += die.value
	}
}

func(hand *Hand) Reroll() {
	for i := range hand.dice {
		hand.dice[i].Roll()
	}
	hand.CalculateValue()
}

func(hand *Hand) RollDie(index int) {
	hand.dice[index].Roll()
	hand.CalculateValue()
}

type Hand struct {
	dice []Die
	totalValue int
}