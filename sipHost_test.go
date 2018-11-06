package sipanonymizer

import (
	"testing"
)

func TestProcessHost(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("192.168.192.100"), []byte("192.***.***.100")},
		{[]byte("192.168.192.100 "), []byte("192.***.***.100 ")},
		{[]byte("192.168.192.100 123"), []byte("192.***.***.100 123")},
		{[]byte("192.168.192.100;"), []byte("192.***.***.100;")},
		{[]byte("192.168.192.100;rport=123;received=8.8.8.8"), []byte("192.***.***.100;rport=123;received=8.8.8.8")},
		{[]byte("192.168.192.100>"), []byte("192.***.***.100>")},
		{[]byte("192.168.192.100:"), []byte("192.***.***.100:")},
		{[]byte("192.168.192.100:5060"), []byte("192.***.***.100:****")},
		{[]byte("192.168.192.100:5060 "), []byte("192.***.***.100:**** ")},
		{[]byte("192.168.192.100:5060 123"), []byte("192.***.***.100:**** 123")},
		{[]byte("192.168.192.100:5060;"), []byte("192.***.***.100:****;")},
		{[]byte("192.168.192.100:5060>"), []byte("192.***.***.100:****>")},
		{[]byte("192.168.192.100:5060:"), []byte("192.***.***.100:****:")},
		{[]byte("domain.com"), []byte("dom***.com")},
		{[]byte("domain.com "), []byte("dom***.com ")},
		{[]byte("domain.com 123"), []byte("dom***.com 123")},
		{[]byte("domain.com;"), []byte("dom***.com;")},
		{[]byte("domain.com;rport=123"), []byte("dom***.com;rport=123")},
		{[]byte("domain.com>"), []byte("dom***.com>")},
		{[]byte("domain.com:"), []byte("dom***.com:")},
		{[]byte("domain.com:5060"), []byte("dom***.com:****")},
		{[]byte("domain.com:5060 "), []byte("dom***.com:**** ")},
		{[]byte("domain.com:5060 123"), []byte("dom***.com:**** 123")},
		{[]byte("domain.com:5060;"), []byte("dom***.com:****;")},
		{[]byte("domain.com:5060>"), []byte("dom***.com:****>")},
		{[]byte("domain.com:5060:"), []byte("dom***.com:****:")},
		{[]byte("sip.domain.com:5060"), []byte("sip.******.com:****")},
		{[]byte("sip.domain.com:5060 "), []byte("sip.******.com:**** ")},
		{[]byte("sip.domain.com:5060 123"), []byte("sip.******.com:**** 123")},
		{[]byte("sip.domain.com:5060;"), []byte("sip.******.com:****;")},
		{[]byte("sip.domain.com:5060>"), []byte("sip.******.com:****>")},
		{[]byte("sip.domain.com:5060:"), []byte("sip.******.com:****:")},
		{[]byte("a.very.complex-domain.co.uk:8080"), []byte("a.****.*******-******.**.uk:****")},
		{[]byte("a.very.complex-domain.co.uk:8080 "), []byte("a.****.*******-******.**.uk:**** ")},
		{[]byte("a.very.complex-domain.co.uk:8080 123"), []byte("a.****.*******-******.**.uk:**** 123")},
		{[]byte("a.very.complex-domain.co.uk:8080;"), []byte("a.****.*******-******.**.uk:****;")},
		{[]byte("a.very.complex-domain.co.uk:8080>"), []byte("a.****.*******-******.**.uk:****>")},
		{[]byte("a.very.complex-domain.co.uk:8080:"), []byte("a.****.*******-******.**.uk:****:")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processHost(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of processHost is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
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
