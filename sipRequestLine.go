package sipanonymizer

/*
 RFC 3261 - https://www.ietf.org/rfc/rfc3261.txt

INVITE sip:01798300765@87.252.61.202;user=phone SIP/2.0
SIP/2.0 200 OK

*/

// processSipRequestLine hides user's personal data in SIP request line
func processSipRequestLine(v []byte) {

	if len(v) > 7 && getString(v, 0, 7) == "SIP/2.0" {
		// it is SIP response, there is nothing to hide
		return
	}

	pos := 0

	// it is SIP request
	// skip METHOD
	for pos = 0; pos < len(v); pos++ {
		if v[pos] == ' ' {
			break
		}
	}

	state := FieldBase

	// Loop through the bytes making up the line
	vLen := len(v)
	for pos < vLen {
		// FSM
		switch state {
		case FieldBase:
			if v[pos] != ' ' {
				if v[pos] == '@' {
					state = FieldHost
					pos++
					continue
				}
				sc := getString(v, pos, pos+4)
				// Not a space so check for uri types
				if sc == "sip:" || sc == "tel:" {
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
				if getString(v, pos, pos+5) == "user=" {
					// there is nothing to hide in the rest of line
					return
				}

			}
		case FieldUser:
			if v[pos] == ':' {
				state = FieldPort
				pos++
				continue
			}
			if v[pos] == ';' || v[pos] == '>' {
				state = FieldBase
				pos++
				continue
			}
			if v[pos] == '@' {
				state = FieldHost
				pos++
				continue
			}
			// mask user
			v[pos] = maskChar

		case FieldHost:
			if v[pos] == ':' {
				state = FieldPort
				pos++
				continue
			}
			if v[pos] == ';' || v[pos] == '>' {
				state = FieldBase
				pos++
				continue
			}
			if v[pos] == '.' {
				pos++
				continue
			}
			if v[pos] == ' ' {
				// there is nothing to hide in the rest of line
				return
			}
			// mask host
			v[pos] = maskChar

		case FieldPort:
			if v[pos] == ';' || v[pos] == '>' || v[pos] == ' ' {
				state = FieldBase
				pos++
				continue
			}
			// mask port
			v[pos] = maskChar
		}
		pos++
	}
}
