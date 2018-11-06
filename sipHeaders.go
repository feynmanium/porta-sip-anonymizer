package sipanonymizer

/*
 RFC 3261 - https://www.ietf.org/rfc/rfc3261.txt

From: "John Doe" <sip:john@87.252.61.202>;tag=bvbvfhehj
f: "John Doe" <sip:john@87.252.61.202>;tag=bvbvfhehj
From: John <sip:john@87.252.61.202>;tag=bvbvfhehj
From: sip:john@87.252.61.202;tag=bvbvfhehj
From: sip:anonymous@anonymous.invalid;tag=bvbvfhehj

*/

// processSipFrom hides user's personal data in SIP h_From
func processSipFrom(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipURL(v[pos+2:])
}

// processSipTo hides user's personal data in SIP h_To
func processSipTo(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipURL(v[pos+2:])
}

// processSipContact hides user's personal data in SIP h_Contact
func processSipContact(v []byte) {
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

// processRoute hides user's personal data in SIP h_Record-Route
func processRoute(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipURL(v[pos+2:])
}

// processPrivacyHeader hides user's personal data in SIP h_PAI and h_RPID
func processPrivacyHeader(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipURL(v[pos+2:])
}
