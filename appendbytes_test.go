package htmlescape_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-htmlescape"
)

func TestAppendBytes(t *testing.T) {

	tests := []struct{
		Bytes []byte
		Slice []byte
		Expected []byte
	}{
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Slice:                                         []byte(""),
			Expected: []byte("Hello world! 🙂 Take a look at this: "),
		},



		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Slice:                                         []byte("This is NOT <b>bold</b>!"),
			Expected: []byte("Hello world! 🙂 Take a look at this: This is NOT &lt;b&gt;bold&lt;/b&gt;!"),
		},



		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Slice:                                         []byte(`"this is a quotation"`),
			Expected: []byte("Hello world! 🙂 Take a look at this: &#34;this is a quotation&#34;"),
		},



		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Slice:                                            []byte("\t3 < 5\r\n3 > 2\x00"),
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x9;3 &lt; 5&#xD;&#xA;3 &gt; 2&#x0;"),
		},
	}

	for testNumber, test := range tests {

		var p []byte = append([]byte(nil), test.Bytes...)
		actual, err := htmlescape.AppendBytes(p, test.Slice)

		if nil != err {
			t.Errorf("For test #%d, did not expect to get an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("SLICE: %q", test.Slice)
			t.Logf("BYTES: %q", test.Bytes)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual resulting value of the append is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("SLICE: %q", test.Slice)
			t.Logf("BYTES: %q", test.Bytes)
			continue
		}
	}
}
