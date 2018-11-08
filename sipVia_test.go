package sipanonymizer

import (
	"testing"
)

func TestProcessSipVia(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_169eac12baa1"), []byte("SIP/2.0/UDP 10.***.*.120;branch=z9hG4bKf_169eac12baa1")},
		{[]byte("SIP/2.0/TCP 10.101.6.120;maddr=10.10.10.10;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa1"), []byte("SIP/2.0/TCP 10.***.*.120;maddr=10.**.**.10;received=8.*.*.8;rport=****;branch=z9hG4bKf_169eac12baa1")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSipVia(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of TestProcessSipVia is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}

func BenchmarkProcessSipVia(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processSipVia([]byte("SIP/2.0/TCP 10.101.6.120;maddr=10.10.10.10;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa1"))
	}
}
