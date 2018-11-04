package main

import (
	"log"
	"time"

	"github.com/marv2097/siprocket"

	"github.com/maxkondr/porta-sip-anonymizer"
)

var (
	rawReq = []byte("INVITE sip:twhite@10.101.5.120:5060 SIP/2.0\r\n" +
		"Via: SIP/2.0/UDP 10.101.6.120;branch=z9hG4bKf_7054e0adfb3_I\r\n" +
		"Via: SIP/2.0/UDP 94.78.45.12;maddr=9.9.9.9;received=8.8.8.8;rport=8090;branch=z9hG4bKf_169eac12baa170\r\n" +
		"Via: SIP/2.0/UDP 192.168.120.100;branch=z9hG4bKf_1234567\r\n" +
		"From: “Andrew Prokop” <sip:aprokop@10.101.6.120:5060>;tag=35b8d8a74ca0f4e34e0adfa7_F10.101.6.120\r\n" +
		"To: sip:twhite@10.101.5.120:5060\r\n" +
		"Call-ID: f169eac17a017b0a4e0adfa8I@10.101.6.120\r\n" +
		"CSeq: 15 INVITE\r\n" +
		"Content-Length: 306\r\n" +
		"Max-Forwards: 70\r\n" +
		"Contact: Andres <sip:aprokop@10.101.6.120:5080;transport=udp>\r\n" +
		"Content-Type: application/sdp\r\n" +
		"User-Agent: Avaya SIP Softphone\r\n" +
		"Supported: replaces\r\n" +
		"X-TEST-HEADER: test\r\n" +
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

func parseRequestSiprocket() {
	defer timeTrack(time.Now(), "siprocket parse request")
	siprocket.Parse(rawReq)
}

func main() {

	// msg := sipanonymizer.ProcessMessage(rawReq)
	// fmt.Printf("%s", msg.GetString())

	parseRequestYakut()
	parseRequestSiprocket()
}
