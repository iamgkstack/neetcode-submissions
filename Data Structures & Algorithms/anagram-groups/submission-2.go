func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int 

		for _, ch := range str {
			count[ch-'a']++
		}
		groups[count] = append(groups[count], str)
	}

	res := make([][]string, 0, len(groups))

	for _, group := range groups {
		res = append(res, group)
	}

	return res
}
