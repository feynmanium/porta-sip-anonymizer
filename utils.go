package sipanonymizer

// Finds the first valid Seperate or notes its type
func indexSep(s []byte) (int, byte) {
	const space = ' '
	const semiColon = ':'
	const eqChar = '='

	vLen := len(s)
	if vLen == 0 {
		return -1, space
	}
	pos := 0
	for pos < vLen {
		if s[pos] == semiColon || s[pos] == eqChar {
			return pos, s[pos]
		}
		pos++
	}
	return -1, space
}

func getIndexSep(v []byte, sep byte) int {
	vLen := len(v)
	pos := 0
	for pos < vLen {
		if v[pos] == sep {
			break
		}
		pos++
	}
	if pos == vLen {
		return -1
	}
	return pos
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

func trimSpace(v []byte) []byte {
	const space = ' '
	n := len(v)
	low, high := 0, n
	for low < n && v[low] == space {
		low++
	}
	for high > low && v[high-1] == space {
		high--
	}
	return v[low:high:high]
}

// used for testing purposes
type testingMaskStruct struct {
	Mask string
}

func getTestingMaskStruct() *testingMaskStruct {
	return &testingMaskStruct{string(maskChar)}
}
