package sipanonymizer

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
				if getString(v, pos, pos+4) == "sip:" {
					if indexAtChar(v) == -1 {
						state = FieldHost
					} else {
						state = FieldUser
					}
					pos = pos + 4
					continue
				}
				if getString(v, pos, pos+5) == "sips:" {
					if indexAtChar(v) == -1 {
						state = FieldHost
					} else {
						state = FieldUser
					}
					pos = pos + 5
					continue
				}
				if getString(v, pos, pos+4) == "tel:" {
					if indexAtChar(v) == -1 {
						state = FieldHost
					} else {
						state = FieldUser
					}
					pos = pos + 4
					continue
				}
				// Look for a Tag identifier
				if getString(v, pos, pos+4) == "tag=" {
					// there is nothing to hide in the rest of line
					return
				}
				// Look for other identifiers and ignore
				if v[pos] == '=' {
					state = FieldIgnore
					pos++
					continue
				}
				// Look for a User Type identifier
				if getString(v, pos, pos+5) == "user=" {
					// there is nothing to hide in the rest of line
					return
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
				state = FieldHost
				pos++
				continue
			}
			// hide displayName
			v[pos] = maskChar

		case FieldHost:
			if v[pos] == ':' {
				state = FieldPort
				pos++
				continue
			}
			if v[pos] == ';' || v[pos] == '>' {
				return
			}
			if v[pos] == '.' {
				pos++
				continue
			}
			// hide displayName
			v[pos] = maskChar

		case FieldPort:
			if v[pos] == ';' || v[pos] == '>' || v[pos] == ' ' {
				return
			}
			// hide displayName
			v[pos] = maskChar
		}
		pos++
	}
}
