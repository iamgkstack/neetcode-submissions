func characterReplacement(s string, k int) int {
    count := make([]int, 26)

    left, maxFreq, maxLen := 0, 0, 0

    for right := 0; right < len(s); right++ {
        idx := s[right] - 'A'
        count[idx]++

        if count[idx] > maxFreq {
            maxFreq = count[idx]
        }

        // if replacement needed > k, shrink window
        for(right-left+1)-maxFreq > k {
            count[s[left]-'A']--
            left++
        }

        if right-left+1 > maxLen {
            maxLen = right-left+1
        }
    }

    return maxLen
}
