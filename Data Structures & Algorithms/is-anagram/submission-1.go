import "reflect"

func isAnagram(s string, t string) bool {
	sCount, tCount := make(map[rune]int), make(map[rune]int)
	for _, runeValue := range s {
		sCount[runeValue] += 1
	}

	for _, runeValue := range t {
		tCount[runeValue] += 1
	}

	return reflect.DeepEqual(sCount, tCount)
}