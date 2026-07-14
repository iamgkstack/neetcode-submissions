type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder

	for _, str := range strs {
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteByte('#')
		builder.WriteString(str)
	}

	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	var res []string

	for i := 0; i < len(encoded); {
		j := i

		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])

		word := encoded[j+1 : j+1+length]
		res = append(res, word)

		i = j+1+length
	}

	return res
}
