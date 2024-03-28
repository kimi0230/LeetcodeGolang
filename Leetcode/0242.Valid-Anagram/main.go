package validanagram

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	record := make(map[rune]int, len(s))

	for _, v := range s {
		record[v]++
	}
	for _, v := range t {
		record[v]--
		if record[v] < 0 {
			return false
		}
	}
	return true
}

func IsAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	record := make(map[rune]int, len(s))

	for _, v := range s {
		record[v]++
	}
	for _, v := range t {
		record[v]--
	}

	for _, v := range record {
		if v != 0 {
			return false
		}
	}
	return true
}
