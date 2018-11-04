package sipanonymizer

// Finds the first @ char
func indexAtChar(s []byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == '@' {
			return i
		}
	}
	return -1
}

// Finds the first valid Seperate or notes its type
func indexSep(s []byte) (int, byte) {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ':' || c == '=' {
			return i, c
		}
	}
	return -1, ' '
}
