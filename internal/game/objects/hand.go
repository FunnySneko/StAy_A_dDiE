package objects

func NewHand(diceCount int) Hand {
	var hand Hand
	hand.Dice = make([]Die, diceCount)
	hand.RollOpportunities = make([]int, diceCount)
	for i := range hand.Dice {
		hand.RollOpportunities[i] = 1
	}
	return hand
}

func (hand *Hand) CalculateValue() {
	hand.totalValue = 0
	for _, die := range hand.Dice {
		hand.totalValue += die.Value
	}
}

func (hand *Hand) Reroll() {
	for i := range hand.Dice {
		hand.Dice[i].Roll()
	}
	hand.CalculateValue()
	for i := range hand.RollOpportunities {
		hand.RollOpportunities[i] = 1
	}
}

func (hand *Hand) RollDie(index int) {
	hand.Dice[index].Roll()
	hand.CalculateValue()
	hand.RollOpportunities[index] = 0
}

func (hand *Hand) SetDie(index, value int) {
	hand.Dice[index].Value = value
	hand.CalculateValue()
}

type Hand struct {
	Dice              []Die
	totalValue        int
	RollOpportunities []int
}
