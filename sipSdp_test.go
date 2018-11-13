package sipanonymizer

import (
	"bytes"
	"testing"
	"text/template"
)

func TestProcessSdpConnection(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("c=IN IP4 10.101.6.120"), "c=IN IP4 10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120"},
		{[]byte("c=IN IP4 sip.domain.com"), "c=IN IP4 sip.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com"},
	}
	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSdpConnection(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of processSdpConnection is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func TestProcessSdpOriginator(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("o=PortaSIP 4530741258397867310 1 IN IP4 217.182.47.207"), "o=PortaSIP 4530741258397867310 1 IN IP4 217.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.207"},
		{[]byte("o=PortaSIP 4530741258397867310 1 IN IP4 sip.domain.com"), "o=PortaSIP 4530741258397867310 1 IN IP4 sip.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com"},
	}
	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSdpOriginator(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of processSdpOriginator is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func TestProcessSdpMedia(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("m=audio 42352 RTP/AVP 0 8 9 18 102 103 101"), "m=audio {{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} RTP/AVP 0 8 9 18 102 103 101"},
	}
	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSdpMedia(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of processSdpMedia is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}
