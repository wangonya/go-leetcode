package goleetcode

func romanToInt(s string) int {
	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	smallest, total := 0, 0

	for _, roman := range s {
		v := romans[roman]

		if smallest < v {
			total += v - smallest - smallest
		} else {
			total += v
		}

		smallest = v
	}

	return total
}
