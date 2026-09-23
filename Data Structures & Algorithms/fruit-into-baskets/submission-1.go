func totalFruit(fruits []int) int {
	freq := make(map[int]int)

	left := 0
	max := 0

	for right := 0; right < len(fruits); right++ {
		freq[fruits[right]]++

		for len(freq) > 2 {
			freq[fruits[left]]--

			if freq[fruits[left]] == 0 {
				delete(freq, fruits[left])
			}

			left++
		}

		if right - left + 1 > max {
			max = right - left + 1
		}


	}

	return max
}
