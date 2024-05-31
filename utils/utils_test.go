package utils

import "testing"

func TestRandomCode(t *testing.T) {
	code := RandomCode()
	t.Log(code)
	if len(code) != 16 {
		t.Error("Unexpected return string length")
	}
}

func TestRandomCode_40characters(t *testing.T) {
	code := RandomCode(31)
	t.Log(code)
	if len(code) != 40 {
		t.Error("Unexpected return string length")
	}
}

func BenchmarkRandomCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomCode()
	}
}
