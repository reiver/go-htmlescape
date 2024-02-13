package htmlescape

func String(str string) string {
	var buffer [bufferSize]byte
	var p []byte = buffer[0:0]

	p, _ = AppendString(p, str)

	return string(p)
}
