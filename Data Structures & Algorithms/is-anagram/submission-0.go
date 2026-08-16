import "reflect"

func isAnagram(s string, t string) bool {
	sCount, tCount := make(map[rune]int), make(map[rune]int)
	for _, runeValue := range s {
		if _, prs := sCount[runeValue]; !prs {
			sCount[runeValue] = 1
		} else {
			sCount[runeValue] += 1
		}
	}

	for _, runeValue := range t {
		if _, prs := tCount[runeValue]; !prs {
			tCount[runeValue] = 1
		} else {
			tCount[runeValue] += 1
		}
	}

	return reflect.DeepEqual(sCount, tCount)
}