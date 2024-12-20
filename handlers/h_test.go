package handlers

import (
	"math/rand"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rng.Intn(21)
	if n%2 == 0 {

	}
	t.Errorf("n is even")

}
