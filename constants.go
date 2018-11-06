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
