func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)

	minDiff := math.MaxInt
	left := 0

	for right := 0; right < len(nums); right++ {
		if right - left + 1 == k {
			diff := nums[right] - nums[left]

			if diff < minDiff {
				minDiff = diff
			}

			left++
		}
	}

	return minDiff
}
