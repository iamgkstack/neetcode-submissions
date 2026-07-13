func isAnagram(s string, t string) bool {
	var count [26]int

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		count[s[i]- 'a']++
		count[t[i]-'a']--
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}

	return true
}
