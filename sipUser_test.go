package sipanonymizer

import (
	"testing"
)

func TestProcessUser(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("1"), []byte("*")},
		{[]byte("12"), []byte("1*")},
		{[]byte("123"), []byte("1**")},
		{[]byte("1234"), []byte("1***")},
		{[]byte("12345"), []byte("123*5")},
		{[]byte("123456"), []byte("123**6")},
		{[]byte("1234567"), []byte("123***7")},
		{[]byte("12345678"), []byte("123****8")},
		{[]byte("123456789"), []byte("123*****9")},
		{[]byte("1234567890"), []byte("123******0")},
		{[]byte("1234567890@1.1.1.1"), []byte("123******0@1.1.1.1")},
		{[]byte("John"), []byte("J***")},
		{[]byte("John@1.1.1.1"), []byte("J***@1.1.1.1")},
		{[]byte("John <sip:1.1.1.1>"), []byte("Joh* <sip:1.1.1.1>")},
		{[]byte("John Doe"), []byte("Joh* **e")},
		{[]byte("\"John Doe\""), []byte("\"Joh* **e\"")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processUser(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of processUser is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}

func BenchmarkProcessUser(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processUser([]byte("\"John Doe\""))
	}
}
