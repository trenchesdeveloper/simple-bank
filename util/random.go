package util

import (
	"math/rand"
	"time"
)

const alpabet = "abcdefghijklmnopqrstuvwxyz"

func init(){
	rand.Seed(time.Now().UnixNano())
}


func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(int64(max-min))
}

func RandomString(n int) string {
	var letters = []rune(alpabet)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return int64(RandomInt(0, 10000))
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "GBP"}

	return currencies[rand.Intn(len(currencies))]
}