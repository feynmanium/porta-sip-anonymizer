package sipanonymizer

import (
	"bytes"
	"strings"
)

type sipHeader struct {
	Value []byte // Sip Value
}

// SipMsg struct contains parsed SIP message.
// Not all headers are parsed, but only those that contain
// user's personal data.
type SipMsg struct {
	// StartLine     sipHeader
	Headers []sipHeader
	// Sdp      SdpMsg
}

func parse(v []byte) (output SipMsg) {

	lines := bytes.Split(v, []byte("\r\n"))

	for i, line := range lines {
		line = bytes.TrimSpace(line)
		if i == 0 {
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
		output.Headers = append(output.Headers, sipHeader{line})
	}

	return
}

// GetString returns string representation
func (msg SipMsg) GetString() string {
	var out bytes.Buffer
	for _, h := range msg.Headers {
		out.Write(h.Value)
		out.Write([]byte("\r\n"))
	}
	return out.String()
}
