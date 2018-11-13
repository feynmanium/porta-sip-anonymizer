package sipanonymizer

import (
	"bytes"
	"testing"
	"text/template"
)

func TestProcessPortaStartLine(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| RECEIVED message from UDP:192.168.64.92:5061 at UDP:192.168.67.224:5060:"), "2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| RECEIVED message from UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.92:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} at UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}:"},
		{[]byte("2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| SENDING message to UDP:192.168.67.224:5060 from UDP:192.168.64.92:5061:"), "2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| SENDING message to UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} from UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.92:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}:"},
	}
	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processPortaStartLine(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of processPortaStartLine is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func BenchmarkProcessPortaStartLine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processPortaStartLine([]byte("2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| SENDING message to UDP:192.168.67.224:5060 from UDP:192.168.64.92:5061:"))
	}
}
