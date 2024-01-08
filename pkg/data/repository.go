package data

import (
	"math/rand"
)

func GetRandomWord() string {

	r := rand.Intn(len(data))
	return data[r]
}
