func checkInclusion(s1 string, s2 string) bool {
	isInclude := false
	if len(s1) > len(s2) {
		return isInclude
	}

	var need [26]int
	var window [26]int

	for i := range s1 {
		need[s1[i] - 'a']++
	}

	left := 0

	for right := 0; right < len(s2); right++ {
		window[s2[right] - 'a']++

		if right - left + 1 == len(s1) {
			
			if window == need {
				isInclude = true
			}

			window[s2[left] - 'a']--
			left++
		}
	}

	return isInclude
}
