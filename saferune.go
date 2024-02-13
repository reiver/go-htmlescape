package htmlescape

var escaped00 []byte = []byte("&#x0;")
var escaped01 []byte = []byte("&#x1;")
var escaped02 []byte = []byte("&#x2;")
var escaped03 []byte = []byte("&#x3;")
var escaped04 []byte = []byte("&#x4;")
var escaped05 []byte = []byte("&#x5;")
var escaped06 []byte = []byte("&#x6;")
var escaped07 []byte = []byte("&#x7;")
var escaped08 []byte = []byte("&#x8;")
var escaped09 []byte = []byte("&#x9;")
var escaped0A []byte = []byte("&#xA;")
var escaped0B []byte = []byte("&#xB;")
var escaped0C []byte = []byte("&#xC;")
var escaped0D []byte = []byte("&#xD;")
var escaped0E []byte = []byte("&#xE;")
var escaped0F []byte = []byte("&#xF;")

var escaped10 []byte = []byte("&#x10;")
var escaped11 []byte = []byte("&#x11;")
var escaped12 []byte = []byte("&#x12;")
var escaped13 []byte = []byte("&#x13;")
var escaped14 []byte = []byte("&#x14;")
var escaped15 []byte = []byte("&#x15;")
var escaped16 []byte = []byte("&#x16;")
var escaped17 []byte = []byte("&#x17;")
var escaped18 []byte = []byte("&#x18;")
var escaped19 []byte = []byte("&#x19;")
var escaped1A []byte = []byte("&#x1A;")
var escaped1B []byte = []byte("&#x1B;")
var escaped1C []byte = []byte("&#x1C;")
var escaped1D []byte = []byte("&#x1D;")
var escaped1E []byte = []byte("&#x1E;")
var escaped1F []byte = []byte("&#x1F;")

const minCtrlBlock1 = 0x7F

var escaped7F []byte = []byte("&#x7F;")
var escaped80 []byte = []byte("&#x80;")
var escaped81 []byte = []byte("&#x81;")
var escaped82 []byte = []byte("&#x82;")
var escaped83 []byte = []byte("&#x83;")
var escaped84 []byte = []byte("&#x84;")
var escaped85 []byte = []byte("&#x85;")
var escaped86 []byte = []byte("&#x86;")
var escaped87 []byte = []byte("&#x87;")
var escaped88 []byte = []byte("&#x88;")
var escaped89 []byte = []byte("&#x89;")
var escaped8A []byte = []byte("&#x8A;")
var escaped8B []byte = []byte("&#x8B;")
var escaped8C []byte = []byte("&#x8C;")
var escaped8D []byte = []byte("&#x8D;")
var escaped8E []byte = []byte("&#x8E;")
var escaped8F []byte = []byte("&#x8F;")
var escaped90 []byte = []byte("&#x90;")
var escaped91 []byte = []byte("&#x91;")
var escaped92 []byte = []byte("&#x92;")
var escaped93 []byte = []byte("&#x93;")
var escaped94 []byte = []byte("&#x94;")
var escaped95 []byte = []byte("&#x95;")
var escaped96 []byte = []byte("&#x96;")
var escaped97 []byte = []byte("&#x97;")
var escaped98 []byte = []byte("&#x98;")
var escaped99 []byte = []byte("&#x99;")
var escaped9A []byte = []byte("&#x9A;")
var escaped9B []byte = []byte("&#x9B;")
var escaped9C []byte = []byte("&#x9C;")
var escaped9D []byte = []byte("&#x9D;")
var escaped9E []byte = []byte("&#x9E;")
var escaped9F []byte = []byte("&#x9F;")

var ctrlBlock0 [][]byte = [][]byte{
	escaped00,
	escaped01,
	escaped02,
	escaped03,
	escaped04,
	escaped05,
	escaped06,
	escaped07,
	escaped08,
	escaped09,
	escaped0A,
	escaped0B,
	escaped0C,
	escaped0D,
	escaped0E,
	escaped0F,

	escaped10,
	escaped11,
	escaped12,
	escaped13,
	escaped14,
	escaped15,
	escaped16,
	escaped17,
	escaped18,
	escaped19,
	escaped1A,
	escaped1B,
	escaped1C,
	escaped1D,
	escaped1E,
	escaped1F,
}

var ctrlBlock1 [][]byte = [][]byte{
	escaped7F,

	escaped80,
	escaped81,
	escaped82,
	escaped83,
	escaped84,
	escaped85,
	escaped86,
	escaped87,
	escaped88,
	escaped89,
	escaped8A,
	escaped8B,
	escaped8C,
	escaped8D,
	escaped8E,
	escaped8F,

	escaped90,
	escaped91,
	escaped92,
	escaped93,
	escaped94,
	escaped95,
	escaped96,
	escaped97,
	escaped98,
	escaped99,
	escaped9A,
	escaped9B,
	escaped9C,
	escaped9D,
	escaped9E,
	escaped9F,
}

var escapedQuotationMark   []byte = []byte("&#34;")  //          "&#x22;" "&quot;"
var escapedApostrophe      []byte = []byte("&#39;")  //          "&#x27;" "&apos;"
var escapedAmpersand       []byte = []byte("&amp;")  // "&#38;"  "&#x26;"
var escapedLessThanSign    []byte = []byte("&lt;")   // "&#60;"  "&#x3C"
var escapedGreaterThanSign []byte = []byte("&gt;")   // "&#62;"  "&#x3E"

// safeRune will only escape the rune 'r' if it is necessary or wise to do so.
// The returned []byte is in UTF-8 — although all the values returned are in the sub-set of Unicode that is also ASCII.
// So it is both valid UTF-8 and ASCII.
func safeRune(r rune) (escapedRune []byte, wasEscaped bool) {
	switch r {
	case '"':
		return escapedQuotationMark, true
	case '\'':
		return escapedApostrophe, true
	case '&':
		return escapedAmpersand, true
	case '<':
		return escapedLessThanSign, true
	case '>':
		return escapedGreaterThanSign, true
	default:
	}

	if int64(0) <= int64(r) && int64(r) < int64(len(ctrlBlock0)) {
		return ctrlBlock0[int(r)], true
	}

	if minCtrlBlock1 <= int64(r) && int64(r) < (minCtrlBlock1 + int64(len(ctrlBlock1))) {
		return ctrlBlock1[int(r) - minCtrlBlock1], true
	}

	return nil, false
}
