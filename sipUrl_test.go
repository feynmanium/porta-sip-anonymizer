package sipanonymizer

import (
	"bytes"
	"testing"
	"text/template"
)

func TestProcessSipURL(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{[]byte("<sip:87.252.61.202>"), "<sip:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("<sips:87.252.61.202>"), "<sips:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("<tel:87.252.61.202>"), "<tel:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("<sip:87.252.61.202:5060>"), "<sip:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<sips:87.252.61.202:5060>"), "<sips:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<tel:87.252.61.202:5060>"), "<tel:87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<sip:anonymous@anonymous.invalid>"), "<sip:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid>"},
		{[]byte("<sips:anonymous@anonymous.invalid>"), "<sips:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid>"},
		{[]byte("<tel:anonymous@anonymous.invalid>"), "<tel:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid>"},
		{[]byte("<sip:anonymous@anonymous.invalid:5060>"), "<sip:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<sips:anonymous@anonymous.invalid:5060>"), "<sips:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<tel:anonymous@anonymous.invalid:5060>"), "<tel:ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@ano{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.invalid:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<sip:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50>"), "<sip:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ep;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50>"},
		{[]byte("<sips:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50>"), "<sips:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ep;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50>"},
		{[]byte("<tel:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50>"), "<tel:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ep;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50>"},
		{[]byte("<sip:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50:5060>"), "<sip:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ep;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<sips:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50:5060>"), "<sips:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ep;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<tel:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50:5060>"), "<tel:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ep;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<sip:john@87.252.61.202>"), "<sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("<sips:john@87.252.61.202>"), "<sips:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("<tel:john@87.252.61.202>"), "<tel:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>"},
		{[]byte("<sip:john@87.252.61.202:5060>"), "<sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<sips:john@87.252.61.202:5060>"), "<sips:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<tel:john@87.252.61.202:5060>"), "<tel:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>"},
		{[]byte("<sip:john@87.252.61.202>;tag=bvbvfhehj"), "<sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("<sips:john@87.252.61.202>;tag=bvbvfhehj"), "<sips:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("<tel:john@87.252.61.202>;tag=bvbvfhehj"), "<tel:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("<sip:john@87.252.61.202:5060>;tag=bvbvfhehj"), "<sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=bvbvfhehj"},
		{[]byte("<sips:john@87.252.61.202:5060>;tag=bvbvfhehj"), "<sips:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=bvbvfhehj"},
		{[]byte("<tel:john@87.252.61.202:5060>;tag=bvbvfhehj"), "<tel:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=bvbvfhehj"},
		{[]byte("John <sip:john@87.252.61.202>;tag=bvbvfhehj"), "Joh{{.Mask}} <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("John <sips:john@87.252.61.202>;tag=bvbvfhehj"), "Joh{{.Mask}} <sips:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("John <tel:john@87.252.61.202>;tag=bvbvfhehj"), "Joh{{.Mask}} <tel:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("John <sip:john@87.252.61.202:5060>;tag=bvbvfhehj"), "Joh{{.Mask}} <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=bvbvfhehj"},
		{[]byte("John <sips:john@87.252.61.202:5060>;tag=bvbvfhehj"), "Joh{{.Mask}} <sips:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=bvbvfhehj"},
		{[]byte("John <tel:john@87.252.61.202:5060>;tag=bvbvfhehj"), "Joh{{.Mask}} <tel:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=bvbvfhehj"},
		{[]byte("\"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), "\"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("\"John Doe\" <sips:john@87.252.61.202>;tag=bvbvfhehj"), "\"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <sips:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("\"John Doe\" <tel:john@87.252.61.202>;tag=bvbvfhehj"), "\"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <tel:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202>;tag=bvbvfhehj"},
		{[]byte("\"John Doe\" <sip:john@87.252.61.202:5060>;tag=bvbvfhehj"), "\"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <sip:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=bvbvfhehj"},
		{[]byte("\"John Doe\" <sips:john@87.252.61.202:5060>;tag=bvbvfhehj"), "\"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <sips:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=bvbvfhehj"},
		{[]byte("\"John Doe\" <tel:john@87.252.61.202:5060>;tag=bvbvfhehj"), "\"Joh{{.Mask}} {{.Mask}}{{.Mask}}e\" <tel:j{{.Mask}}{{.Mask}}{{.Mask}}@87.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.202:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=bvbvfhehj"},
	}
	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSipURL(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of TestProcessSipURL is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func BenchmarkProcessSipURL(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processSipURL([]byte("\"John Doe\" <sip:john@87.252.61.202:5060>;tag=bvbvfhehj"))
	}
}
