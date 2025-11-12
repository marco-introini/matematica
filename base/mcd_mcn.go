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

// McdSlice calcola il massimo comun divisore di uno slice di numeri
func McdSlice(numbers []int) int {
	// Gestione casi limite
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}

	// Calcola il MCD di tutti i numeri nello slice
	risultato := numbers[0]
	for i := 1; i < len(numbers); i++ {
		risultato = Mcd(risultato, numbers[i])
		// Se il MCD diventa 1, possiamo fermarci
		if risultato == 1 {
			break
		}
	}

	return risultato
}

func Mcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	return (a * b) / Mcd(a, b)
}

// McmSlice calcola il minimo comune multiplo di uno slice di numeri
func McmSlice(numbers []int) int {
	// Gestione casi limite
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}

	// Calcola il MCM di tutti i numeri nello slice
	risultato := numbers[0]
	for i := 1; i < len(numbers); i++ {
		risultato = Mcm(risultato, numbers[i])
		// Se uno dei numeri è 0, il risultato è 0
		if risultato == 0 {
			break
		}
	}

	return risultato
}
