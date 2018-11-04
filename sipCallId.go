package sipanonymizer

// processSipCallID hides user's personal data in SIP URL
func processSipCallID(v []byte) {
	pos := indexAtChar(v)
	if pos == -1 {
		// there is nothing to hide
		return
	}
	pos++

	vLen := len(v)
	for pos < vLen {
		// don't hide '.' in domain or IP
		if v[pos] != '.' {
			v[pos] = maskChar
		}
		pos++
	}
}
