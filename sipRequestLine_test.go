package sipanonymizer

import "testing"

func TestProcessSipRequestLine(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("INVITE sip:abc@87.252.61.202;user=phone SIP/2.0"), []byte("INVITE sip:***@**.***.**.***;user=phone SIP/2.0")},
		{[]byte("INVITE sips:abc@87.252.61.202;user=phone SIP/2.0"), []byte("INVITE sips:***@**.***.**.***;user=phone SIP/2.0")},
		{[]byte("INVITE tel:abc@87.252.61.202;user=phone SIP/2.0"), []byte("INVITE tel:***@**.***.**.***;user=phone SIP/2.0")},
		{[]byte("INVITE sip:abc@domain.com SIP/2.0"), []byte("INVITE sip:***@******.*** SIP/2.0")},
		{[]byte("REGISTER sips:ss2.biloxi.example.com SIP/2.0"), []byte("REGISTER sips:***.******.*******.*** SIP/2.0")},
		{[]byte("BYE sip:123456@host.domain.com SIP/2.0"), []byte("BYE sip:******@****.******.*** SIP/2.0")},
		{[]byte("BYE sip:abc@8.8.8.8 SIP/2.0"), []byte("BYE sip:***@*.*.*.* SIP/2.0")},
		{[]byte("ACK sip:abc@domain.com SIP/2.0"), []byte("ACK sip:***@******.*** SIP/2.0")},
		{[]byte("CANCEL sip:abc@domain.com SIP/2.0"), []byte("CANCEL sip:***@******.*** SIP/2.0")},
		{[]byte("NOTIFY sip:abc@domain.com SIP/2.0"), []byte("NOTIFY sip:***@******.*** SIP/2.0")},
		{[]byte("MESSAGE sip:abc@domain.com SIP/2.0"), []byte("MESSAGE sip:***@******.*** SIP/2.0")},
		{[]byte("SIP/2.0 200 OK"), []byte("SIP/2.0 200 OK")},
		{[]byte("SIP/2.0 401 Unauthorized"), []byte("SIP/2.0 401 Unauthorized")},
		{[]byte("SIP/2.0 100 Trying"), []byte("SIP/2.0 100 Trying")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSipRequestLine(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of ProcessSipRequestLine is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}
