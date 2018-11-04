package sipanonymizer

// ProcessMessage reads in messages from bytes
func ProcessMessage(v []byte) SipMsg {
	return parse(v)
}

// Get a string from a slice of bytes
// Checks the bounds to avoid any range errors
func getString(sl []byte, from, to int) string {
	// Limit if over cap
	if from > cap(sl) {
		return ""
	}
	if to > cap(sl) {
		return string(sl[from:])
	}
	return string(sl[from:to])
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
