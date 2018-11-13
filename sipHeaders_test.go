package sipanonymizer

import (
	"bytes"
	"testing"
	"text/template"
)

func TestProcessSipFrom(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("From: \"John Doe\" <sip:john@87.252.61.202:5070>;tag=bvbvfhehj"), "From: \"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=bvbvfhehj"},
		{[]byte("f: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), "f: \"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("From: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), "From: Joh{{.Mask}} <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("f: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), "f: Joh{{.Mask}} <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("From: <sip:john@87.252.61.202>;tag=bvbvfhehj"), "From: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("f: <sip:john@87.252.61.202>;tag=bvbvfhehj"), "f: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("From: <sip:john@87.252.61.202>"), "From: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("f: <sip:john@87.252.61.202>"), "f: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("From: <sip:87.252.61.202>"), "From: <sip:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("f: <sip:87.252.61.202>"), "f: <sip:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("From: <sip:anonymous@anonymous.invalid>"), "From: <sip:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid>"},
		{[]byte("f: <sip:anonymous@anonymous.invalid>"), "f: <sip:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid>"},
		//
		{[]byte("To: \"John Doe\" <sip:john@87.252.61.202:5070>;tag=bvbvfhehj"), "To: \"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=bvbvfhehj"},
		{[]byte("t: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), "t: \"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("To: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), "To: Joh{{.Mask}} <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("t: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), "t: Joh{{.Mask}} <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("To: <sip:john@87.252.61.202>;tag=bvbvfhehj"), "To: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("t: <sip:john@87.252.61.202>;tag=bvbvfhehj"), "t: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("To: <sip:john@87.252.61.202>"), "To: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("t: <sip:john@87.252.61.202>"), "t: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("To: <sip:87.252.61.202>"), "To: <sip:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("t: <sip:87.252.61.202>"), "t: <sip:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("To: <sip:anonymous@anonymous.invalid>"), "To: <sip:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid>"},
		{[]byte("t: <sip:anonymous@anonymous.invalid>"), "t: <sip:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid>"},
		//
		{[]byte("Contact: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), "Contact: \"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("m: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), "m: \"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("Contact: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), "Contact: Joh{{.Mask}} <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("m: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), "m: Joh{{.Mask}} <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("Contact: <sip:john@87.252.61.202>;tag=bvbvfhehj"), "Contact: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("m: <sip:john@87.252.61.202>;tag=bvbvfhehj"), "m: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("Contact: <sip:john@87.252.61.202>"), "Contact: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("m: <sip:john@87.252.61.202>"), "m: <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("Contact: sip:87.252.61.202;transport=udp"), "Contact: sip:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202;transport=udp"},
		{[]byte("Contact: sip:87.252.61.202:5060;transport=udp"), "Contact: sip:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};transport=udp"},
		{[]byte("m: <sip:87.252.61.202>"), "m: <sip:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("Contact: sip:87.252.61.202"), "Contact: sip:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202"},
		{[]byte("Contact: <sip:anonymous@anonymous.invalid>"), "Contact: <sip:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid>"},
		{[]byte("m: <sip:anonymous@anonymous.invalid>"), "m: <sip:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid>"},
		//
		{[]byte("Record-Route: <sip:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50>"), "Record-Route: <sip:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ep;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50>"},
		{[]byte("Record-Route: <sip:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50:5060>"), "Record-Route: <sip:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ep;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("Route: <sip:192.168.67.224;lr;ob;pinhole=UDP:192.168.64.50;ep>"), "Route: <sip:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ob;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50;ep>"},
		{[]byte("Route: <sip:192.168.67.224;lr;ob;pinhole=UDP:192.168.64.50:5060;ep>"), "Route: <sip:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ob;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};ep>"},
		//

	}
	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processURLBasedHeader(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of processURLBasedHeader is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func TestProcessSipCallIDHeader(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("Call-id: vjnejivnreivujreiuvjnie"), "Call-id: vjnejivnreivujreiuvjnie"},
		{[]byte("i: vjnejivnreivujreiuvjnie"), "i: vjnejivnreivujreiuvjnie"},
		{[]byte("Call-id: 1232312@192.168.1.10"), "Call-id: 1232312@192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.10"},
		{[]byte("i: 1232312@192.168.1.10"), "i: 1232312@192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.10"},
		{[]byte("Call-id: 1232312@sip.domain.com"), "Call-id: 1232312@sip.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com"},
		{[]byte("i: 1232312@sip.domain.com"), "i: 1232312@sip.{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com"},
	}
	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		ProcessSipCallID(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of TestProcessSipCallIDHeader is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func TestProcessSipViaHeader(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("Via: SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_169eac12baa1"), "Via: SIP/2.0/UDP 10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120;branch=z9hG4bKf_169eac12baa1"},
		{[]byte("v: SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_169eac12baa1"), "v: SIP/2.0/UDP 10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120;branch=z9hG4bKf_169eac12baa1"},
		{[]byte("Via: SIP/2.0/TCP 10.101.6.120;maddr=10.10.10.10;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa1"), "Via: SIP/2.0/TCP 10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120;maddr=10.{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.10;received=8.{{.Mask}}.{{.Mask}}.8;rport={{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};branch=z9hG4bKf_169eac12baa1"},
		{[]byte("v: SIP/2.0/TCP 10.101.6.120;maddr=10.10.10.10;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa1"), "v: SIP/2.0/TCP 10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120;maddr=10.{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.10;received=8.{{.Mask}}.{{.Mask}}.8;rport={{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};branch=z9hG4bKf_169eac12baa1"},
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
			t.Errorf("Result of TestProcessSipViaHeader is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}
