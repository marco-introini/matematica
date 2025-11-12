package base

import (
	"math/rand"
	"testing"
)

func assertCorrectValue(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func generaInteroCasuale(min, max int) int {
	return min + rand.Intn(max-min)
}
