func minimumRecolors(blocks string, k int) int {
	left := 0
	whiteCount := 0
	minOperation := k

	for right := 0; right < len(blocks); right++ {
		if blocks[right] == 'W' {
			whiteCount++
		}

		if right - left + 1 == k {
			if whiteCount < minOperation {
				minOperation = whiteCount
			}

			if blocks[left] == 'W' {
				whiteCount--
			}

			left++
		}
	}

	return minOperation
}
