package objects

import (
	"math/rand"
)

func (die *Die) Roll() {
	oldValue := die.value
	for {
		die.value = rand.Intn(6)
		die.value++
		if die.value != oldValue {break }
	}
} 

type Die struct {
	value int
}