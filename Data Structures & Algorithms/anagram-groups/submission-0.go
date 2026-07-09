func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		chars := []rune(str)

		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		key := string(chars)

		groups[key] = append(groups[key], str)
	}

	result := make([][]string, 0, len(groups))

	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
