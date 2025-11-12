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

func assertCorrectSlice(t testing.TB, got, want []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("lunghezza diversa: got %d want %d", len(got), len(want))
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("elemento %d: got %d want %d", i, got[i], want[i])
		}
	}
}

func generaInteroCasuale(min, max int) int {
	return min + rand.Intn(max-min)
}
