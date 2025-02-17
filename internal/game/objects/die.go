package objects

import (
	"math/rand"
)

func (die *Die) Roll() {
	oldValue := die.Value
	for {
		die.Value = rand.Intn(6)
		die.Value++
		if die.Value != oldValue {
			break
		}
	}
}

type Die struct {
	Value int
}
