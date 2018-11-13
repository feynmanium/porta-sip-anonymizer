package sipanonymizer

import (
	"bytes"
	"testing"

	"text/template"
)

func TestProcessSipRequestLine(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("INVITE sip:abc@87.252.61.202;user=phone SIP/2.0"), "INVITE sip:a{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202;user=phone SIP/2.0"},
		{[]byte("INVITE sips:abc@87.252.61.202;user=phone SIP/2.0"), "INVITE sips:a{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202;user=phone SIP/2.0"},
		{[]byte("INVITE tel:abc@87.252.61.202;user=phone SIP/2.0"), "INVITE tel:a{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202;user=phone SIP/2.0"},
		{[]byte("INVITE sip:abc@domain.com SIP/2.0"), "INVITE sip:a{{.Mask}}{{.Mask}}@dom{{.Mask}}{{.Mask}}{{.Mask}}.com SIP/2.0"},
		{[]byte("REGISTER sips:ss2.biloxi.example.com SIP/2.0"), "REGISTER sips:ss2.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com SIP/2.0"},
		{[]byte("REGISTER sips:ss2.biloxi.example.com:5060 SIP/2.0"), "REGISTER sips:ss2.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} SIP/2.0"},
		{[]byte("BYE sip:123456@host.domain.com SIP/2.0"), "BYE sip:123{{.Mask}}{{.Mask}}6@host.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com SIP/2.0"},
		{[]byte("BYE sip:abc@8.8.8.8 SIP/2.0"), "BYE sip:a{{.Mask}}{{.Mask}}@8.{{.Mask}}.{{.Mask}}.8 SIP/2.0"},
		{[]byte("ACK sip:abc@domain.com SIP/2.0"), "ACK sip:a{{.Mask}}{{.Mask}}@dom{{.Mask}}{{.Mask}}{{.Mask}}.com SIP/2.0"},
		{[]byte("CANCEL sip:abc@domain.com SIP/2.0"), "CANCEL sip:a{{.Mask}}{{.Mask}}@dom{{.Mask}}{{.Mask}}{{.Mask}}.com SIP/2.0"},
		{[]byte("NOTIFY sip:abc@domain.com SIP/2.0"), "NOTIFY sip:a{{.Mask}}{{.Mask}}@dom{{.Mask}}{{.Mask}}{{.Mask}}.com SIP/2.0"},
		{[]byte("MESSAGE sip:abc@domain.com SIP/2.0"), "MESSAGE sip:a{{.Mask}}{{.Mask}}@dom{{.Mask}}{{.Mask}}{{.Mask}}.com SIP/2.0"},
		{[]byte("SIP/2.0 200 OK"), "SIP/2.0 200 OK"},
		{[]byte("SIP/2.0 401 Unauthorized"), "SIP/2.0 401 Unauthorized"},
		{[]byte("SIP/2.0 100 Trying"), "SIP/2.0 100 Trying"},
	}

	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSipRequestLine(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of ProcessSipRequestLine is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func BenchmarkProcessSipRequestLine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processSipRequestLine([]byte("INVITE sip:abc@87.252.61.202;user=phone SIP/2.0"))
	}
}
