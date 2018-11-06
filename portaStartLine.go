package sipanonymizer

import (
	"bytes"
	"fmt"
)

/*
2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| RECEIVED message from UDP:192.168.64.92:5061 at UDP:192.168.67.224:5060:
2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| SENDING message to UDP:192.168.67.224:5060 from UDP:192.168.64.92:5061:
*/

// processPortaStartLine hides user's personal data in Porta's start line
func processPortaStartLine(v []byte) {
	pos := bytes.LastIndexByte(v, '|')
	if pos < 0 {
		fmt.Println("error has occured, there are no '|' chars in line:\n", v)
		return
	}
	state := FieldText

	vLen := len(v)
	for pos < vLen {
		// FSM
		switch state {
		case FieldText:
			tr := getBytes(v, pos, pos+4)
			if bytes.Equal(tr, []byte("UDP:")) ||
				bytes.Equal(tr, []byte("TCP:")) ||
				bytes.Equal(tr, []byte("TLS:")) {
				pos = pos + 4
				pos = pos + processHost(v[pos:])
				continue
			}
		}
		pos++
	}

}
