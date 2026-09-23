func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)

	baseSatisfied := 0

	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			baseSatisfied += customers[i]
		}
	}

	left := 0
	extra := 0
	maxExtra := 0


	for right := 0; right < n; right++ {
		if grumpy[right] == 1 {
			extra += customers[right]
		}

		if right - left + 1 == minutes {
			if extra > maxExtra {
				maxExtra = extra
			}

			if grumpy[left] == 1 {
				extra -= customers[left]
			}

			left++
		}
	}
	return baseSatisfied + maxExtra
}
