package sipanonymizer

import (
	"bytes"
)

func getNextLine(v []byte) ([]byte, []byte) {
	if v == nil {
		return nil, nil
	}
	sepPos := getIndexSep(v, '\n')
	if sepPos < 0 {
		return v, nil
	}
	return v[:sepPos], v[sepPos+2:]
}

// parse incoming array of bytes and modify it accordingly
// NOTE: modifies incoming array
func parse(v []byte) {
	i := -1
	line := []byte(nil)
	for {
		if line, v = getNextLine(v); line == nil {
			break
		} else {
			i++
			line = trimSpace(line)
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
					lhdr := line[0:spos]
					switch {
					case bytes.Equal(lhdr, viaCapBytes) || bytes.Equal(lhdr, viaBytes) || bytes.Equal(lhdr, []byte("v")):
						ProcessSipVia(line)
					case bytes.Equal(lhdr, fromCapBytes) || bytes.Equal(lhdr, toCapBytes) || bytes.Equal(lhdr, contactCapBytes) ||
						bytes.Equal(lhdr, routeCapBytes) || bytes.Equal(lhdr, recordRouteCapBytes) || bytes.Equal(lhdr, rpidCapBytes) ||
						bytes.Equal(lhdr, paiCapBytes) ||
						bytes.Equal(lhdr, fromBytes) || bytes.Equal(lhdr, toBytes) || bytes.Equal(lhdr, contactBytes) ||
						bytes.Equal(lhdr, routeBytes) || bytes.Equal(lhdr, recordRouteBytes) || bytes.Equal(lhdr, rpidBytes) ||
						bytes.Equal(lhdr, paiBytes) ||
						bytes.Equal(lhdr, []byte("f")) || bytes.Equal(lhdr, []byte("t")) || bytes.Equal(lhdr, []byte("m")):
						processURLBasedHeader(line)
					case bytes.Equal(lhdr, callIDCapBytes) || bytes.Equal(lhdr, callIDBytes) || bytes.Equal(lhdr, []byte("i")):
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
		}
	}
}
