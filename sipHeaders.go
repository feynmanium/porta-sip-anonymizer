package sipanonymizer

/*
 RFC 3261 - https://www.ietf.org/rfc/rfc3261.txt

From: "John Doe" <sip:john@87.252.61.202>;tag=bvbvfhehj
f: "John Doe" <sip:john@87.252.61.202>;tag=bvbvfhehj
From: John <sip:john@87.252.61.202>;tag=bvbvfhehj
From: sip:john@87.252.61.202;tag=bvbvfhehj
From: sip:anonymous@anonymous.invalid;tag=bvbvfhehj

*/

// ProcessSipFrom hides user's personal data in SIP h_From
func ProcessSipFrom(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipURL(v[pos+2 : len(v)])
}

// ProcessSipTo hides user's personal data in SIP h_To
func ProcessSipTo(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipURL(v[pos+2 : len(v)])
}

// ProcessSipContact hides user's personal data in SIP h_Contact
func ProcessSipContact(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipURL(v[pos+2 : len(v)])
}

// ProcessSipCallID hides user's personal data in SIP h_Call-id
func ProcessSipCallID(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipCallID(v[pos+2 : len(v)])
}

// ProcessSipVia hides user's personal data in SIP h_Via
func ProcessSipVia(v []byte) {
	pos, sep := indexSep(v)
	if sep != ':' {
		// not a SIP header
		return
	}

	processSipVia(v[pos+2 : len(v)])
}
