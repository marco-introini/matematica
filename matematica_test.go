package main

import "testing"

func TestMCD_ricorsiva(t *testing.T) {
	t.Run("MCD 110 e 11", func(t *testing.T) {
		got := MCD_ricorsiva(110, 11)
		want := 11

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("MCD 50 e 4", func(t *testing.T) {
		got := MCD_ricorsiva(50, 4)
		want := 2

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
