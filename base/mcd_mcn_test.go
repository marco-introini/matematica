package base

import (
	"fmt"
	"testing"
)

func TestMCD(t *testing.T) {
	t.Run("MCD 110 e 11", func(t *testing.T) {
		got := Mcd(110, 11)
		want := 11
		assertCorrectValue(t, got, want)
	})
	t.Run("MCD 50 e 4", func(t *testing.T) {
		got := Mcd(50, 4)
		want := 2
		assertCorrectValue(t, got, want)
	})
}

func TestMCM(t *testing.T) {
	t.Run("MCM 110 e 11", func(t *testing.T) {
		got := Mcm(110, 11)
		want := 110
		assertCorrectValue(t, got, want)
	})
}

func TestMcdSlice(t *testing.T) {
	numbers := []int{110, 10, 50, 4}
	got := McdSlice(numbers)
	want := 2
	assertCorrectValue(t, got, want)
}

func ExampleMcd() {
	mcd := Mcd(110, 11)
	fmt.Println(mcd)
	// Output: 11
}

func BenchmarkMcd(b *testing.B) {
	for b.Loop() {
		//a := generaInteroCasuale(1, 1000000)
		//b := generaInteroCasuale(1, 1000000)
		Mcd(1100, 13222)
	}
}
