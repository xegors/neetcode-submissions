func isAnagram(s string, t string) bool {
	sCount := make(map[rune]int)
	for _, runeValue := range s {
		sCount[runeValue] += 1
	}

	for _, runeValue := range t {
		if _, prs := sCount[runeValue]; !prs {
			return false
		}
		sCount[runeValue] -= 1
		if sCount[runeValue] < 0 { return false}
	}

	for _, value := range sCount {
		if value != 0 { return false }
	}

	return true
}