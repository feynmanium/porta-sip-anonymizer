package sipanonymizer

import "testing"

func TestProcessSipURL(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("<sip:87.252.61.202>"), []byte("<sip:87.***.**.202>")},
		{[]byte("<sips:87.252.61.202>"), []byte("<sips:87.***.**.202>")},
		{[]byte("<tel:87.252.61.202>"), []byte("<tel:87.***.**.202>")},
		{[]byte("<sip:87.252.61.202:5060>"), []byte("<sip:87.***.**.202:****>")},
		{[]byte("<sips:87.252.61.202:5060>"), []byte("<sips:87.***.**.202:****>")},
		{[]byte("<tel:87.252.61.202:5060>"), []byte("<tel:87.***.**.202:****>")},
		{[]byte("<sip:anonymous@anonymous.invalid>"), []byte("<sip:ano*****s@ano******.invalid>")},
		{[]byte("<sips:anonymous@anonymous.invalid>"), []byte("<sips:ano*****s@ano******.invalid>")},
		{[]byte("<tel:anonymous@anonymous.invalid>"), []byte("<tel:ano*****s@ano******.invalid>")},
		{[]byte("<sip:anonymous@anonymous.invalid:5060>"), []byte("<sip:ano*****s@ano******.invalid:****>")},
		{[]byte("<sips:anonymous@anonymous.invalid:5060>"), []byte("<sips:ano*****s@ano******.invalid:****>")},
		{[]byte("<tel:anonymous@anonymous.invalid:5060>"), []byte("<tel:ano*****s@ano******.invalid:****>")},
		{[]byte("<sip:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50>"), []byte("<sip:192.***.**.224;lr;ep;pinhole=UDP:192.***.**.50>")},
		{[]byte("<sips:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50>"), []byte("<sips:192.***.**.224;lr;ep;pinhole=UDP:192.***.**.50>")},
		{[]byte("<tel:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50>"), []byte("<tel:192.***.**.224;lr;ep;pinhole=UDP:192.***.**.50>")},
		{[]byte("<sip:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50:5060>"), []byte("<sip:192.***.**.224;lr;ep;pinhole=UDP:192.***.**.50:****>")},
		{[]byte("<sips:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50:5060>"), []byte("<sips:192.***.**.224;lr;ep;pinhole=UDP:192.***.**.50:****>")},
		{[]byte("<tel:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50:5060>"), []byte("<tel:192.***.**.224;lr;ep;pinhole=UDP:192.***.**.50:****>")},
		{[]byte("<sip:john@87.252.61.202>"), []byte("<sip:j***@87.***.**.202>")},
		{[]byte("<sips:john@87.252.61.202>"), []byte("<sips:j***@87.***.**.202>")},
		{[]byte("<tel:john@87.252.61.202>"), []byte("<tel:j***@87.***.**.202>")},
		{[]byte("<sip:john@87.252.61.202:5060>"), []byte("<sip:j***@87.***.**.202:****>")},
		{[]byte("<sips:john@87.252.61.202:5060>"), []byte("<sips:j***@87.***.**.202:****>")},
		{[]byte("<tel:john@87.252.61.202:5060>"), []byte("<tel:j***@87.***.**.202:****>")},
		{[]byte("<sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("<sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("<sips:john@87.252.61.202>;tag=bvbvfhehj"), []byte("<sips:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("<tel:john@87.252.61.202>;tag=bvbvfhehj"), []byte("<tel:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("<sip:john@87.252.61.202:5060>;tag=bvbvfhehj"), []byte("<sip:j***@87.***.**.202:****>;tag=bvbvfhehj")},
		{[]byte("<sips:john@87.252.61.202:5060>;tag=bvbvfhehj"), []byte("<sips:j***@87.***.**.202:****>;tag=bvbvfhehj")},
		{[]byte("<tel:john@87.252.61.202:5060>;tag=bvbvfhehj"), []byte("<tel:j***@87.***.**.202:****>;tag=bvbvfhehj")},
		{[]byte("John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("Joh* <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("John <sips:john@87.252.61.202>;tag=bvbvfhehj"), []byte("Joh* <sips:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("John <tel:john@87.252.61.202>;tag=bvbvfhehj"), []byte("Joh* <tel:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("John <sip:john@87.252.61.202:5060>;tag=bvbvfhehj"), []byte("Joh* <sip:j***@87.***.**.202:****>;tag=bvbvfhehj")},
		{[]byte("John <sips:john@87.252.61.202:5060>;tag=bvbvfhehj"), []byte("Joh* <sips:j***@87.***.**.202:****>;tag=bvbvfhehj")},
		{[]byte("John <tel:john@87.252.61.202:5060>;tag=bvbvfhehj"), []byte("Joh* <tel:j***@87.***.**.202:****>;tag=bvbvfhehj")},
		{[]byte("\"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("\"Joh* **e\" <sip:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("\"John Doe\" <sips:john@87.252.61.202>;tag=bvbvfhehj"), []byte("\"Joh* **e\" <sips:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("\"John Doe\" <tel:john@87.252.61.202>;tag=bvbvfhehj"), []byte("\"Joh* **e\" <tel:j***@87.***.**.202>;tag=bvbvfhehj")},
		{[]byte("\"John Doe\" <sip:john@87.252.61.202:5060>;tag=bvbvfhehj"), []byte("\"Joh* **e\" <sip:j***@87.***.**.202:****>;tag=bvbvfhehj")},
		{[]byte("\"John Doe\" <sips:john@87.252.61.202:5060>;tag=bvbvfhehj"), []byte("\"Joh* **e\" <sips:j***@87.***.**.202:****>;tag=bvbvfhehj")},
		{[]byte("\"John Doe\" <tel:john@87.252.61.202:5060>;tag=bvbvfhehj"), []byte("\"Joh* **e\" <tel:j***@87.***.**.202:****>;tag=bvbvfhehj")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSipURL(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of TestProcessSipURL is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}

func BenchmarkProcessSipURL(b *testing.B) {
	for n := 0; n < b.N; n++ {
		processSipURL([]byte("\"John Doe\" <sip:john@87.252.61.202:5060>;tag=bvbvfhehj"))
	}
}
