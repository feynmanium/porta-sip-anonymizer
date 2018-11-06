package sipanonymizer

import "testing"

func TestProcessSipCallID(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("vjnejivnreivujreiuvjnie"), []byte("vjnejivnreivujreiuvjnie")},
		{[]byte("1232312@192.168.1.10"), []byte("1232312@192.***.*.10")},
		{[]byte("1232312@sip.domain.com"), []byte("1232312@sip.******.com")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSipCallID(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of ProcessSipRequestLine is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}

func BenchmarkProcessSipCallID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processSipCallID([]byte("1232312@sip.domain.com"))
	}
}
