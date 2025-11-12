package base

func ScomposizioneFattoriPrimi(num int) []int {
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
