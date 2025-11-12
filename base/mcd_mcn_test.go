package base

import (
	"fmt"
	"testing"
)

// mcd - Massimo comune divisore

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

func ExampleMcdSlice() {
	numeri := []int{110, 10, 50, 4}
	mcd := McdSlice(numeri)
	fmt.Println(mcd)
	// Output: 2
}

func BenchmarkMcd(b *testing.B) {
	for b.Loop() {
		//a := generaInteroCasuale(1, 1000000)
		//b := generaInteroCasuale(1, 1000000)
		Mcd(1100, 13222)
	}
}

// mcm - Minimo comune multiplo

func TestMcm(t *testing.T) {
	t.Run("MCM 110 e 11", func(t *testing.T) {
		got := Mcm(110, 11)
		want := 110
		assertCorrectValue(t, got, want)
	})
}

func TestMcmSlice(t *testing.T) {
	t.Run("MCM di [4, 6, 8]", func(t *testing.T) {
		numbers := []int{4, 6, 8}
		got := McmSlice(numbers)
		want := 24
		assertCorrectValue(t, got, want)
	})
	t.Run("MCM di [3, 5, 7]", func(t *testing.T) {
		numbers := []int{3, 5, 7}
		got := McmSlice(numbers)
		want := 105
		assertCorrectValue(t, got, want)
	})
	t.Run("MCM di [12, 18, 24]", func(t *testing.T) {
		numbers := []int{12, 18, 24}
		got := McmSlice(numbers)
		want := 72
		assertCorrectValue(t, got, want)
	})
}

func ExampleMcmSlice() {
	numeri := []int{4, 6, 8}
	mcm := McmSlice(numeri)
	fmt.Println(mcm)
	// Output: 24
}
