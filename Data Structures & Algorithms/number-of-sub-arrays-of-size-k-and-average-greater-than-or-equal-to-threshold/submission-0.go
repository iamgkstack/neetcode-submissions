func numOfSubarrays(arr []int, k int, threshold int) int {
	left := 0
	sum := 0
	count := 0

	for right := 0; right < len(arr); right++ {
		sum += arr[right]

		if right - left + 1 == k {
			if sum / k >= threshold {
				count++
			}

			sum -= arr[left]
			left++
		}
	}

	return count
}
