func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		ele := target-nums[i]

		if v, ok := m[ele]; ok {
			if v > i {
				return []int{i, v}
			}
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return nil
}
