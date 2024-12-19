package main

import (
	"testing"
)

func TestUnpackEmptyString(t *testing.T) {
	r := Unpack("")

	if r != "" {
		t.Error("expected empty output")
	}
}

func TestUnpackInvalidString(t *testing.T) {
	r := Unpack("45")

	if r != "" {
		t.Error("expected empty output")
	}
}

func TestUnpackNoNumbers(t *testing.T) {
	r := Unpack("abcd")

	if r != "abcd" {
		t.Error("expected same string")
	}
}

func TestUnpackValidString(t *testing.T) {
	r := Unpack("a4bc2d5e")

	if r != "aaaabccddddde" {
		t.Error("wrong output")
	}
}

func TestUnpackEscapeChars(t *testing.T) {
	r := Unpack("qwe\\4\\5")

	if r != "qwe45" {
		t.Errorf("wrong output: expected qwe45, got %s", r)
	}
}

func TestUnpackEscapeCharsTwoDigits(t *testing.T) {
	r := Unpack("qwe\\45")

	if r != "qwe44444" {
		t.Errorf("wrong output: expected qwe44444, got %s", r)
	}
}

func TestUnpackEscapeCharsBackslashes(t *testing.T) {
	r := Unpack("qwe\\\\5")

	if r != "qwe\\\\\\\\\\" {
		t.Errorf("wrong output: expected qwe\\\\\\\\\\, got %s", r)
	}
}
