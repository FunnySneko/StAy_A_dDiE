package objects

func NewHand(diceCount int) Hand {
	return Hand{
		dice:make([]Die, diceCount),
	}
}

func(hand *Hand) RollDie(index int) {
	hand.dice[index].Roll()
}

func(hand *Hand) GetPoints() int {
	var points int
	for _, die := range hand.dice {
		points += die.value
	}
	return points
}

type Hand struct {
	dice []Die
}