package sipanonymizer

import (
	"bytes"
)

/*
 RFC 3261 - https://www.ietf.org/rfc/rfc3261.txt

INVITE sip:01798300765@87.252.61.202;user=phone SIP/2.0
REGISTER sips:ss2.biloxi.example.com:5060 SIP/2.0
SIP/2.0 200 OK

*/

// processSipRequestLine hides user's personal data in SIP request line
func processSipRequestLine(v []byte) {

	if len(v) > 7 && bytes.Equal(getBytes(v, 0, 7), []byte("SIP/2.0")) {
		// it is SIP response, there is nothing to hide
		return
	}

	// it is SIP request
	// skip METHOD
	vLen := len(v)

	pos := bytes.IndexByte(v, ':')
	if pos < 0 {
		return
	}
	pos++

	atPos := bytes.IndexByte(v, '@')
	if atPos < 0 {
		// there is no user part
		processHost(v[pos:])
		return
	}

	// Loop through the bytes making up the line
	for pos < vLen {
		if pos == atPos {
			pos++
			processHost(v[pos:])
			return
		}
		// mask user
		pos = pos + processUser(v[pos:])
	}
}
