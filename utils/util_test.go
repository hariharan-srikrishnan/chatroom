package utils

import (
	"testing"
)

func TestStrToBytes(t *testing.T) {
	input := "OK"
	target := []byte("OK")
	converted := StrToBytes(&input)
	if len(target) != len(converted) {
		t.Fail()
	}
	for i := 0; i < len(target); i++ {
		if target[i] != converted[i] {
			t.Fail()
		}
	}
}

func TestBytesToStr(t *testing.T) {
	input := []byte("OK")
	target := "OK"
	converted := BytesToStr(input)
	if len(target) != len(converted) {
		t.Fail()
	}
	for i := 0; i < len(target); i++ {
		if target[i] != converted[i] {
			t.Fail()
		}
	}
}