package main

import (
	"fmt"
)

func MCD_iterativa(a, b int) int {
	// Algoritmo di Euclide
	// Versione Iterativa
	for b != 0 {
		a, b = b, a%b
		// equivalente a
		// a=b
		// b=a%b
	}
	return a
}

func MCD_ricorsiva(a, b int) int {
	// Algoritmo di Euclide
	// Versione Ricorsiva
	if b == 0 {
		return a
	}
	return MCD_ricorsiva(b, a%b)
}

func mcm(a, b int) int {
	return (a * b) / MCD_ricorsiva(a, b)
}

func scomposizioneFattoriPrimi(num int) []int {
	var fattori []int
	divisore := 2

	for {
		if num <= 1 {
			break
		}
		if num%divisore == 0 {
			fattori = append(fattori, divisore)
			num = num / divisore
			divisore = 2
		}
		divisore++
	}

	return fattori
}

func esempioChiamataMatematica() {
	a := 111
	b := 11
	println(MCD_ricorsiva(a, b))
	println(MCD_iterativa(a, b))
	println(mcm(a, b))
	fmt.Println(scomposizioneFattoriPrimi(110))
}
