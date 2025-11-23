package util

import (
	"math/rand"
	"time"
)

type Shufflable interface {
	Len() int
	Swap(i, j int)
}

// use sort.*Slice to get Shufflable
func Shuffle(s Shufflable) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := s.Len(); n > 0; n-- {
		s.Swap(r.Intn(n), n-1)
	}
}
