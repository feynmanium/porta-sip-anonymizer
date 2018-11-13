package sipanonymizer

import (
	"bytes"
	"testing"
	"text/template"
)

func TestProcessHost(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("192.168.192.100"), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100"},
		{[]byte("192.168.192.100 "), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100 "},
		{[]byte("192.168.192.100 123"), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100 123"},
		{[]byte("192.168.192.100;"), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100;"},
		{[]byte("192.168.192.100;rport=123;received=8.8.8.8"), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100;rport=123;received=8.8.8.8"},
		{[]byte("192.168.192.100>"), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100>"},
		{[]byte("192.168.192.100:"), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100:"},
		{[]byte("192.168.192.100:5060"), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}"},
		{[]byte("192.168.192.100:5060 "), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} "},
		{[]byte("192.168.192.100:5060 123"), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} 123"},
		{[]byte("192.168.192.100:5060;"), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};"},
		{[]byte("192.168.192.100:5060>"), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("192.168.192.100:5060:"), "192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}:"},
		{[]byte("domain.com"), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com"},
		{[]byte("domain.com "), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com "},
		{[]byte("domain.com 123"), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com 123"},
		{[]byte("domain.com;"), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com;"},
		{[]byte("domain.com;rport=123"), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com;rport=123"},
		{[]byte("domain.com>"), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com>"},
		{[]byte("domain.com:"), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com:"},
		{[]byte("domain.com:5060"), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}"},
		{[]byte("domain.com:5060 "), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} "},
		{[]byte("domain.com:5060 123"), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} 123"},
		{[]byte("domain.com:5060;"), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};"},
		{[]byte("domain.com:5060>"), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("domain.com:5060:"), "dom{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}:"},
		{[]byte("sip.domain.com:5060"), "sip.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}"},
		{[]byte("sip.domain.com:5060 "), "sip.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} "},
		{[]byte("sip.domain.com:5060 123"), "sip.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} 123"},
		{[]byte("sip.domain.com:5060;"), "sip.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};"},
		{[]byte("sip.domain.com:5060>"), "sip.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("sip.domain.com:5060:"), "sip.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}:"},
		{[]byte("a.very.complex-domain.co.uk:8080"), "a.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}-{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.uk:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}"},
		{[]byte("a.very.complex-domain.co.uk:8080 "), "a.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}-{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.uk:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} "},
		{[]byte("a.very.complex-domain.co.uk:8080 123"), "a.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}-{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.uk:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} 123"},
		{[]byte("a.very.complex-domain.co.uk:8080;"), "a.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}-{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.uk:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};"},
		{[]byte("a.very.complex-domain.co.uk:8080>"), "a.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}-{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.uk:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("a.very.complex-domain.co.uk:8080:"), "a.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}-{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.uk:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}:"},
	}
	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processHost(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of processHost is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func BenchmarkProcessHostIPPort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processHost([]byte("192.168.192.100:5060"))
	}
}

func BenchmarkProcessHostDomainPort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processHost([]byte("a.very.complex-domain.co.uk:8080"))
	}
}
