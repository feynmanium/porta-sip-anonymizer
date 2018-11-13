package sipanonymizer

import (
	"bytes"
)

// Finds the first valid Seperate or notes its type
func indexSep(s []byte) (int, byte) {
	pos := bytes.IndexAny(s, ":=")
	if pos > 0 {
		return pos, s[pos]
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
