package helpers

import (
	"math/rand"
	"time"
)

type Customer struct {
	CustomerId      string
	FirstName       string
	CellPhoneNumber string
}

func RandomeNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
