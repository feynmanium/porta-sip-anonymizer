package sipanonymizer

import (
	"bytes"
)

/*
192.168.192.10
8.8.8.8
domain.com
sip.domain.com
*/

// processHost hides most part in the Host(IPv4 or domain) like:
// 192.x.x.10 or sip.xxxxx.com or gxxxxx.com
func processHost(v []byte) int {
	if len(v) == 0 {
		return 0
	}

	pos := 0
	state := FieldPreserveFirstOctet
	dotSeen := 0
	dotCount := 0
	hostLastPos := bytes.IndexAny(v, ":;> ")
	if hostLastPos < 0 {
		hostLastPos = len(v) - 1
	}
	// portLastPos :=

	if v[pos] > 64 {
		// looks like a domain
		state = FieldPreserveFirstDomain

		dotCount = bytes.Count(v[0:hostLastPos], []byte("."))
	}

	vLen := len(v)
	for pos < vLen {
		// FSM
		switch state {
		case FieldPreserveFirstOctet:
			if v[pos] == '.' {
				pos++
				dotSeen++
				state = FieldHostIP
				continue
			}
		case FieldPreserveFirstDomain:
			if v[pos] == '.' {
				dotSeen++
				if dotCount == 1 {
					state = FieldPreserveLastDomain
				} else {
					state = FieldHost
				}
				pos++
				continue
			}
			if dotCount == 1 && pos > 2 {
				// in case of domain.com form
				// just transform into dom***.com
				v[pos] = maskChar
				pos++
				continue
			}
		case FieldHostIP:
			if v[pos] == '.' {
				pos++
				dotSeen++
				continue
			}
			if v[pos] == ':' {
				state = FieldPort
				pos++
				continue
			}
			if v[pos] == ' ' ||
				v[pos] == ';' ||
				v[pos] == '>' {
				// it seems host is over
				return pos
			}
			if dotSeen != 3 {
				v[pos] = maskChar
			}
		case FieldHost:
			if v[pos] == '.' {
				dotSeen++
				if dotCount == dotSeen {
					state = FieldPreserveLastDomain
				}
				pos++
				continue
			}
			if v[pos] == '-' {
				pos++
				continue
			}
			v[pos] = maskChar
		case FieldPreserveLastDomain:
			if pos == hostLastPos {
				if v[pos] == ':' {
					state = FieldPort
					pos++
					continue
				}
				// it seems end of host
				return pos
			}
		case FieldPort:
			if v[pos] < '0' || v[pos] > '9' {
				// port is over
				return pos
			}
			v[pos] = maskChar
		}
		pos++
	}
	return pos
}
