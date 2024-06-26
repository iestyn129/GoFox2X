package gofox2x

import (
	"testing"
)

func TestSFSType_String(t *testing.T) {
	var val1 = makeSFSType("ok", true).String()
	var val2 = "(bool) ok: true"
	if val1 != val2 {
		t.Error("val1 != val2")
	}
}

func TestSFSTypeMaking(t *testing.T) {
	var val1 = makeSFSType("ok", true)
	var val2 = sfsType{SFSBOOL, "ok", true}
	if val1 != val2 {
		t.Error("val1 != val2")
	}
}
