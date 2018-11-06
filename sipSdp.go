package sipanonymizer

import (
	"bytes"
)

/*
c=IN IP4 10.101.6.120
c=IN IP4 sip.domain.com
*/
func processSdpConnection(v []byte) {
	pos := bytes.LastIndexByte(v, byte(' '))
	processHost(v[pos+1:])
}

/*
o=PortaSIP 4530741258397867310 1 IN IP4 217.182.47.207
o=PortaSIP 4530741258397867310 1 IN IP4 sip.domain.com
*/
func processSdpOriginator(v []byte) {
	pos := bytes.LastIndexByte(v, byte(' '))
	processHost(v[pos+1:])
}

/*
m=audio 42352 RTP/AVP 0 8 9 18 102 103 101
*/
func processSdpMedia(v []byte) {
	pos := bytes.IndexByte(v, byte(' ')) + 1
	end := pos + bytes.IndexByte(v[pos:], byte(' '))
	for pos < end {
		v[pos] = maskChar
		pos++
	}
}
