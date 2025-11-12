package base

func Mcd2(a, b int) int {
	// Algoritmo di Euclide
	// Versione Iterativa
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Mcd(a, b int) int {
	// Algoritmo di Euclide
	// Versione Ricorsiva
	if b == 0 {
		return a
	}
	return Mcd(b, a%b)
}

func Mcm(a, b int) int {
	return (a * b) / Mcd(a, b)
}
