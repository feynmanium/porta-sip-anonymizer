package sipanonymizer

import (
	"bytes"
	"testing"
	"text/template"
)

var (
	rawReq = []byte("2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| RECEIVED message from UDP:192.168.64.92:5061 at UDP:192.168.67.224:5060:\n\t" +
		"INVITE sip:twhite@10.101.5.120:5060 SIP/2.0\n\t" +
		"Via: SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_7054e0adfb3_I\n\t" +
		"Via: SIP/2.0/UDP 94.78.45.12;maddr=9.9.9.9;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa170\n\t" +
		"Via: SIP/2.0/UDP 192.168.120.100;branch=z9hG4bKf_1234567\n\t" +
		"Record-Route: <sip:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50:5060>\n\t" +
		"Route: <sip:192.168.67.224;lr;ob;pinhole=UDP:192.168.64.50:5060;ep>\n\t" +
		"From: “Andrew Prokop” <sip:aprokop@10.101.6.120:5060>;tag=35b8d8a74ca0f4e34e0adfa7_F10.101.6.120\n\t" +
		"To: sip:twhite@10.101.5.120:5060\n\t" +
		"Call-ID: f169eac17a017b0a4e0adfa8I@10.101.6.120\n\t" +
		"CSeq: 15 INVITE\n\t" +
		"Max-Forwards: 70\n\t" +
		"Contact: Andres <sip:aprokop@10.101.6.120:5080;transport=udp>\n\t" +
		"Content-Type: application/sdp\n\t" +
		"Content-Length: 306\n\t" +
		"User-Agent: Avaya SIP Softphone\n\t" +
		"Supported: replaces\n\t" +
		"X-TEST-HEADER: test\n\t" +
		"\n\t" +
		"v=0\n\t" +
		"o=PortaSIP 4530741258397867310 1 IN IP4 217.182.47.207\n\t" +
		"s=s=sip:aprokop@10.101.6.120\n\t" +
		"t=0 0\n\t" +
		"c=IN IP4 10.101.6.120\n\t" +
		"a=msid-semantic:WMS 95f5780e-f8a9-44cd-b3f9-32e329fa7144\n\t" +
		"m=audio 42352 RTP/AVP 0 8 9 18 102 103 101\n\t" +
		"c=IN IP4 217.182.47.207\n\t" +
		"a=rtpmap:101 telephone-event/8000\n\t" +
		"a=rtpmap:102 iLBC/8000\n\t" +
		"a=rtpmap:103 opus/48000/2\n\t" +
		"a=fmtp:101 0-15\n\t" +
		"a=fmtp:102 mode=20\n\t" +
		"a=fmtp:103 maxplaybackrate=16000;maxaveragebitrate=24000;useinbandfec=1;usedtx=1\n\t" +
		"a=fmtp:18 annexb=no\n\t" +
		"a=ptime:20\n\t" +
		"a=sendrecv\n\t" +
		"a=ssrc:1585112755 cname:1Lv6hOmlHPqDqSXM\n\t" +
		"a=ssrc:1585112755 msid:95f5780e-f8a9-44cd-b3f9-32e329fa7144 bed6e71f-e555-494e-b5e9-4bbc0c505c57\n\t" +
		"a=ssrc:1585112755 mslabel:95f5780e-f8a9-44cd-b3f9-32e329fa7144\n\t" +
		"a=ssrc:1585112755 label:bed6e71f-e555-494e-b5e9-4bbc0c505c57\n\t" +
		"m=video 61144 RTP/AVP 108 34 99\n\t" +
		"c=IN IP4 217.182.47.207\n\t" +
		"a=rtpmap:108 VP8/90000\n\t" +
		"a=rtpmap:34 H263/90000\n\t" +
		"a=rtpmap:99 H264/90000\n\t" +
		"a=fmtp:108 max-fr=30;max-fs=3600\n\t" +
		"a=fmtp:34 CIF=1;QCIF=2;SQCIF=2\n\t" +
		"a=sendrecv\n\t" +
		"a=ssrc:1474426549 cname:1Lv6hOmlHPqDqSXM\n\t" +
		"a=ssrc:1474426549 msid:95f5780e-f8a9-44cd-b3f9-32e329fa7144 d11")
	parsedReq = []byte("2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| RECEIVED message from UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.92:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} at UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}:\n\t" +
		"INVITE sip:twh{{.Mask}}{{.Mask}}e@10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} SIP/2.0\n\t" +
		"Via: SIP/2.0/UDP 10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120;branch=z9hG4bKf_7054e0adfb3_I\n\t" +
		"Via: SIP/2.0/UDP 94.{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.12;maddr=9.{{.Mask}}.{{.Mask}}.9;received=8.{{.Mask}}.{{.Mask}}.8;rport={{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};branch=z9hG4bKf_169eac12baa170\n\t" +
		"Via: SIP/2.0/UDP 192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}{{.Mask}}.100;branch=z9hG4bKf_1234567\n\t" +
		"Record-Route: <sip:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ep;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>\n\t" +
		"Route: <sip:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224;lr;ob;pinhole=UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};ep>\n\t" +
		"From: “{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} {{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} <sip:apr{{.Mask}}{{.Mask}}{{.Mask}}p@10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;tag=35b8d8a74ca0f4e34e0adfa7_F10.101.6.120\n\t" +
		"To: sip:twh{{.Mask}}{{.Mask}}e@10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}\n\t" +
		"Call-ID: f169eac17a017b0a4e0adfa8I@10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120\n\t" +
		"CSeq: 15 INVITE\n\t" +
		"Max-Forwards: 70\n\t" +
		"Contact: And{{.Mask}}{{.Mask}}{{.Mask}} <sip:apr{{.Mask}}{{.Mask}}{{.Mask}}p@10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};transport=udp>\n\t" +
		"Content-Type: application/sdp\n\t" +
		"Content-Length: 306\n\t" +
		"User-Agent: Avaya SIP Softphone\n\t" +
		"Supported: replaces\n\t" +
		"X-TEST-HEADER: test\n\t" +
		"\n\t" +
		"v=0\n\t" +
		"o=PortaSIP 4530741258397867310 1 IN IP4 217.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.207\n\t" +
		"s=s=sip:aprokop@10.101.6.120\n\t" +
		"t=0 0\n\t" +
		"c=IN IP4 10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120\n\t" +
		"a=msid-semantic:WMS 95f5780e-f8a9-44cd-b3f9-32e329fa7144\n\t" +
		"m=audio {{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} RTP/AVP 0 8 9 18 102 103 101\n\t" +
		"c=IN IP4 217.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.207\n\t" +
		"a=rtpmap:101 telephone-event/8000\n\t" +
		"a=rtpmap:102 iLBC/8000\n\t" +
		"a=rtpmap:103 opus/48000/2\n\t" +
		"a=fmtp:101 0-15\n\t" +
		"a=fmtp:102 mode=20\n\t" +
		"a=fmtp:103 maxplaybackrate=16000;maxaveragebitrate=24000;useinbandfec=1;usedtx=1\n\t" +
		"a=fmtp:18 annexb=no\n\t" +
		"a=ptime:20\n\t" +
		"a=sendrecv\n\t" +
		"a=ssrc:1585112755 cname:1Lv6hOmlHPqDqSXM\n\t" +
		"a=ssrc:1585112755 msid:95f5780e-f8a9-44cd-b3f9-32e329fa7144 bed6e71f-e555-494e-b5e9-4bbc0c505c57\n\t" +
		"a=ssrc:1585112755 mslabel:95f5780e-f8a9-44cd-b3f9-32e329fa7144\n\t" +
		"a=ssrc:1585112755 label:bed6e71f-e555-494e-b5e9-4bbc0c505c57\n\t" +
		"m=video {{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} RTP/AVP 108 34 99\n\t" +
		"c=IN IP4 217.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.207\n\t" +
		"a=rtpmap:108 VP8/90000\n\t" +
		"a=rtpmap:34 H263/90000\n\t" +
		"a=rtpmap:99 H264/90000\n\t" +
		"a=fmtp:108 max-fr=30;max-fs=3600\n\t" +
		"a=fmtp:34 CIF=1;QCIF=2;SQCIF=2\n\t" +
		"a=sendrecv\n\t" +
		"a=ssrc:1474426549 cname:1Lv6hOmlHPqDqSXM\n\t" +
		"a=ssrc:1474426549 msid:95f5780e-f8a9-44cd-b3f9-32e329fa7144 d11")
	rawResp = []byte("2018-10-31T14:38:52.352724Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|995| RECEIVED message from UDP:192.168.64.50:5060 at UDP:192.168.67.224:5060:\n\t" +
		"SIP/2.0 200 OK\n\t" +
		"Via: SIP/2.0/UDP 192.168.2.242:5060;received=22.23.24.25;branch=z9hG4bK5ea22bdd74d079b9;alias;rport=5060\n\t" +
		"To: <sip:JohnSmith@mycompany.com>;tag=aprqu3hicnhaiag03-2s7kdq2000ob4\n\t" +
		"From: sip:HarryJones@mycompany.com;tag=89ddf2f1700666f272fb861443003888\n\t" +
		"CSeq: 57413 REGISTER\n\t" +
		"Call-ID: b5deab6380c4e57fa20486e493c68324\n\t" +
		"Contact: <sip:JohnSmith@192.168.2.242:5060>;expires=192\n\t" +
		"Content-Type: application/sdp\n\t" +
		"Content-Length: 306\n\t" +
		"\n\t" +
		"v=0\n\t" +
		"o=PortaSIP 4530741258397867310 1 IN IP4 217.182.47.207\n\t" +
		"s=s=sip:aprokop@10.101.6.120\n\t" +
		"t=0 0\n\t" +
		"c=IN IP4 10.101.6.120\n\t" +
		"a=msid-semantic:WMS 95f5780e-f8a9-44cd-b3f9-32e329fa7144\n\t" +
		"m=audio 42352 RTP/AVP 0 8 9 18 102 103 101\n\t" +
		"c=IN IP4 217.182.47.207\n\t" +
		"a=rtpmap:101 telephone-event/8000\n\t" +
		"a=rtpmap:102 iLBC/8000\n\t" +
		"a=rtpmap:103 opus/48000/2\n\t" +
		"a=fmtp:101 0-15\n\t" +
		"a=fmtp:102 mode=20\n\t" +
		"a=fmtp:103 maxplaybackrate=16000;maxaveragebitrate=24000;useinbandfec=1;usedtx=1\n\t" +
		"a=fmtp:18 annexb=no\n\t" +
		"a=ptime:20\n\t" +
		"a=sendrecv\n\t" +
		"a=ssrc:1585112755 cname:1Lv6hOmlHPqDqSXM\n\t" +
		"a=ssrc:1585112755 msid:95f5780e-f8a9-44cd-b3f9-32e329fa7144 bed6e71f-e555-494e-b5e9-4bbc0c505c57\n\t" +
		"a=ssrc:1585112755 mslabel:95f5780e-f8a9-44cd-b3f9-32e329fa7144\n\t" +
		"a=ssrc:1585112755 label:bed6e71f-e555-494e-b5e9-4bbc0c505c57\n\t" +
		"m=video 61144 RTP/AVP 108 34 99\n\t" +
		"c=IN IP4 217.182.47.207\n\t" +
		"a=rtpmap:108 VP8/90000\n\t" +
		"a=rtpmap:34 H263/90000\n\t" +
		"a=rtpmap:99 H264/90000\n\t" +
		"a=fmtp:108 max-fr=30;max-fs=3600\n\t" +
		"a=fmtp:34 CIF=1;QCIF=2;SQCIF=2\n\t" +
		"a=sendrecv\n\t" +
		"a=ssrc:1474426549 cname:1Lv6hOmlHPqDqSXM\n\t" +
		"a=ssrc:1474426549 msid:95f5780e-f8a9-44cd-b3f9-32e329fa7144 d11")
	parsedResp = []byte("2018-10-31T14:38:52.352724Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|995| RECEIVED message from UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.50:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} at UDP:192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.224:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}:\n\t" +
		"SIP/2.0 200 OK\n\t" +
		"Via: SIP/2.0/UDP 192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.242:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}};received=22.{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.25;branch=z9hG4bK5ea22bdd74d079b9;alias;rport={{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}\n\t" +
		"To: <sip:Joh{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}h@myc{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com>;tag=aprqu3hicnhaiag03-2s7kdq2000ob4\n\t" +
		"From: sip:Har{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}s@myc{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}.com;tag=89ddf2f1700666f272fb861443003888\n\t" +
		"CSeq: 57413 REGISTER\n\t" +
		"Call-ID: b5deab6380c4e57fa20486e493c68324\n\t" +
		"Contact: <sip:Joh{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}h@192.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.242:{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}>;expires=192\n\t" +
		"Content-Type: application/sdp\n\t" +
		"Content-Length: 306\n\t" +
		"\n\t" +
		"v=0\n\t" +
		"o=PortaSIP 4530741258397867310 1 IN IP4 217.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.207\n\t" +
		"s=s=sip:aprokop@10.101.6.120\n\t" +
		"t=0 0\n\t" +
		"c=IN IP4 10.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}.120\n\t" +
		"a=msid-semantic:WMS 95f5780e-f8a9-44cd-b3f9-32e329fa7144\n\t" +
		"m=audio {{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} RTP/AVP 0 8 9 18 102 103 101\n\t" +
		"c=IN IP4 217.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.207\n\t" +
		"a=rtpmap:101 telephone-event/8000\n\t" +
		"a=rtpmap:102 iLBC/8000\n\t" +
		"a=rtpmap:103 opus/48000/2\n\t" +
		"a=fmtp:101 0-15\n\t" +
		"a=fmtp:102 mode=20\n\t" +
		"a=fmtp:103 maxplaybackrate=16000;maxaveragebitrate=24000;useinbandfec=1;usedtx=1\n\t" +
		"a=fmtp:18 annexb=no\n\t" +
		"a=ptime:20\n\t" +
		"a=sendrecv\n\t" +
		"a=ssrc:1585112755 cname:1Lv6hOmlHPqDqSXM\n\t" +
		"a=ssrc:1585112755 msid:95f5780e-f8a9-44cd-b3f9-32e329fa7144 bed6e71f-e555-494e-b5e9-4bbc0c505c57\n\t" +
		"a=ssrc:1585112755 mslabel:95f5780e-f8a9-44cd-b3f9-32e329fa7144\n\t" +
		"a=ssrc:1585112755 label:bed6e71f-e555-494e-b5e9-4bbc0c505c57\n\t" +
		"m=video {{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}}{{.Mask}} RTP/AVP 108 34 99\n\t" +
		"c=IN IP4 217.{{.Mask}}{{.Mask}}{{.Mask}}.{{.Mask}}{{.Mask}}.207\n\t" +
		"a=rtpmap:108 VP8/90000\n\t" +
		"a=rtpmap:34 H263/90000\n\t" +
		"a=rtpmap:99 H264/90000\n\t" +
		"a=fmtp:108 max-fr=30;max-fs=3600\n\t" +
		"a=fmtp:34 CIF=1;QCIF=2;SQCIF=2\n\t" +
		"a=sendrecv\n\t" +
		"a=ssrc:1474426549 cname:1Lv6hOmlHPqDqSXM\n\t" +
		"a=ssrc:1474426549 msid:95f5780e-f8a9-44cd-b3f9-32e329fa7144 d11")
)

func TestProcessSipEntryRequest(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{rawReq, string(parsedReq)},
	}

	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		parse(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of ProcessSipEntryRequest is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func TestProcessSipEntryResponse(t *testing.T) {
	tables := []struct {
		src  []byte
		want string
	}{
		{rawResp, string(parsedResp)},
	}

	ts := getTestingMaskStruct()
	for _, table := range tables {
		line := make([]byte, len(table.src))
		copy(line, table.src)
		parse(line)
		templ := template.Must(template.New("want").Parse(table.want))
		var buf bytes.Buffer
		templ.Execute(&buf, ts)
		if string(line) != string(buf.Bytes()) {
			t.Errorf("Result of ProcessSipEntryResponse is incorrect:\n src  %s\n want %s\n got  %s",
				table.src, string(buf.Bytes()), line)
		}
	}
}

func BenchmarkProcessSipEntryRequest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parse(rawReq)
	}
}

func BenchmarkProcessSipEntryResponse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		parse(rawResp)
	}
}
