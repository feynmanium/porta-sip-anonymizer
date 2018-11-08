package sipanonymizer

func processURLBasedHeader(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipURL(v[pos+2:])
}

// ProcessSipCallID hides user's personal data in SIP h_Call-id
func ProcessSipCallID(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipCallID(v[pos+2:])
}

// ProcessSipVia hides user's personal data in SIP h_Via
func ProcessSipVia(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipVia(v[pos+2:])
}
