func containsNearbyDuplicate(nums []int, k int) bool {
	window := make(map[int]bool)
	left := 0

	for right := 0; right < len(nums); right++ {
		if window[nums[right]] {
			return true
		}

		window[nums[right]] = true

		if right - left >= k {
			delete(window, nums[left])
			left++
		}
	}

	return false
}
