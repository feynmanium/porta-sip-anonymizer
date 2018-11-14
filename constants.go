package sipanonymizer

const (
	FieldBase = iota
	FieldName
	FieldNameQ
	FieldUser
	FieldHost
	FieldHostIP
	FieldPort
	FieldText
	FieldPreserveFirstOctet  // preserve first octet in IP like 192.x.x.x
	FieldPreserveLastOctet   // preserve last octet in IP like x.x.x.10
	FieldPreserveFirstDomain // preserve first part of domain like gooxxx.xxx
	FieldPreserveLastDomain  // preserve first part of domain like xxxxxx.com
)

const maskChar byte = '*'
const atBytes = '@'

var lineSep = []byte("\n\t")
var spaceBytes = []byte(" ")
var udpBytes = []byte("UDP:")
var tcpBytes = []byte("TCP:")
var tlsBytes = []byte("TLS:")
var dotBytes = []byte(".")
var sip20Bytes = []byte("SIP/2.0")
var pinholeBytes = []byte("pinhole=")
var sipBytes = []byte("sip:")
var sipsBytes = []byte("sips:")
var telBytes = []byte("tel:")
var sipCapBytes = []byte("SIP")
var rportBytes = []byte("rport=")
var maddrBytes = []byte("maddr=")
var receivedBytes = []byte("received=")
var viaBytes = []byte("via")
var viaCapBytes = []byte("Via")
var fromBytes = []byte("from")
var fromCapBytes = []byte("From")
var toBytes = []byte("to")
var toCapBytes = []byte("To")
var contactBytes = []byte("contact")
var contactCapBytes = []byte("Contact")
var routeBytes = []byte("route")
var routeCapBytes = []byte("Route")
var recordRouteBytes = []byte("record-route")
var recordRouteCapBytes = []byte("Record-Route")
var rpidBytes = []byte("remote-party-id")
var rpidCapBytes = []byte("Remote-Party-Id")
var callIDBytes = []byte("call-id")
var callIDCapBytes = []byte("Call-ID")
var paiBytes = []byte("p-asserted-identity")
var paiCapBytes = []byte("P-Asserted-Identity")

var userEnd = string("<\"'@")
var headerSep = string(":=")

// var hostEnd = string(":;> ")
