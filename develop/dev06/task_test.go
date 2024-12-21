package main

import "testing"

func TestCutDefault(t *testing.T) {
	f_split = []string{"Id"}
	t.Cleanup(func() {
		f_split = []string{}
	})

	s := []string{"Id\tName", "1\tПетя"}
	s = cut(s)

	if len(s) != 2 {
		t.Fatal("cut failed")
	}

	if s[0] != "Id" || s[1] != "1" {
		t.Error("wrong result")
	}
}

func TestCutSpaceDelim(t *testing.T) {
	f_split = []string{"Id"}
	d = " "
	t.Cleanup(func() {
		f_split = []string{}
		d = ""
	})

	s := []string{"Id Name", "1 Петя"}
	s = cut(s)

	if len(s) != 2 {
		t.Fatal("cut failed")
	}

	if s[0] != "Id" || s[1] != "1" {
		t.Error("wrong result")
	}
}

func TestCutNoEmpty(t *testing.T) {
	f_split = []string{"Id"}
	s = true
	d = " "
	t.Cleanup(func() {
		f_split = []string{}
		s = false
		d = ""
	})

	s := []string{"Id Name", "1 Петя", "2Вася", "Тест "}
	s = cut(s)

	if len(s) != 3 {
		t.Fatal("cut failed")
	}

	if s[2] != "Тест" {
		t.Error("wrong result")
	}
}
