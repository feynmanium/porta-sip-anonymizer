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
	atPos := getIndexSep(v, '@')
	pinholePos := bytes.Index(v, pinholeBytes)
	schemePos := bytes.Index(v, sipBytes)
	schemeLen := 4

	if schemePos < 0 {
		schemePos = bytes.Index(v, sipsBytes)
		schemeLen = 5
		if schemePos < 0 {
			schemePos = bytes.Index(v, telBytes)
			schemeLen = 4
		}
	}

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
			// Not a space so check for uri types
			if pos == schemePos {
				pos = pos + schemeLen
				if atPos < 0 {
					// there is no user part
					pos = pos + processHost(v[pos+1:])
					if pinholePos > 0 {
						// pinhole=UDP:
						// len("pinhole=") = 8
						pos = pinholePos + 8 + 4
						processHost(v[pos:])
					}
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

		case FieldNameQ:
			if v[pos] == '"' {
				state = FieldName
				pos++
				continue
			}
			// hide displayName
			pos = pos + processUser(v[pos:])
			continue

		case FieldName:
			if v[pos] == '<' || v[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			// hide displayName
			pos = pos + processUser(v[pos:])
			continue

		case FieldUser:
			if pos == atPos {
				pos++
				processHost(v[pos:])
				return
			}
			// hide displayName
			pos = pos + processUser(v[pos:])
			continue
		}
		pos++
	}
}
