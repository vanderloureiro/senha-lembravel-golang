package data

import "math/rand"

func GetOne() string {

	r := rand.Intn(292)
	return data[r]
}
