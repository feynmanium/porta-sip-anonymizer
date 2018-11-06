package sipanonymizer

import "testing"

func TestProcessSdpConnection(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("c=IN IP4 10.101.6.120"), []byte("c=IN IP4 10.***.*.120")},
		{[]byte("c=IN IP4 sip.domain.com"), []byte("c=IN IP4 sip.******.com")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSdpConnection(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of processSdpConnection is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}

func TestProcessSdpOriginator(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("o=PortaSIP 4530741258397867310 1 IN IP4 217.182.47.207"), []byte("o=PortaSIP 4530741258397867310 1 IN IP4 217.***.**.207")},
		{[]byte("o=PortaSIP 4530741258397867310 1 IN IP4 sip.domain.com"), []byte("o=PortaSIP 4530741258397867310 1 IN IP4 sip.******.com")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSdpOriginator(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of processSdpOriginator is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}

func TestProcessSdpMedia(t *testing.T) {
	tables := []struct {
		src  []byte
		want []byte
	}{
		{[]byte("m=audio 42352 RTP/AVP 0 8 9 18 102 103 101"), []byte("m=audio ***** RTP/AVP 0 8 9 18 102 103 101")},
	}
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		processSdpMedia(line)
		if string(line) != string(table.want) {
			t.Errorf("Result of processSdpMedia is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, table.want, line)
		}
	}
}
