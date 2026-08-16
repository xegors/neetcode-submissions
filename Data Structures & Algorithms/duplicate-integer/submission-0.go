func hasDuplicate(nums []int) bool {
	count := make(map[int]bool)

	for _, num := range nums {
		if _, prs := count[num]; !prs {
			count[num] = true
		} else {
			return true
		}
	}
	return false
}