package sipanonymizer

import (
	"bytes"
)

/*
 RFC 3261 - https://www.ietf.org/rfc/rfc3261.txt

"John Doe" <sip:john@87.252.61.202>;tag=bvbvfhehj
John <sip:john@87.252.61.202>;tag=bvbvfhehj
sip:john@87.252.61.202;tag=bvbvfhehj
sip:anonymous@anonymous.invalid;tag=bvbvfhehj

*/

// processSipURL hides user's personal data in SIP URL
func processSipURL(v []byte) {
	pos := 0
	state := FieldBase
	atPos := bytes.IndexByte(v, '@')

	// Loop through the bytes making up the line
	vLen := len(v)
	for pos < vLen {
		// FSM
		switch state {
		case FieldBase:
			if v[pos] == '"' && pos == 0 {
				state = FieldNameQ
				pos++
				continue
			}
			if v[pos] != ' ' {
				// Not a space so check for uri types
				if bytes.Equal(getBytes(v, pos, pos+4), []byte("sip:")) {
					pos = pos + 4
					if atPos < 0 {
						// there is no user part
						processHost(v[pos+1:])
						return
					}
					state = FieldUser
					continue
				}
				if bytes.Equal(getBytes(v, pos, pos+5), []byte("sips:")) {
					pos = pos + 5
					if atPos < 0 {
						// there is no user part
						processHost(v[pos+1:])
						return
					}
					state = FieldUser
					continue
				}
				if bytes.Equal(getBytes(v, pos, pos+4), []byte("tel:")) {
					pos = pos + 4
					if atPos < 0 {
						// there is no user part
						processHost(v[pos+1:])
						return
					}
					state = FieldUser
					continue
				}
				// Check for other chrs
				if v[pos] != '<' && v[pos] != '>' && v[pos] != ';' {
					state = FieldName
					continue
				}
			}

		case FieldNameQ:
			if v[pos] == '"' {
				state = FieldName
				pos++
				continue
			}
			// hide displayName
			v[pos] = maskChar

		case FieldName:
			if v[pos] == '<' || v[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			// hide displayName
			v[pos] = maskChar

		case FieldUser:
			if v[pos] == '@' {
				pos++
				processHost(v[pos:])
				return
			}
			// hide displayName
			v[pos] = maskChar
		}
		pos++
	}
}
