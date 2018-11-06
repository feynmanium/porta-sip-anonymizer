package sipanonymizer

import (
	"bytes"
)

// processSipCallID hides user's personal data in SIP URL
func processSipCallID(v []byte) {
	pos := bytes.IndexByte(v, '@')
	if pos < 0 {
		// there is nothing to hide
		return
	}
	pos++

	processHost(v[pos:])
}
