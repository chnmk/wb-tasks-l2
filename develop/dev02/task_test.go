package main

import (
	"testing"
)

func TestUnpackEmptyString(t *testing.T) {
	_, err := Unpack("")

	if err == nil {
		t.Error("expected empty output")
	}
}

func TestUnpackInvalidString(t *testing.T) {
	_, err := Unpack("45")

	if err == nil {
		t.Error("expected empty output")
	}
}

func TestUnpackNoNumbers(t *testing.T) {
	r, err := Unpack("abcd")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if r != "abcd" {
		t.Error("expected same string")
	}
}

func TestUnpackValidString(t *testing.T) {
	r, err := Unpack("a4bc2d5e")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if r != "aaaabccddddde" {
		t.Error("wrong output")
	}
}

func TestUnpackEscapeChars(t *testing.T) {
	r, err := Unpack("qwe\\4\\5")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if r != "qwe45" {
		t.Errorf("wrong output: expected qwe45, got %s", r)
	}
}

func TestUnpackEscapeCharsTwoDigits(t *testing.T) {
	r, err := Unpack("qwe\\45")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if r != "qwe44444" {
		t.Errorf("wrong output: expected qwe44444, got %s", r)
	}
}

func TestUnpackEscapeCharsBackslashes(t *testing.T) {
	r, err := Unpack("qwe\\\\5")
	if err != nil {
		t.Fatalf(err.Error())
	}

	if r != "qwe\\\\\\\\\\" {
		t.Errorf("wrong output: expected qwe\\\\\\\\\\, got %s", r)
	}
}
