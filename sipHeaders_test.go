package sipanonymizer

import "testing"

func TestProcessSipFrom(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("From: \"John Doe\" <sip:john@87.252.61.202:5070>;tag=bvbvfhehj"), []byte("From: \"********\" <sip:****@**.***.**.***:****>;tag=bvbvfhehj")},
		{[]byte("f: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("f: \"********\" <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("From: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("From: **** <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("f: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("f: **** <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("From: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("From: <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("f: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("f: <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("From: <sip:john@87.252.61.202>"), []byte("From: <sip:****@**.***.**.***>")},
		{[]byte("f: <sip:john@87.252.61.202>"), []byte("f: <sip:****@**.***.**.***>")},
		{[]byte("From: <sip:87.252.61.202>"), []byte("From: <sip:**.***.**.***>")},
		{[]byte("f: <sip:87.252.61.202>"), []byte("f: <sip:**.***.**.***>")},
		{[]byte("From: <sip:anonymous@anonymous.invalid>"), []byte("From: <sip:*********@*********.*******>")},
		{[]byte("f: <sip:anonymous@anonymous.invalid>"), []byte("f: <sip:*********@*********.*******>")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		ProcessSipFrom(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of TestProcessSipFrom is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}

func TestProcessSipTo(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("To: \"John Doe\" <sip:john@87.252.61.202:5070>;tag=bvbvfhehj"), []byte("To: \"********\" <sip:****@**.***.**.***:****>;tag=bvbvfhehj")},
		{[]byte("t: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("t: \"********\" <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("To: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("To: **** <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("t: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("t: **** <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("To: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("To: <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("t: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("t: <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("To: <sip:john@87.252.61.202>"), []byte("To: <sip:****@**.***.**.***>")},
		{[]byte("t: <sip:john@87.252.61.202>"), []byte("t: <sip:****@**.***.**.***>")},
		{[]byte("To: <sip:87.252.61.202>"), []byte("To: <sip:**.***.**.***>")},
		{[]byte("t: <sip:87.252.61.202>"), []byte("t: <sip:**.***.**.***>")},
		{[]byte("To: <sip:anonymous@anonymous.invalid>"), []byte("To: <sip:*********@*********.*******>")},
		{[]byte("t: <sip:anonymous@anonymous.invalid>"), []byte("t: <sip:*********@*********.*******>")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		ProcessSipTo(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of TestProcessSipTo is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}

func TestProcessSipContact(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("Contact: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("Contact: \"********\" <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("m: \"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("m: \"********\" <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("Contact: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("Contact: **** <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("m: John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("m: **** <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("Contact: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("Contact: <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("m: <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("m: <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("Contact: <sip:john@87.252.61.202>"), []byte("Contact: <sip:****@**.***.**.***>")},
		{[]byte("m: <sip:john@87.252.61.202>"), []byte("m: <sip:****@**.***.**.***>")},
		{[]byte("Contact: sip:87.252.61.202;transport=udp"), []byte("Contact: sip:**.***.**.***;transport=udp")},
		{[]byte("Contact: sip:87.252.61.202:5060;transport=udp"), []byte("Contact: sip:**.***.**.***:****;transport=udp")},
		{[]byte("m: <sip:87.252.61.202>"), []byte("m: <sip:**.***.**.***>")},
		{[]byte("Contact: <sip:anonymous@anonymous.invalid>"), []byte("Contact: <sip:*********@*********.*******>")},
		{[]byte("m: <sip:anonymous@anonymous.invalid>"), []byte("m: <sip:*********@*********.*******>")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		ProcessSipContact(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of TestProcessSipContact is incorrect:\n src  %s\n want %s\n got  %s",
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
		{[]byte("Call-id: 1232312@192.168.1.10"), []byte("Call-id: 1232312@***.***.*.**")},
		{[]byte("i: 1232312@192.168.1.10"), []byte("i: 1232312@***.***.*.**")},
		{[]byte("Call-id: 1232312@sip.domain.com"), []byte("Call-id: 1232312@***.******.***")},
		{[]byte("i: 1232312@sip.domain.com"), []byte("i: 1232312@***.******.***")},
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
		{[]byte("Via: SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_169eac12baa1"), []byte("Via: SIP/2.0/UDP **.***.*.***;branch=z9hG4bKf_169eac12baa1")},
		{[]byte("v: SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_169eac12baa1"), []byte("v: SIP/2.0/UDP **.***.*.***;branch=z9hG4bKf_169eac12baa1")},
		{[]byte("Via: SIP/2.0/TCP 10.101.6.120;maddr=10.10.10.10;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa1"), []byte("Via: SIP/2.0/TCP **.***.*.***;maddr=**.**.**.**;received=*.*.*.*;rport=****;branch=z9hG4bKf_169eac12baa1")},
		{[]byte("v: SIP/2.0/TCP 10.101.6.120;maddr=10.10.10.10;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa1"), []byte("v: SIP/2.0/TCP **.***.*.***;maddr=**.**.**.**;received=*.*.*.*;rport=****;branch=z9hG4bKf_169eac12baa1")},
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
