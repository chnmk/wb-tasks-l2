package main

import (
	"testing"
)

func TestFilter(t *testing.T) {
	pattern = "Втор"
	t.Cleanup(func() { pattern = "" })

	s := []string{"Первый", "Второй"}
	s = filterFile(s)

	if len(s) != 1 {
		t.Fatal("filter failed")
	}

	if s[0] != "Второй" {
		t.Error("wrong result")
	}
}

func TestFilterAfter(t *testing.T) {
	pattern = "Втор"
	A = 1
	t.Cleanup(func() {
		pattern = ""
		A = 0
	})

	s := []string{"Первый", "Второй", "Третий"}
	s = filterFile(s)

	if len(s) != 2 {
		t.Fatal("filter failed")
	}

	if s[0] != "Второй" || s[1] != "Третий" {
		t.Error("wrong result")
	}
}

func TestFilterBefore(t *testing.T) {
	pattern = "Трет"
	B = 1
	t.Cleanup(func() {
		pattern = ""
		B = 0
	})

	s := []string{"Первый", "Второй", "Третий"}
	s = filterFile(s)

	if len(s) != 2 {
		t.Fatal("filter failed")
	}

	if s[0] != "Второй" || s[1] != "Третий" {
		t.Error("wrong result")
	}
}

func TestFilterAround(t *testing.T) {
	pattern = "Трет"
	C = 1
	t.Cleanup(func() {
		pattern = ""
		C = 0
	})

	s := []string{"Первый", "Второй", "Третий", "Четвёртый"}
	s = filterFile(s)

	if len(s) != 3 {
		t.Fatal("filter failed")
	}

	if s[0] != "Второй" || s[1] != "Третий" || s[2] != "Четвёртый" {
		t.Log(s)
		t.Error("wrong result")
	}
}

func TestFilterCount(t *testing.T) {
	pattern = "Втор"
	c = 3
	t.Cleanup(func() {
		pattern = ""
		c = 0
	})

	s := []string{"Первый", "Второй", "Третий", "Второй2"}
	s = filterFile(s)

	if len(s) != 1 {
		t.Fatal("filter failed")
	}

	if s[0] != "Второй" {
		t.Error("wrong result")
	}
}

func TestFilterIgnoreCase(t *testing.T) {
	pattern = "втор"
	i = true
	t.Cleanup(func() {
		pattern = ""
		i = false
	})

	s := []string{"Первый", "Второй", "Третий"}
	s = filterFile(s)

	if len(s) != 1 {
		t.Fatal("filter failed")
	}

	if s[0] != "Второй" {
		t.Error("wrong result")
	}
}

func TestFilterInvert(t *testing.T) {
	pattern = "Втор"
	v = true
	t.Cleanup(func() {
		pattern = ""
		v = false
	})

	s := []string{"Первый", "Второй", "Третий"}
	s = filterFile(s)

	if len(s) != 2 {
		t.Fatal("filter failed")
	}

	if s[0] != "Первый" || s[1] != "Третий" {
		t.Error("wrong result")
	}
}

func TestFilterFixed(t *testing.T) {
	pattern = "Втор"
	F = true
	t.Cleanup(func() {
		pattern = ""
		F = false
	})

	s := []string{"Первый", "Втор", "Второй"}
	s = filterFile(s)

	if len(s) != 1 {
		t.Fatal("filter failed")
	}

	if s[0] != "Втор" {
		t.Error("wrong result")
	}
}

func TestFilterLineNum(t *testing.T) {
	pattern = "Втор"
	n = true
	t.Cleanup(func() {
		pattern = ""
		n = false
	})

	s := []string{"Первый", "Второй"}
	s = filterFile(s)

	if len(s) != 1 {
		t.Fatal("filter failed")
	}

	if s[0] != "2 Второй" {
		t.Log(s)
		t.Error("wrong result")
	}
}
