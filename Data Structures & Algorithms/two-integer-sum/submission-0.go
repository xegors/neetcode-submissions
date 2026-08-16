func twoSum(nums []int, target int) []int {
	count := make(map[int]int)

	for i, val := range nums {
		diff := target - val
		if j, found := count[diff]; found {
			return []int{j, i}
		}
		count[val] = i
	}
	return []int{}
}