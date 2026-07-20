func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	lastSeen := make(map[rune]int)

	left, maxLen := 0, 0

	for right, ch := range runes {
		if idx, exists := lastSeen[ch]; exists && idx >= left {
			left = idx+1
		}

		lastSeen[ch] = right

		if right -left + 1 > maxLen {
			maxLen = right-left+1
		}
	}

	return maxLen
}
