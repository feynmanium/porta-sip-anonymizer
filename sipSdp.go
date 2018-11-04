package sipanonymizer

// hideLastParam hised the last param(IP or domain) in line
func hideLastParam(v []byte) {
	pos := len(v) - 1
	for pos > 0 {
		if v[pos] == '.' {
			pos--
		} else if v[pos] != ' ' {
			v[pos] = maskChar
			pos--
		} else {
			return
		}
	}
}

/*
c=IN IP4 10.101.6.120
c=IN IP4 sip.domain.com
*/
func processSdpConnection(v []byte) {
	_, sep := indexSep(v)
	if sep != '=' {
		// not a SDP attr
		return
	}
	hideLastParam(v)
}

/*
o=PortaSIP 4530741258397867310 1 IN IP4 217.182.47.207
o=PortaSIP 4530741258397867310 1 IN IP4 sip.domain.com
*/
func processSdpOriginator(v []byte) {
	_, sep := indexSep(v)
	if sep != '=' {
		// not a SDP attr
		return
	}
	hideLastParam(v)
}

/*
m=audio 42352 RTP/AVP 0 8 9 18 102 103 101
*/
func processSdpMedia(v []byte) {
	_, sep := indexSep(v)
	if sep != '=' {
		// not a SDP attr
		return
	}
	pos := 0
	state := FieldMedia

	// Loop through the bytes making up the line
	for pos < len(v) {
		// FSM
		switch state {
		case FieldMedia:
			if v[pos] == ' ' {
				state = FieldPort
				pos++
				continue
			}

		case FieldPort:
			if v[pos] == ' ' {
				// there is nothing left to hide in line
				return
			}
			v[pos] = maskChar
		}
		pos++
	}
}
