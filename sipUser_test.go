package sipanonymizer

import (
	"bytes"
	"testing"
	"text/template"
)

func TestProcessUser(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("1"), "{{.Mask}}"},
		{[]byte("12"), "1{{.Mask}}"},
		{[]byte("123"), "1{{.Mask}}{{.Mask}}"},
		{[]byte("1234"), "1{{.Mask}}{{.Mask}}{{.Mask}}"},
		{[]byte("12345"), "123{{.Mask}}5"},
		{[]byte("123456"), "123{{.Mask}}{{.Mask}}6"},
		{[]byte("1234567"), "123{{.Mask}}{{.Mask}}{{.Mask}}7"},
		{[]byte("12345678"), "123{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}8"},
		{[]byte("123456789"), "123{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}9"},
		{[]byte("1234567890"), "123{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}0"},
		{[]byte("1234567890@1.1.1.1"), "123{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}0@1.1.1.1"},
		{[]byte("John"), "J{{.Mask}}{{.Mask}}{{.Mask}}"},
		{[]byte("John@1.1.1.1"), "J{{.Mask}}{{.Mask}}{{.Mask}}@1.1.1.1"},
		{[]byte("John <sip:1.1.1.1>"), "Joh{{.Mask}} <sip:1.1.1.1>"},
		{[]byte("John Doe"), "Joh{{.Mask}} {{.Mask}}{{.Mask}}e"},
		{[]byte("\"John Doe\""), "\"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\""},
	}
	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processUser(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of processUser is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func BenchmarkProcessUser(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processUser([]byte("\"John Doe\""))
	}
}
