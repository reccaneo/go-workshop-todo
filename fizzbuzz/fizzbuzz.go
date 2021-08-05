package fizzbuzz

import (
	"math/rand"
	"strconv"
	"time"
)

func Say(n int) string {
	if n%5 == 0 && n%3 == 0 {
		return "FizzBuzz"
	} else if n%5 == 0 {
		return "Buzz"
	} else if n%3 == 0 {
		return "Fizz"
	}

	return strconv.Itoa(n)
}

func RandomFizzBuzzOriginal(rd Intner) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	n := r1.Intn(100) + 1

	return Say(n)
}

type Intner interface {
	Intn(n int) int
}

func RandomFizzBuzz(rd Intner) string {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)
	n := rd.Intn(100) + 1

	return Say(n)
}

type RandomFizzBuzzHandler struct {
	random Intner
}

func (r RandomFizzBuzzHandler) Handler() string {
	n := r.random.Intn(100) + 1

	return Say(n)
}
