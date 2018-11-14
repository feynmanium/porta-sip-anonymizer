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
	lines := bytes.Split(v, spaceBytes)
	if len(lines) == 0 {
		fmt.Println("error has occured, there are no ':' chars in line:\n", v)
		return
	}

	for _, line := range lines {
		tr := getBytes(line, 0, 4)
		if bytes.Equal(tr, udpBytes) ||
			bytes.Equal(tr, tcpBytes) ||
			bytes.Equal(tr, tlsBytes) {
			processHost(line[5:])
		}
	}
}
