package sipanonymizer

import "testing"

func TestProcessSipFrom(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("From: \"John Doe\" <sip:john@87.252.61.202:5070>;tag=bvbvfhehj"), []byte("From: \"Joh* **e\" <sip:j***@87.***.**.202:****>;tag=bvbvfhehj")},
		{[]byte("f: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("f: \"Joh* **e\" <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("From: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("From: Joh* <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("f: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("f: Joh* <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("From: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("From: <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("f: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("f: <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("From: <sip:john@87.252.61.202>"), []byte("From: <sip:j***@87.***.**.202>")},
		{[]byte("f: <sip:john@87.252.61.202>"), []byte("f: <sip:j***@87.***.**.202>")},
		{[]byte("From: <sip:87.252.61.202>"), []byte("From: <sip:87.***.**.202>")},
		{[]byte("f: <sip:87.252.61.202>"), []byte("f: <sip:87.***.**.202>")},
		{[]byte("From: <sip:anonymous@anonymous.invalid>"), []byte("From: <sip:ano*****s@ano******.invalid>")},
		{[]byte("f: <sip:anonymous@anonymous.invalid>"), []byte("f: <sip:ano*****s@ano******.invalid>")},
		//
		{[]byte("To: \"John Doe\" <sip:john@87.252.61.202:5070>;tag=bvbvfhehj"), []byte("To: \"Joh* **e\" <sip:j***@87.***.**.202:****>;tag=bvbvfhehj")},
		{[]byte("t: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("t: \"Joh* **e\" <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("To: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("To: Joh* <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("t: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("t: Joh* <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("To: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("To: <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("t: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("t: <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("To: <sip:john@87.252.61.202>"), []byte("To: <sip:j***@87.***.**.202>")},
		{[]byte("t: <sip:john@87.252.61.202>"), []byte("t: <sip:j***@87.***.**.202>")},
		{[]byte("To: <sip:87.252.61.202>"), []byte("To: <sip:87.***.**.202>")},
		{[]byte("t: <sip:87.252.61.202>"), []byte("t: <sip:87.***.**.202>")},
		{[]byte("To: <sip:anonymous@anonymous.invalid>"), []byte("To: <sip:ano*****s@ano******.invalid>")},
		{[]byte("t: <sip:anonymous@anonymous.invalid>"), []byte("t: <sip:ano*****s@ano******.invalid>")},
		//
		{[]byte("Contact: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("Contact: \"Joh* **e\" <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("m: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("m: \"Joh* **e\" <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("Contact: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("Contact: Joh* <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("m: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("m: Joh* <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("Contact: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("Contact: <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("m: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("m: <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("Contact: <sip:john@87.252.61.202>"), []byte("Contact: <sip:j***@87.***.**.202>")},
		{[]byte("m: <sip:john@87.252.61.202>"), []byte("m: <sip:j***@87.***.**.202>")},
		{[]byte("Contact: sip:87.252.61.202;transport=udp"), []byte("Contact: sip:87.***.**.202;transport=udp")},
		{[]byte("Contact: sip:87.252.61.202:5060;transport=udp"), []byte("Contact: sip:87.***.**.202:****;transport=udp")},
		{[]byte("m: <sip:87.252.61.202>"), []byte("m: <sip:87.***.**.202>")},
		{[]byte("Contact: sip:87.252.61.202"), []byte("Contact: sip:87.***.**.202")},
		{[]byte("Contact: <sip:anonymous@anonymous.invalid>"), []byte("Contact: <sip:ano*****s@ano******.invalid>")},
		{[]byte("m: <sip:anonymous@anonymous.invalid>"), []byte("m: <sip:ano*****s@ano******.invalid>")},
		//
		{[]byte("Record-Route: <sip:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50>"), []byte("Record-Route: <sip:192.***.**.224;lr;ep;pinhole=UDP:192.***.**.50>")},
		{[]byte("Record-Route: <sip:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50:5060>"), []byte("Record-Route: <sip:192.***.**.224;lr;ep;pinhole=UDP:192.***.**.50:****>")},
		{[]byte("Route: <sip:192.168.67.224;lr;ob;pinhole=UDP:192.168.64.50;ep>"), []byte("Route: <sip:192.***.**.224;lr;ob;pinhole=UDP:192.***.**.50;ep>")},
		{[]byte("Route: <sip:192.168.67.224;lr;ob;pinhole=UDP:192.168.64.50:5060;ep>"), []byte("Route: <sip:192.***.**.224;lr;ob;pinhole=UDP:192.***.**.50:****;ep>")},
		//

	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processURLBasedHeader(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of processURLBasedHeader is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}

func TestProcessSipCallIDHeader(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("Call-id: vjnejivnreivujreiuvjnie"), []byte("Call-id: vjnejivnreivujreiuvjnie")},
		{[]byte("i: vjnejivnreivujreiuvjnie"), []byte("i: vjnejivnreivujreiuvjnie")},
		{[]byte("Call-id: 1232312@192.168.1.10"), []byte("Call-id: 1232312@192.***.*.10")},
		{[]byte("i: 1232312@192.168.1.10"), []byte("i: 1232312@192.***.*.10")},
		{[]byte("Call-id: 1232312@sip.domain.com"), []byte("Call-id: 1232312@sip.******.com")},
		{[]byte("i: 1232312@sip.domain.com"), []byte("i: 1232312@sip.******.com")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		ProcessSipCallID(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of TestProcessSipCallIDHeader is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}

func TestProcessSipViaHeader(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("Via: SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_169eac12baa1"), []byte("Via: SIP/2.0/UDP 10.***.*.120;branch=z9hG4bKf_169eac12baa1")},
		{[]byte("v: SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_169eac12baa1"), []byte("v: SIP/2.0/UDP 10.***.*.120;branch=z9hG4bKf_169eac12baa1")},
		{[]byte("Via: SIP/2.0/TCP 10.101.6.120;maddr=10.10.10.10;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa1"), []byte("Via: SIP/2.0/TCP 10.***.*.120;maddr=10.**.**.10;received=8.*.*.8;rport=****;branch=z9hG4bKf_169eac12baa1")},
		{[]byte("v: SIP/2.0/TCP 10.101.6.120;maddr=10.10.10.10;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa1"), []byte("v: SIP/2.0/TCP 10.***.*.120;maddr=10.**.**.10;received=8.*.*.8;rport=****;branch=z9hG4bKf_169eac12baa1")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSipVia(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of TestProcessSipViaHeader is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}
