package sipanonymizer

import (
	"bytes"
	"testing"
	"text/template"
)

func TestProcessSipCallID(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("vjnejivnreivujreiuvjnie"), "vjnejivnreivujreiuvjnie"},
		{[]byte("1232312@192.168.1.10"), "1232312@192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.10"},
		{[]byte("1232312@sip.domain.com"), "1232312@sip.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com"},
	}
	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSipCallID(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of ProcessSipRequestLine is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func BenchmarkProcessSipCallID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processSipCallID([]byte("1232312@sip.domain.com"))
	}
}
