package htmlescape

import (
	"io"
)

func WriteBytes(writer io.Writer, bytes []byte) (n int, err error) {
	var buffer [bufferSize]byte
	var p []byte = buffer[0:0]

	p, err = AppendBytes(p, bytes)
	if nil != err {
		return 0, err
	}

	return writer.Write(p)
}

func WriteRune(writer io.Writer, r rune) (n int, err error) {
	var buffer [bufferSize]byte
	var p []byte = buffer[0:0]

	p = AppendRune(p,r)

	return writer.Write(p)
}

func WriteString(writer io.Writer, str string) (n int, err error) {
	var buffer [bufferSize]byte
	var p []byte = buffer[0:0]

	p, err = AppendString(p, str)
	if nil != err {
		return 0, err
	}

	return writer.Write(p)
}
