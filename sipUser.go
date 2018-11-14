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

// processUser hides most part in the username like:
// Joh* **e or 123***
// 123*5
// 123**6
// 123***7
// ...
func processUser(v []byte) int {
	if len(v) == 0 {
		return 0
	}

	pos := 0
	if v[pos] == '"' {
		pos = 1
	}
	userLen := bytes.IndexAny(v[pos:], userEnd) + pos
	if userLen < 0 {
		userLen = len(v)
	}

	preserveFirst := 0
	preserveLastPos := -1

	if userLen > 1 && userLen < 5 {
		// 1*
		// 1**
		// 1***
		preserveFirst = 1
		preserveLastPos = -1
	} else if userLen >= 5 {
		// 123*5
		// 123**6
		// 123***7
		// ...
		preserveFirst = 3
		preserveLastPos = userLen - 1
	}

	for pos < userLen {
		if v[pos] == ' ' {
			pos++
			continue
		}
		if preserveFirst != 0 {
			preserveFirst--
			pos++
			continue
		}
		if pos != preserveLastPos {
			v[pos] = maskChar
		}
		pos++
	}
	return userLen
}
