func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]bool)

    left, maxLen := 0, 0

    for right := 0; right < len(s); right++ {
        for seen[s[right]] {
            delete(seen, s[left])
            left++
        }
        seen[s[right]] = true

        windowLen := right - left + 1

        if windowLen > maxLen {
            maxLen = windowLen
        }
    }

    return maxLen
}
