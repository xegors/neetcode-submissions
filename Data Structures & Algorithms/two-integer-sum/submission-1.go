func twoSum(nums []int, target int) []int {
	count := make(map[int]int) // какой индекс у value

	for i, val := range nums {
		diff := target - val // какой хотим найти элемент
		if j, prs := count[diff]; prs {
			return []int{j, i}
		}
		count[val] = i
	}
	return []int{}
}