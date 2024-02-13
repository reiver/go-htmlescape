package htmlescape

import (
	"testing"

	"bytes"
)

func TestSafeRune(t *testing.T) {

	tests := []struct{
		Rune rune
		ExpectedEscaped []byte
		ExpectedWasEscaped bool
	}{
		{
			Rune:                   '\x00', // NUL
			ExpectedEscaped: []byte("&#x0;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x01', // SOH
			ExpectedEscaped: []byte("&#x1;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x02', // STX
			ExpectedEscaped: []byte("&#x2;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x03', // ETX
			ExpectedEscaped: []byte("&#x3;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x04', // EOT
			ExpectedEscaped: []byte("&#x4;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x05', // ENQ
			ExpectedEscaped: []byte("&#x5;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x06', // ACK
			ExpectedEscaped: []byte("&#x6;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x07', // BEL
			ExpectedEscaped: []byte("&#x7;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x08', // BS
			ExpectedEscaped: []byte("&#x8;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x09', // HT
			ExpectedEscaped: []byte("&#x9;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x0A', // LF
			ExpectedEscaped: []byte("&#xA;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x0B', // VT
			ExpectedEscaped: []byte("&#xB;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x0C', // FF
			ExpectedEscaped: []byte("&#xC;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x0D', // CR
			ExpectedEscaped: []byte("&#xD;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x0E', // SO
			ExpectedEscaped: []byte("&#xE;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                   '\x0F', // SI
			ExpectedEscaped: []byte("&#xF;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x10', // DLE
			ExpectedEscaped: []byte("&#x10;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x11', // DC1
			ExpectedEscaped: []byte("&#x11;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x12', // DC2
			ExpectedEscaped: []byte("&#x12;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x13', // DC3
			ExpectedEscaped: []byte("&#x13;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x14', // DC4
			ExpectedEscaped: []byte("&#x14;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x15', // NAK
			ExpectedEscaped: []byte("&#x15;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x16', // SYN
			ExpectedEscaped: []byte("&#x16;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x17', // ETB
			ExpectedEscaped: []byte("&#x17;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x18', // CAN
			ExpectedEscaped: []byte("&#x18;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x19', // EM
			ExpectedEscaped: []byte("&#x19;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x1A', // SUB
			ExpectedEscaped: []byte("&#x1A;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x1B', // ESC
			ExpectedEscaped: []byte("&#x1B;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x1C', // FS
			ExpectedEscaped: []byte("&#x1C;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x1D', // GS
			ExpectedEscaped: []byte("&#x1D;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x1E', // RS
			ExpectedEscaped: []byte("&#x1E;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune:                    '\x1F', // US
			ExpectedEscaped: []byte("&#x1F;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: ' ', // 0x20
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '!', // 0x21
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '"', // 0x22 == 34
			ExpectedEscaped: []byte("&#34;"), // "&quot;"
			ExpectedWasEscaped: true,
		},
		{
			Rune: '#', // 0x23
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '$', // 0x24
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '%', // 0x25
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '&', // 0x26
			ExpectedEscaped: []byte("&amp;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\'', // 0x27
			ExpectedEscaped: []byte("&#39;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '(', // 0x28
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: ')', // 0x29
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '*', // 0x2A
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '+', // 0x2B
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: ',', // 0x2C
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '-', // 0x2D
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '.', // 0x2E
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '/', // 0x2F
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '0', // 0x30
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '1', // 0x31
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '2', // 0x32
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '3', // 0x33
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '4', // 0x34
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '5', // 0x35
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '6', // 0x36
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '7', // 0x37
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '8', // 0x38
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '9', // 0x39
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: ':', // 0x3A
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: ';', // 0x3B
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '<', // 0x3C
			ExpectedEscaped: []byte("&lt;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '=', // 0x3D
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '>', // 0x3E
			ExpectedEscaped: []byte("&gt;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '?', // 0x3F
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '@', // 0x40
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'A', // 0x41
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'B', // 0x42
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'C', // 0x43
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'D', // 0x44
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'E', // 0x45
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'F', // 0x46
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'G', // 0x47
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'H', // 0x48
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'I', // 0x49
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'J', // 0x4A
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'K', // 0x4B
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'L', // 0x4C
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'M', // 0x4D
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'N', // 0x4E
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'O', // 0x4F
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'P', // 0x5
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'Q', // 0x51
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'R', // 0x52
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'S', // 0x53
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'T', // 0x54
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'U', // 0x55
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'V', // 0x56
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'W', // 0x57
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'X', // 0x58
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'Y', // 0x59
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'Z', // 0x5A
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '[', // 0x5B
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\\', // 0x5C
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: ']', // 0x5D
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '^', // 0x5E
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '_', // 0x5F
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '`', // 0x60
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'a', // 0x61
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'b', // 0x62
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'c', // 0x63
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'd', // 0x64
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'e', // 0x65
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'f', // 0x66
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'g', // 0x67
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'h', // 0x68
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'i', // 0x69
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'j', // 0x6A
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'k', // 0x6B
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'l', // 0x6C
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'm', // 0x6D
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'n', // 0x6E
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'o', // 0x6F
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'p', // 0x70
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'q', // 0x71
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'r', // 0x72
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 's', // 0x73
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 't', // 0x74
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'u', // 0x75
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'v', // 0x76
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'w', // 0x77
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'x', // 0x78
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'y', // 0x79
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: 'z', // 0x7A
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '{', // 0x7B
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '|', // 0x7C
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '}', // 0x7D
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '~', // 0x7E
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\x7F', // DEL
			ExpectedEscaped: []byte("&#x7F;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x80', // PAD
			ExpectedEscaped: []byte("&#x80;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x81', // HOP
			ExpectedEscaped: []byte("&#x81;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x82', // BPH
			ExpectedEscaped: []byte("&#x82;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x83', // NBH
			ExpectedEscaped: []byte("&#x83;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x84', // IND
			ExpectedEscaped: []byte("&#x84;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x85', // NEL
			ExpectedEscaped: []byte("&#x85;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x86', // SSA
			ExpectedEscaped: []byte("&#x86;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x87', // ESA
			ExpectedEscaped: []byte("&#x87;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x88', // HTS
			ExpectedEscaped: []byte("&#x88;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x89', // HTJ
			ExpectedEscaped: []byte("&#x89;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x8A', // VTS
			ExpectedEscaped: []byte("&#x8A;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x8B', // PLD
			ExpectedEscaped: []byte("&#x8B;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x8C', // PLU
			ExpectedEscaped: []byte("&#x8C;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x8D', // RI
			ExpectedEscaped: []byte("&#x8D;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x8E', // SS2
			ExpectedEscaped: []byte("&#x8E;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x8F', // SS3
			ExpectedEscaped: []byte("&#x8F;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x90', // DCS
			ExpectedEscaped: []byte("&#x90;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x91', // PU1
			ExpectedEscaped: []byte("&#x91;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x92', // PU2
			ExpectedEscaped: []byte("&#x92;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x93', // STS
			ExpectedEscaped: []byte("&#x93;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x94', // CCH
			ExpectedEscaped: []byte("&#x94;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x95', // MW
			ExpectedEscaped: []byte("&#x95;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x96', // SPA
			ExpectedEscaped: []byte("&#x96;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x97', // EPA
			ExpectedEscaped: []byte("&#x97;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x98', // SOS
			ExpectedEscaped: []byte("&#x98;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x99', // SGC (SGCI)
			ExpectedEscaped: []byte("&#x99;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x9A', // SCI
			ExpectedEscaped: []byte("&#x9A;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x9B', // CSI
			ExpectedEscaped: []byte("&#x9B;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x9C', // ST
			ExpectedEscaped: []byte("&#x9C;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x9D', // OSC
			ExpectedEscaped: []byte("&#x9D;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x9E', // PM
			ExpectedEscaped: []byte("&#x9E;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\x9F', // APC
			ExpectedEscaped: []byte("&#x9F;"),
			ExpectedWasEscaped: true,
		},
		{
			Rune: '\u00A0',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00A1',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00A2',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00A3',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00A4',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00A5',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00A6',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00A7',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00A8',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00A9',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00AA',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00AB',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00AC',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00AD',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00AE',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00AF',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00B0',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00B1',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00B2',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00B3',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00B4',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00B5',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00B6',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00B7',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00B8',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00B9',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00BA',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00BB',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00BC',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00BD',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00BE',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00BF',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
		{
			Rune: '\u00C0',
			ExpectedEscaped: nil,
			ExpectedWasEscaped: false,
		},
	}

	for testNumber, test := range tests {

		actualEscaped, actuallyWasEscaped := safeRune(test.Rune)

		{
			expected := test.ExpectedWasEscaped
			actual   := actuallyWasEscaped

			if expected != actual {
				t.Errorf("For test #%d, the actual value for was-escaped is not what was expected." , testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("RUNE: %q (%U)", test.Rune, test.Rune)
				continue
			}
		}

		{
			expected := test.ExpectedEscaped
			actual   := actualEscaped

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual escaped-rune is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("RUNE: %q (%U)", test.Rune, test.Rune)
				continue
			}
		}
	}
}
