func characterReplacement(s string, k int) int {
	freq := [26]int{}

	left := 0
	maxFreq := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		index := s[right] - 'A'

		freq[index]++

		if freq[index] > maxFreq {
			maxFreq = freq[index]
		}

		for right - left + 1 - maxFreq > k {
			freq[s[left] - 'A']--

			left++
		}

		if right - left + 1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
