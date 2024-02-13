package htmlescape_test

import (
	"testing"

	"bytes"

	"sourcecode.social/reiver/go-htmlescape"
)

func TestAppendRune(t *testing.T) {

	tests := []struct{
		Bytes []byte
		Rune rune
		Expected []byte
	}{
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x00,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x0;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x01,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x1;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x02,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x2;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x03,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x3;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x04,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x4;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x05,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x5;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x06,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x6;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x07,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x7;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x08,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x8;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x09,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x9;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x0A,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#xA;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x0B,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#xB;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x0C,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#xC;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x0D,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#xD;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x0E,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#xE;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x0F,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#xF;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x10,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x10;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x11,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x11;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x12,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x12;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x13,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x13;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x14,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x14;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x15,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x15;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x16,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x16;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x17,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x17;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x18,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x18;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x19,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x19;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x1A,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x1A;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x1B,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x1B;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x1C,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x1C;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x1D,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x1D;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x1E,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x1E;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x1F,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x1F;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x20,
			Expected: []byte("Hello world! 🙂 Take a look at this:  "),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x21,
			Expected: []byte("Hello world! 🙂 Take a look at this: !"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x22,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#34;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x23,
			Expected: []byte("Hello world! 🙂 Take a look at this: #"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x24,
			Expected: []byte("Hello world! 🙂 Take a look at this: $"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x25,
			Expected: []byte("Hello world! 🙂 Take a look at this: %"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x26,
			Expected: []byte("Hello world! 🙂 Take a look at this: &amp;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x27,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#39;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x28,
			Expected: []byte("Hello world! 🙂 Take a look at this: ("),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x29,
			Expected: []byte("Hello world! 🙂 Take a look at this: )"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x2A,
			Expected: []byte("Hello world! 🙂 Take a look at this: *"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x2B,
			Expected: []byte("Hello world! 🙂 Take a look at this: +"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x2C,
			Expected: []byte("Hello world! 🙂 Take a look at this: ,"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x2D,
			Expected: []byte("Hello world! 🙂 Take a look at this: -"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x2E,
			Expected: []byte("Hello world! 🙂 Take a look at this: ."),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x2F,
			Expected: []byte("Hello world! 🙂 Take a look at this: /"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x30,
			Expected: []byte("Hello world! 🙂 Take a look at this: 0"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x31,
			Expected: []byte("Hello world! 🙂 Take a look at this: 1"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x32,
			Expected: []byte("Hello world! 🙂 Take a look at this: 2"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x33,
			Expected: []byte("Hello world! 🙂 Take a look at this: 3"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x34,
			Expected: []byte("Hello world! 🙂 Take a look at this: 4"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x35,
			Expected: []byte("Hello world! 🙂 Take a look at this: 5"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x36,
			Expected: []byte("Hello world! 🙂 Take a look at this: 6"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x37,
			Expected: []byte("Hello world! 🙂 Take a look at this: 7"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x38,
			Expected: []byte("Hello world! 🙂 Take a look at this: 8"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x39,
			Expected: []byte("Hello world! 🙂 Take a look at this: 9"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x3A,
			Expected: []byte("Hello world! 🙂 Take a look at this: :"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x3B,
			Expected: []byte("Hello world! 🙂 Take a look at this: ;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x3C,
			Expected: []byte("Hello world! 🙂 Take a look at this: &lt;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x3D,
			Expected: []byte("Hello world! 🙂 Take a look at this: ="),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x3E,
			Expected: []byte("Hello world! 🙂 Take a look at this: &gt;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x3F,
			Expected: []byte("Hello world! 🙂 Take a look at this: ?"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x40,
			Expected: []byte("Hello world! 🙂 Take a look at this: @"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x41,
			Expected: []byte("Hello world! 🙂 Take a look at this: A"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x42,
			Expected: []byte("Hello world! 🙂 Take a look at this: B"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x43,
			Expected: []byte("Hello world! 🙂 Take a look at this: C"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x44,
			Expected: []byte("Hello world! 🙂 Take a look at this: D"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x45,
			Expected: []byte("Hello world! 🙂 Take a look at this: E"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x46,
			Expected: []byte("Hello world! 🙂 Take a look at this: F"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x47,
			Expected: []byte("Hello world! 🙂 Take a look at this: G"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x48,
			Expected: []byte("Hello world! 🙂 Take a look at this: H"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x49,
			Expected: []byte("Hello world! 🙂 Take a look at this: I"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x4A,
			Expected: []byte("Hello world! 🙂 Take a look at this: J"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x4B,
			Expected: []byte("Hello world! 🙂 Take a look at this: K"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x4C,
			Expected: []byte("Hello world! 🙂 Take a look at this: L"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x4D,
			Expected: []byte("Hello world! 🙂 Take a look at this: M"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x4E,
			Expected: []byte("Hello world! 🙂 Take a look at this: N"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x4F,
			Expected: []byte("Hello world! 🙂 Take a look at this: O"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x50,
			Expected: []byte("Hello world! 🙂 Take a look at this: P"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x51,
			Expected: []byte("Hello world! 🙂 Take a look at this: Q"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x52,
			Expected: []byte("Hello world! 🙂 Take a look at this: R"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x53,
			Expected: []byte("Hello world! 🙂 Take a look at this: S"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x54,
			Expected: []byte("Hello world! 🙂 Take a look at this: T"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x55,
			Expected: []byte("Hello world! 🙂 Take a look at this: U"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x56,
			Expected: []byte("Hello world! 🙂 Take a look at this: V"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x57,
			Expected: []byte("Hello world! 🙂 Take a look at this: W"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x58,
			Expected: []byte("Hello world! 🙂 Take a look at this: X"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x59,
			Expected: []byte("Hello world! 🙂 Take a look at this: Y"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x5A,
			Expected: []byte("Hello world! 🙂 Take a look at this: Z"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x5B,
			Expected: []byte("Hello world! 🙂 Take a look at this: ["),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x5C,
			Expected: []byte("Hello world! 🙂 Take a look at this: \\"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x5D,
			Expected: []byte("Hello world! 🙂 Take a look at this: ]"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x5E,
			Expected: []byte("Hello world! 🙂 Take a look at this: ^"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x5F,
			Expected: []byte("Hello world! 🙂 Take a look at this: _"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x60,
			Expected: []byte("Hello world! 🙂 Take a look at this: `"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x61,
			Expected: []byte("Hello world! 🙂 Take a look at this: a"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x62,
			Expected: []byte("Hello world! 🙂 Take a look at this: b"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x63,
			Expected: []byte("Hello world! 🙂 Take a look at this: c"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x64,
			Expected: []byte("Hello world! 🙂 Take a look at this: d"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x65,
			Expected: []byte("Hello world! 🙂 Take a look at this: e"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x66,
			Expected: []byte("Hello world! 🙂 Take a look at this: f"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x67,
			Expected: []byte("Hello world! 🙂 Take a look at this: g"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x68,
			Expected: []byte("Hello world! 🙂 Take a look at this: h"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x69,
			Expected: []byte("Hello world! 🙂 Take a look at this: i"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x6A,
			Expected: []byte("Hello world! 🙂 Take a look at this: j"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x6B,
			Expected: []byte("Hello world! 🙂 Take a look at this: k"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x6C,
			Expected: []byte("Hello world! 🙂 Take a look at this: l"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x6D,
			Expected: []byte("Hello world! 🙂 Take a look at this: m"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x6E,
			Expected: []byte("Hello world! 🙂 Take a look at this: n"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x6F,
			Expected: []byte("Hello world! 🙂 Take a look at this: o"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x70,
			Expected: []byte("Hello world! 🙂 Take a look at this: p"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x71,
			Expected: []byte("Hello world! 🙂 Take a look at this: q"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x72,
			Expected: []byte("Hello world! 🙂 Take a look at this: r"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x73,
			Expected: []byte("Hello world! 🙂 Take a look at this: s"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x74,
			Expected: []byte("Hello world! 🙂 Take a look at this: t"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x75,
			Expected: []byte("Hello world! 🙂 Take a look at this: u"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x76,
			Expected: []byte("Hello world! 🙂 Take a look at this: v"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x77,
			Expected: []byte("Hello world! 🙂 Take a look at this: w"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x78,
			Expected: []byte("Hello world! 🙂 Take a look at this: x"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x79,
			Expected: []byte("Hello world! 🙂 Take a look at this: y"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x7A,
			Expected: []byte("Hello world! 🙂 Take a look at this: z"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x7B,
			Expected: []byte("Hello world! 🙂 Take a look at this: {"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x7C,
			Expected: []byte("Hello world! 🙂 Take a look at this: |"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x7D,
			Expected: []byte("Hello world! 🙂 Take a look at this: }"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x7E,
			Expected: []byte("Hello world! 🙂 Take a look at this: ~"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x7F,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x7F;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x80,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x80;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x81,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x81;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x82,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x82;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x83,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x83;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x84,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x84;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x85,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x85;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x86,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x86;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x87,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x87;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x88,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x88;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x89,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x89;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x8A,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x8A;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x8B,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x8B;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x8C,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x8C;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x8D,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x8D;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x8E,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x8E;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x8F,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x8F;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x90,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x90;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x91,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x91;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x92,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x92;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x93,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x93;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x94,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x94;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x95,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x95;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x96,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x96;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x97,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x97;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x98,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x98;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x99,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x99;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x9A,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x9A;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x9B,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x9B;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x9C,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x9C;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x9D,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x9D;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x9E,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x9E;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0x9F,
			Expected: []byte("Hello world! 🙂 Take a look at this: &#x9F;"),
		},
		{
			Bytes:    []byte("Hello world! 🙂 Take a look at this: "),
			Rune: 0xA0,
			Expected: []byte("Hello world! 🙂 Take a look at this: \u00A0"),
		},
	}

	for testNumber, test := range tests {

		var p []byte = append([]byte(nil), test.Bytes...)
		actual := htmlescape.AppendRune(p, test.Rune)

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual resulting value of the append is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("RUNE: %q (%U)", test.Rune, test.Rune)
			t.Logf("BYTES: %q", test.Bytes)
			continue
		}
	}
}
