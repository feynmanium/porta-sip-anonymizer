package sipanonymizer

// processSipCallID hides user's personal data in SIP URL
func processSipCallID(v []byte) {
	pos := getIndexSep(v, '@')
	if pos < 0 {
		return
	}
	processHost(v[pos+1:])
}
