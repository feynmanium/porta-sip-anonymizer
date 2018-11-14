package sipanonymizer

import (
	"bytes"
)

/*
 RFC 3261 - https://www.ietf.org/rfc/rfc3261.txt - 8.1.1.7 Via

 The Via header field indicates the transport used for the transaction
and identifies the location where the response is to be sent.  A Via
header field value is added only after the transport that will be
used to reach the next hop has been selected (which may involve the
usage of the procedures in [4]).

SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_169eac12baa17054e0adfb3_I
SIP/2.0/TCP 10.101.6.120;maddr=9.9.9.9;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa17054e0adfb3_I
*/

func processSipVia(v []byte) {

	// skip SIP/2.0/UDP
	pos := bytes.Index(v, sipCapBytes)
	if pos < 0 {
		// not a Via
		return
	}
	pos = pos + 12
	seenHost := false
	state := FieldBase
	rportPos := bytes.Index(v, rportBytes)
	maddrPos := bytes.Index(v, maddrBytes)
	recvPos := bytes.Index(v, receivedBytes)

	// Loop through the bytes making up the line
	vLen := len(v)
	for pos < vLen {
		// FSM
		switch state {
		case FieldBase:
			if v[pos] == ';' {
				pos++
				continue
			}
			// we are right after 'SIP/2.0/UDP ', at the begining of host
			if !seenHost {
				pos = pos + processHost(v[pos:])
				seenHost = true
				continue
			}
			// Look for a Rport identifier
			if pos == rportPos {
				state = FieldPort
				pos = pos + 6
				continue
			}
			// Look for a maddr identifier
			if pos == maddrPos {
				pos = pos + 6
				pos = pos + processHost(v[pos:])
				continue
			}
			// Look for a recevived identifier
			if pos == recvPos {
				pos = pos + 9
				pos = pos + processHost(v[pos:])
				continue
			}
		case FieldPort:
			if v[pos] == ';' {
				state = FieldBase
				pos++
				continue
			}
			v[pos] = maskChar
		}
		pos++
	}
}
