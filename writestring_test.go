package htmlescape_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-htmlescape"
)

func TestWriteString(t *testing.T) {

	tests := []struct{
		String string
		Expected []byte
	}{
		{
			String:          "This is NOT <b>bold</b>!",
			Expected: []byte("This is NOT &lt;b&gt;bold&lt;/b&gt;!"),
		},



		{
			String:              `"this is a quotation"`,
			Expected: []byte("&#34;this is a quotation&#34;"),
		},



		{
			String:             "\t3 < 5\r\n3 > 2\x00",
			Expected: []byte("&#x9;3 &lt; 5&#xD;&#xA;3 &gt; 2&#x0;"),
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer

		_, err := htmlescape.WriteString(&buffer, test.String)

		if nil != err {
			t.Errorf("For test #%d, did not expect to get an error but actually got one." , testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("STRING: %q", test.String)
			continue
		}

		{
			expected := test.Expected
			actual := buffer.Bytes()

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual resulting value of the append is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("STRING: %q", test.String)
				continue
			}
		}
	}
}
