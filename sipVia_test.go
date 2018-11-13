package sipanonymizer

import (
	"bytes"
	"testing"
	"text/template"
)

func TestProcessSipVia(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_169eac12baa1"), "SIP/2.0/UDP 10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120;branch=z9hG4bKf_169eac12baa1"},
		{[]byte("SIP/2.0/TCP 10.101.6.120;maddr=10.10.10.10;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa1"), "SIP/2.0/TCP 10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120;maddr=10.{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.10;received=8.{{.Mask}}.{{.Mask}}.8;rport={{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};branch=z9hG4bKf_169eac12baa1"},
	}
	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSipVia(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of TestProcessSipVia is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func BenchmarkProcessSipVia(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processSipVia([]byte("SIP/2.0/TCP 10.101.6.120;maddr=10.10.10.10;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa1"))
	}
}
