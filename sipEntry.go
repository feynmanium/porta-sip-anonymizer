package sipanonymizer

import (
	"bytes"
)

func parse(v []byte) []byte {
	lines := bytes.Split(v, []byte("\n\t"))
	output := [][]byte{}

	for i, line := range lines {
		line = bytes.TrimSpace(line)
		if i == 0 {
			processPortaStartLine(line)
		} else if i == 1 {
			// For the first line parse the request
			processSipRequestLine(line)
		} else {
			// For subsequent lines split in sep (: for sip, = for sdp)
			spos, stype := indexSep(line)
			if spos > 1 && stype == ':' {
				// SIP: Break up into header and value
				lhdr := bytes.ToLower(line[0:spos])
				switch {
				case bytes.Equal(lhdr, []byte("via")) || bytes.Equal(lhdr, []byte("v")):
					ProcessSipVia(line)
				case bytes.Equal(lhdr, []byte("from")) || bytes.Equal(lhdr, []byte("f")) ||
					bytes.Equal(lhdr, []byte("to")) || bytes.Equal(lhdr, []byte("t")) ||
					bytes.Equal(lhdr, []byte("contact")) || bytes.Equal(lhdr, []byte("m")) ||
					bytes.Equal(lhdr, []byte("route")) || bytes.Equal(lhdr, []byte("record-route")) ||
					bytes.Equal(lhdr, []byte("remote-party-id")) || bytes.Equal(lhdr, []byte("p-asserted-identity")):
					processURLBasedHeader(line)
				case bytes.Equal(lhdr, []byte("call-id")) || bytes.Equal(lhdr, []byte("i")):
					ProcessSipCallID(line)
				}
			} else if spos == 1 && stype == '=' {
				// SDP: Break up into header and value
				lhdr := line[0]
				// Switch on the line header
				switch {
				case lhdr == 'm':
					processSdpMedia(line)
				case lhdr == 'c':
					processSdpConnection(line)
				case lhdr == 'o':
					processSdpOriginator(line)
				} // End of Switch
			}
		}
		output = append(output, line)
	}
	return bytes.Join(output, []byte("\n\t"))
}
