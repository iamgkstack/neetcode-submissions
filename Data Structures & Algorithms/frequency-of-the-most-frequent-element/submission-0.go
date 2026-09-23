func maxFrequency(nums []int, k int) int {
	left := 0
	sum := 0
	maxFreq := 0

	sort.Ints(nums)

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for nums[right] * (right - left + 1) - sum > k {
			sum -= nums[left]

			left++
		}

		if right - left + 1 > maxFreq {
			maxFreq = right - left + 1
		}
	}

	return maxFreq

}
