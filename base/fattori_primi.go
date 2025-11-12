package base

func ScomposizioneFattoriPrimi(num int) []int {
	var fattori []int
	divisore := 2

	for num > 1 {
		if num%divisore == 0 {
			fattori = append(fattori, divisore)
			num = num / divisore
			// Non incrementare il divisore, continua a provare lo stesso
		} else {
			divisore++
		}
	}

	return fattori
}
