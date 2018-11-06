package sipanonymizer

import (
	"testing"
)

func TestProcessPortaStartLine(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| RECEIVED message from UDP:192.168.64.92:5061 at UDP:192.168.67.224:5060:"), []byte("2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| RECEIVED message from UDP:192.***.**.92:**** at UDP:192.***.**.224:****:")},
		{[]byte("2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| SENDING message to UDP:192.168.67.224:5060 from UDP:192.168.64.92:5061:"), []byte("2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| SENDING message to UDP:192.***.**.224:**** from UDP:192.***.**.92:****:")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processPortaStartLine(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of processPortaStartLine is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}

func BenchmarkProcessPortaStartLine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processPortaStartLine([]byte("2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| SENDING message to UDP:192.168.67.224:5060 from UDP:192.168.64.92:5061:"))
	}
}
