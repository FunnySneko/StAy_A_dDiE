package objects

import (
	"math/rand"
)

func (die *Die) Roll() {
	die.value = rand.Intn(6)
	die.value++
}

type Die struct {
	value int
}