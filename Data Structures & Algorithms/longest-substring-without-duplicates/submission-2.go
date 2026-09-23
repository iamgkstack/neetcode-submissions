func lengthOfLongestSubstring(s string) int {
	freq := make(map[byte]int)

	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		freq[s[right]]++

		for freq[s[right]] > 1 {
			freq[s[left]]--

			if freq[s[left]] == 0 {
				delete(freq, s[left])
			}

			left++
		}

		if right - left + 1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
