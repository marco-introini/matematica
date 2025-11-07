package main

func MCD_iterativa(a, b int) int {
	// Algoritmo di Euclide
	// Versione Iterativa
	for b != 0 {
		a, b = b, a%b
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
