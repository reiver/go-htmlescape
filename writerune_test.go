package htmlescape_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-htmlescape"
)

func TestWriteRune(t *testing.T) {

	tests := []struct{
		Rune rune
		Expected []byte
	}{
		{
			Rune: 0x00,
			Expected: []byte("&#x0;"),
		},
		{
			Rune: 0x01,
			Expected: []byte("&#x1;"),
		},
		{
			Rune: 0x02,
			Expected: []byte("&#x2;"),
		},
		{
			Rune: 0x03,
			Expected: []byte("&#x3;"),
		},
		{
			Rune: 0x04,
			Expected: []byte("&#x4;"),
		},
		{
			Rune: 0x05,
			Expected: []byte("&#x5;"),
		},
		{
			Rune: 0x06,
			Expected: []byte("&#x6;"),
		},
		{
			Rune: 0x07,
			Expected: []byte("&#x7;"),
		},
		{
			Rune: 0x08,
			Expected: []byte("&#x8;"),
		},
		{
			Rune: 0x09,
			Expected: []byte("&#x9;"),
		},
		{
			Rune: 0x0A,
			Expected: []byte("&#xA;"),
		},
		{
			Rune: 0x0B,
			Expected: []byte("&#xB;"),
		},
		{
			Rune: 0x0C,
			Expected: []byte("&#xC;"),
		},
		{
			Rune: 0x0D,
			Expected: []byte("&#xD;"),
		},
		{
			Rune: 0x0E,
			Expected: []byte("&#xE;"),
		},
		{
			Rune: 0x0F,
			Expected: []byte("&#xF;"),
		},
		{
			Rune: 0x10,
			Expected: []byte("&#x10;"),
		},
		{
			Rune: 0x11,
			Expected: []byte("&#x11;"),
		},
		{
			Rune: 0x12,
			Expected: []byte("&#x12;"),
		},
		{
			Rune: 0x13,
			Expected: []byte("&#x13;"),
		},
		{
			Rune: 0x14,
			Expected: []byte("&#x14;"),
		},
		{
			Rune: 0x15,
			Expected: []byte("&#x15;"),
		},
		{
			Rune: 0x16,
			Expected: []byte("&#x16;"),
		},
		{
			Rune: 0x17,
			Expected: []byte("&#x17;"),
		},
		{
			Rune: 0x18,
			Expected: []byte("&#x18;"),
		},
		{
			Rune: 0x19,
			Expected: []byte("&#x19;"),
		},
		{
			Rune: 0x1A,
			Expected: []byte("&#x1A;"),
		},
		{
			Rune: 0x1B,
			Expected: []byte("&#x1B;"),
		},
		{
			Rune: 0x1C,
			Expected: []byte("&#x1C;"),
		},
		{
			Rune: 0x1D,
			Expected: []byte("&#x1D;"),
		},
		{
			Rune: 0x1E,
			Expected: []byte("&#x1E;"),
		},
		{
			Rune: 0x1F,
			Expected: []byte("&#x1F;"),
		},
		{
			Rune: 0x20,
			Expected: []byte(" "),
		},
		{
			Rune: 0x21,
			Expected: []byte("!"),
		},
		{
			Rune: 0x22,
			Expected: []byte("&#34;"),
		},
		{
			Rune: 0x23,
			Expected: []byte("#"),
		},
		{
			Rune: 0x24,
			Expected: []byte("$"),
		},
		{
			Rune: 0x25,
			Expected: []byte("%"),
		},
		{
			Rune: 0x26,
			Expected: []byte("&amp;"),
		},
		{
			Rune: 0x27,
			Expected: []byte("&#39;"),
		},
		{
			Rune: 0x28,
			Expected: []byte("("),
		},
		{
			Rune: 0x29,
			Expected: []byte(")"),
		},
		{
			Rune: 0x2A,
			Expected: []byte("*"),
		},
		{
			Rune: 0x2B,
			Expected: []byte("+"),
		},
		{
			Rune: 0x2C,
			Expected: []byte(","),
		},
		{
			Rune: 0x2D,
			Expected: []byte("-"),
		},
		{
			Rune: 0x2E,
			Expected: []byte("."),
		},
		{
			Rune: 0x2F,
			Expected: []byte("/"),
		},
		{
			Rune: 0x30,
			Expected: []byte("0"),
		},
		{
			Rune: 0x31,
			Expected: []byte("1"),
		},
		{
			Rune: 0x32,
			Expected: []byte("2"),
		},
		{
			Rune: 0x33,
			Expected: []byte("3"),
		},
		{
			Rune: 0x34,
			Expected: []byte("4"),
		},
		{
			Rune: 0x35,
			Expected: []byte("5"),
		},
		{
			Rune: 0x36,
			Expected: []byte("6"),
		},
		{
			Rune: 0x37,
			Expected: []byte("7"),
		},
		{
			Rune: 0x38,
			Expected: []byte("8"),
		},
		{
			Rune: 0x39,
			Expected: []byte("9"),
		},
		{
			Rune: 0x3A,
			Expected: []byte(":"),
		},
		{
			Rune: 0x3B,
			Expected: []byte(";"),
		},
		{
			Rune: 0x3C,
			Expected: []byte("&lt;"),
		},
		{
			Rune: 0x3D,
			Expected: []byte("="),
		},
		{
			Rune: 0x3E,
			Expected: []byte("&gt;"),
		},
		{
			Rune: 0x3F,
			Expected: []byte("?"),
		},
		{
			Rune: 0x40,
			Expected: []byte("@"),
		},
		{
			Rune: 0x41,
			Expected: []byte("A"),
		},
		{
			Rune: 0x42,
			Expected: []byte("B"),
		},
		{
			Rune: 0x43,
			Expected: []byte("C"),
		},
		{
			Rune: 0x44,
			Expected: []byte("D"),
		},
		{
			Rune: 0x45,
			Expected: []byte("E"),
		},
		{
			Rune: 0x46,
			Expected: []byte("F"),
		},
		{
			Rune: 0x47,
			Expected: []byte("G"),
		},
		{
			Rune: 0x48,
			Expected: []byte("H"),
		},
		{
			Rune: 0x49,
			Expected: []byte("I"),
		},
		{
			Rune: 0x4A,
			Expected: []byte("J"),
		},
		{
			Rune: 0x4B,
			Expected: []byte("K"),
		},
		{
			Rune: 0x4C,
			Expected: []byte("L"),
		},
		{
			Rune: 0x4D,
			Expected: []byte("M"),
		},
		{
			Rune: 0x4E,
			Expected: []byte("N"),
		},
		{
			Rune: 0x4F,
			Expected: []byte("O"),
		},
		{
			Rune: 0x50,
			Expected: []byte("P"),
		},
		{
			Rune: 0x51,
			Expected: []byte("Q"),
		},
		{
			Rune: 0x52,
			Expected: []byte("R"),
		},
		{
			Rune: 0x53,
			Expected: []byte("S"),
		},
		{
			Rune: 0x54,
			Expected: []byte("T"),
		},
		{
			Rune: 0x55,
			Expected: []byte("U"),
		},
		{
			Rune: 0x56,
			Expected: []byte("V"),
		},
		{
			Rune: 0x57,
			Expected: []byte("W"),
		},
		{
			Rune: 0x58,
			Expected: []byte("X"),
		},
		{
			Rune: 0x59,
			Expected: []byte("Y"),
		},
		{
			Rune: 0x5A,
			Expected: []byte("Z"),
		},
		{
			Rune: 0x5B,
			Expected: []byte("["),
		},
		{
			Rune: 0x5C,
			Expected: []byte("\\"),
		},
		{
			Rune: 0x5D,
			Expected: []byte("]"),
		},
		{
			Rune: 0x5E,
			Expected: []byte("^"),
		},
		{
			Rune: 0x5F,
			Expected: []byte("_"),
		},
		{
			Rune: 0x60,
			Expected: []byte("`"),
		},
		{
			Rune: 0x61,
			Expected: []byte("a"),
		},
		{
			Rune: 0x62,
			Expected: []byte("b"),
		},
		{
			Rune: 0x63,
			Expected: []byte("c"),
		},
		{
			Rune: 0x64,
			Expected: []byte("d"),
		},
		{
			Rune: 0x65,
			Expected: []byte("e"),
		},
		{
			Rune: 0x66,
			Expected: []byte("f"),
		},
		{
			Rune: 0x67,
			Expected: []byte("g"),
		},
		{
			Rune: 0x68,
			Expected: []byte("h"),
		},
		{
			Rune: 0x69,
			Expected: []byte("i"),
		},
		{
			Rune: 0x6A,
			Expected: []byte("j"),
		},
		{
			Rune: 0x6B,
			Expected: []byte("k"),
		},
		{
			Rune: 0x6C,
			Expected: []byte("l"),
		},
		{
			Rune: 0x6D,
			Expected: []byte("m"),
		},
		{
			Rune: 0x6E,
			Expected: []byte("n"),
		},
		{
			Rune: 0x6F,
			Expected: []byte("o"),
		},
		{
			Rune: 0x70,
			Expected: []byte("p"),
		},
		{
			Rune: 0x71,
			Expected: []byte("q"),
		},
		{
			Rune: 0x72,
			Expected: []byte("r"),
		},
		{
			Rune: 0x73,
			Expected: []byte("s"),
		},
		{
			Rune: 0x74,
			Expected: []byte("t"),
		},
		{
			Rune: 0x75,
			Expected: []byte("u"),
		},
		{
			Rune: 0x76,
			Expected: []byte("v"),
		},
		{
			Rune: 0x77,
			Expected: []byte("w"),
		},
		{
			Rune: 0x78,
			Expected: []byte("x"),
		},
		{
			Rune: 0x79,
			Expected: []byte("y"),
		},
		{
			Rune: 0x7A,
			Expected: []byte("z"),
		},
		{
			Rune: 0x7B,
			Expected: []byte("{"),
		},
		{
			Rune: 0x7C,
			Expected: []byte("|"),
		},
		{
			Rune: 0x7D,
			Expected: []byte("}"),
		},
		{
			Rune: 0x7E,
			Expected: []byte("~"),
		},
		{
			Rune: 0x7F,
			Expected: []byte("&#x7F;"),
		},
		{
			Rune: 0x80,
			Expected: []byte("&#x80;"),
		},
		{
			Rune: 0x81,
			Expected: []byte("&#x81;"),
		},
		{
			Rune: 0x82,
			Expected: []byte("&#x82;"),
		},
		{
			Rune: 0x83,
			Expected: []byte("&#x83;"),
		},
		{
			Rune: 0x84,
			Expected: []byte("&#x84;"),
		},
		{
			Rune: 0x85,
			Expected: []byte("&#x85;"),
		},
		{
			Rune: 0x86,
			Expected: []byte("&#x86;"),
		},
		{
			Rune: 0x87,
			Expected: []byte("&#x87;"),
		},
		{
			Rune: 0x88,
			Expected: []byte("&#x88;"),
		},
		{
			Rune: 0x89,
			Expected: []byte("&#x89;"),
		},
		{
			Rune: 0x8A,
			Expected: []byte("&#x8A;"),
		},
		{
			Rune: 0x8B,
			Expected: []byte("&#x8B;"),
		},
		{
			Rune: 0x8C,
			Expected: []byte("&#x8C;"),
		},
		{
			Rune: 0x8D,
			Expected: []byte("&#x8D;"),
		},
		{
			Rune: 0x8E,
			Expected: []byte("&#x8E;"),
		},
		{
			Rune: 0x8F,
			Expected: []byte("&#x8F;"),
		},
		{
			Rune: 0x90,
			Expected: []byte("&#x90;"),
		},
		{
			Rune: 0x91,
			Expected: []byte("&#x91;"),
		},
		{
			Rune: 0x92,
			Expected: []byte("&#x92;"),
		},
		{
			Rune: 0x93,
			Expected: []byte("&#x93;"),
		},
		{
			Rune: 0x94,
			Expected: []byte("&#x94;"),
		},
		{
			Rune: 0x95,
			Expected: []byte("&#x95;"),
		},
		{
			Rune: 0x96,
			Expected: []byte("&#x96;"),
		},
		{
			Rune: 0x97,
			Expected: []byte("&#x97;"),
		},
		{
			Rune: 0x98,
			Expected: []byte("&#x98;"),
		},
		{
			Rune: 0x99,
			Expected: []byte("&#x99;"),
		},
		{
			Rune: 0x9A,
			Expected: []byte("&#x9A;"),
		},
		{
			Rune: 0x9B,
			Expected: []byte("&#x9B;"),
		},
		{
			Rune: 0x9C,
			Expected: []byte("&#x9C;"),
		},
		{
			Rune: 0x9D,
			Expected: []byte("&#x9D;"),
		},
		{
			Rune: 0x9E,
			Expected: []byte("&#x9E;"),
		},
		{
			Rune: 0x9F,
			Expected: []byte("&#x9F;"),
		},
		{
			Rune: 0xA0,
			Expected: []byte("\u00A0"),
		},
	}

	for testNumber, test := range tests {

		var buffer bytes.Buffer
		_, err := htmlescape.WriteRune(&buffer, test.Rune)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("RUNE: %q (%U)", test.Rune, test.Rune)
			continue
		}

		{
			actual := buffer.Bytes()
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual resulting value of the append is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("RUNE: %q (%U)", test.Rune, test.Rune)
				continue
			}
		}
	}
}
