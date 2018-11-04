package sipanonymizer

import "testing"

func TestProcessSipURL(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("\"John Doe\" <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("\"********\" <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("John <sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("**** <sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("<sip:john@87.252.61.202>;tag=bvbvfhehj"), []byte("<sip:****@**.***.**.***>;tag=bvbvfhehj")},
		{[]byte("<sip:john@87.252.61.202>"), []byte("<sip:****@**.***.**.***>")},
		{[]byte("<sip:87.252.61.202>"), []byte("<sip:**.***.**.***>")},
		{[]byte("<sip:anonymous@anonymous.invalid>"), []byte("<sip:*********@*********.*******>")},
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
