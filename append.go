package htmlescape

import (
	"unicode/utf8"
)

func AppendBytes(p []byte, bytes []byte) ([]byte, error) {
	for 0 < len(bytes) {
		r, size := utf8.DecodeRune(bytes)
		if utf8.RuneError == r && 1 == size {
			return p, errNotUTF8
		}
		if size <= 0 {
			return p, errInternalError
		}
		bytes = bytes[size:]

		p = AppendRune(p, r)
	}

	return p, nil
}

func AppendRune(p []byte, r rune) []byte {

	escapedRune, wasEscaped := safeRune(r)
	if wasEscaped {
		return append(p, escapedRune...)
	}

	return append(p, string(r)...)
}

func AppendString(p []byte, str string) ([]byte, error) {
	for 0 < len(str) {
		r, size := utf8.DecodeRuneInString(str)
		if utf8.RuneError == r && 1 == size {
			return p, errNotUTF8
		}
		if size <= 0 {
			return p, errInternalError
		}
		str = str[size:]

		p = AppendRune(p, r)
	}

	return p, nil
}
