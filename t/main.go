package main

import (
	"fmt"
	"log"
	"time"

	"github.com/marv2097/siprocket"

	"github.com/maxkondr/porta-sip-anonymizer"
)

var (
	rawReq = []byte("2018-10-31T14:38:49.293806Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|981| RECEIVED message from UDP:192.168.64.92:5061 at UDP:192.168.67.224:5060:\n\t" +
		"INVITE sip:twhite@10.101.5.120:5060 SIP/2.0\n\t" +
		"Via: SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_7054e0adfb3_I\n\t" +
		"Via: SIP/2.0/UDP 94.78.45.12;maddr=9.9.9.9;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa170\n\t" +
		"Via: SIP/2.0/UDP 192.168.120.100;branch=z9hG4bKf_1234567\n\t" +
		"Record-Route: <sip:192.168.67.224;lr;ep;pinhole=UDP:192.168.64.50:5060>\n\t" +
		"Route: <sip:192.168.67.224;lr;ob;pinhole=UDP:192.168.64.50:5060;ep>\n\t" +
		"From: \"Andrew Prokop\" <sip:aprokop@10.101.6.120:5060>;tag=35b8d8a74ca0f4e34e0adfa7_F10.101.6.120\n\t" +
		"To: sip:twhite@10.101.5.120:5060\n\t" +
		"Call-ID: f169eac17a017b0a4e0adfa8I@10.101.6.120\n\t" +
		"CSeq: 15 INVITE\n\t" +
		"Max-Forwards: 70\n\t" +
		"Contact: Andres <sip:aprokop@10.101.6.120:5080;transport=udp>\n\t" +
		"Content-Type: application/sdp\n\t" +
		"Content-Length: 306\n\t" +
		"User-Agent: Avaya SIP Softphone\n\t" +
		"P-Asserted-Identity: <sip:123001@192.168.67.224>\n\t" +
		"Remote-Party-ID: <sip:123001@192.168.67.224>;party=calling\n\t" +
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

	rawResp = []byte("2018-10-31T14:38:52.352724Z|edgeproxy[4783]|ab6d9db1-643c74bc@192.168.64.92|IS|995| RECEIVED message from UDP:192.168.64.50:5060 at UDP:192.168.67.224:5060:\n\t" +
		"SIP/2.0 200 OK\r\n" +
		"Via: SIP/2.0/UDP 192.168.2.242:5060;received=22.23.24.25;branch=z9hG4bK5ea22bdd74d079b9;alias;rport=5060\r\n" +
		"To: <sip:JohnSmith@mycompany.com>;tag=aprqu3hicnhaiag03-2s7kdq2000ob4\r\n" +
		"From: sip:HarryJones@mycompany.com;tag=89ddf2f1700666f272fb861443003888\r\n" +
		"CSeq: 57413 REGISTER\r\n" +
		"Call-ID: b5deab6380c4e57fa20486e493c68324\r\n" +
		"Contact: <sip:JohnSmith@192.168.2.242:5060>;expires=192\r\n" +
		"Content-Type: application/sdp\r\n" +
		"Content-Length: 306\r\n" +
		"\r\n" +
		"v=0\r\n" +
		"o=PortaSIP 4530741258397867310 1 IN IP4 217.182.47.207\r\n" +
		"s=s=sip:aprokop@10.101.6.120\r\n" +
		"t=0 0\r\n" +
		"c=IN IP4 10.101.6.120\r\n" +
		"a=msid-semantic:WMS 95f5780e-f8a9-44cd-b3f9-32e329fa7144\r\n" +
		"m=audio 42352 RTP/AVP 0 8 9 18 102 103 101\r\n" +
		"c=IN IP4 217.182.47.207\r\n" +
		"a=rtpmap:101 telephone-event/8000\r\n" +
		"a=rtpmap:102 iLBC/8000\r\n" +
		"a=rtpmap:103 opus/48000/2\r\n" +
		"a=fmtp:101 0-15\r\n" +
		"a=fmtp:102 mode=20\r\n" +
		"a=fmtp:103 maxplaybackrate=16000;maxaveragebitrate=24000;useinbandfec=1;usedtx=1\r\n" +
		"a=fmtp:18 annexb=no\r\n" +
		"a=ptime:20\r\n" +
		"a=sendrecv\r\n" +
		"a=ssrc:1585112755 cname:1Lv6hOmlHPqDqSXM\r\n" +
		"a=ssrc:1585112755 msid:95f5780e-f8a9-44cd-b3f9-32e329fa7144 bed6e71f-e555-494e-b5e9-4bbc0c505c57\r\n" +
		"a=ssrc:1585112755 mslabel:95f5780e-f8a9-44cd-b3f9-32e329fa7144\r\n" +
		"a=ssrc:1585112755 label:bed6e71f-e555-494e-b5e9-4bbc0c505c57\r\n" +
		"m=video 61144 RTP/AVP 108 34 99\r\n" +
		"c=IN IP4 217.182.47.207\r\n" +
		"a=rtpmap:108 VP8/90000\r\n" +
		"a=rtpmap:34 H263/90000\r\n" +
		"a=rtpmap:99 H264/90000\r\n" +
		"a=fmtp:108 max-fr=30;max-fs=3600\r\n" +
		"a=fmtp:34 CIF=1;QCIF=2;SQCIF=2\r\n" +
		"a=sendrecv\r\n" +
		"a=ssrc:1474426549 cname:1Lv6hOmlHPqDqSXM\r\n" +
		"a=ssrc:1474426549 msid:95f5780e-f8a9-44cd-b3f9-32e329fa7144 d11")
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func parseRequestYakut() {
	defer timeTrack(time.Now(), "yakut parse request")
	sipanonymizer.ProcessMessage(rawReq)
}

func parseResponseYakut() {
	defer timeTrack(time.Now(), "yakut parse request")
	sipanonymizer.ProcessMessage(rawResp)
}

func parseRequestSiprocket() {
	defer timeTrack(time.Now(), "siprocket parse request")
	siprocket.Parse(rawReq)
}

func parseResponseSiprocket() {
	defer timeTrack(time.Now(), "siprocket parse response")
	siprocket.Parse(rawResp)
}

func main() {

	// fmt.Println(string(rawReq))
	// msg := sipanonymizer.ProcessMessage(rawReq)
	// fmt.Printf("%s", msg)

	sipanonymizer.ProcessMessage(rawReq)
	fmt.Printf("%s", rawReq)

	// msg := sipanonymizer.ProcessMessage(rawResp)
	// fmt.Printf("%s", msg)

	// parseRequestYakut()
	// parseRequestSiprocket()

	// parseResponseYakut()
	// parseResponseSiprocket()
}
