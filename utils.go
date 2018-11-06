package sipanonymizer

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

// Get a slice from a slice of bytes
// Checks the bounds to avoid any range errors
func getBytes(sl []byte, from, to int) []byte {
	// Limit if over cap
	if from > cap(sl) {
		return nil
	}
	if to > cap(sl) {
		return sl[from:]
	}
	return sl[from:to]
}
