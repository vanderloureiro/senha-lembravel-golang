package generator

import "github.com/vanderloureiro/senha-lembravel-golang/pkg/data"

func GeneratePassword() string {

	return data.GetRandomWord()
}
