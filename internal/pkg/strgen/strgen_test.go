package strgen

import (
	"goapp/pkg/util"
	"testing"
)

func TestRandomHex(t *testing.T) {
	hex, err := util.RandomHex(5)

	if err != nil {
		t.Errorf("Error while generating hex string: %v", err)
	}

	//We need 10 hexadecimal characters to represent 5 bytes
	if len(hex) != 10 {
		t.Errorf("Invalid hexadecimal string length, (%d)", len(hex))
	}

	for _, ch := range hex {
		if (ch < '0' || ch > '9') && (ch < 'a' || ch > 'f') {
			t.Errorf("Invalid hexadecimal character, %v", ch)
		}
	}
}

func BenchmarkRandomHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		util.RandomHex(5)
	}
}
