package sipanonymizer

import (
	"bytes"
	"strings"
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
				lhdr := strings.ToLower(string(line[0:spos]))
				switch {
				case lhdr == "v" || lhdr == "via":
					ProcessSipVia(line)
				case lhdr == "f" || lhdr == "from":
					ProcessSipFrom(line)
				case lhdr == "t" || lhdr == "to":
					ProcessSipTo(line)
				case lhdr == "m" || lhdr == "contact":
					ProcessSipContact(line)
				case lhdr == "i" || lhdr == "call-id":
					ProcessSipCallID(line)
				}
			} else if spos == 1 && stype == '=' {
				// SDP: Break up into header and value
				lhdr := strings.ToLower(string(line[0]))
				// Switch on the line header
				switch {
				case lhdr == "m":
					processSdpMedia(line)
				case lhdr == "c":
					processSdpConnection(line)
				case lhdr == "o":
					processSdpOriginator(line)
				} // End of Switch

			}
		}
		output = append(output, line)
	}
	return bytes.Join(output, []byte("\n\t"))
}
