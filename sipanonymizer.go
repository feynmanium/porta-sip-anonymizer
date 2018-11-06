package sipanonymizer

// ProcessMessage reads in messages from bytes
func ProcessMessage(v []byte) string {
	return string(parse(v))
}
