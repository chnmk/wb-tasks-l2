package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestDefaultSort(t *testing.T) {
	s1 := "ббб"
	s2 := "ааа"

	list := []string{s1, s2}

	list, err := sortList(list)
	if err != nil {
		t.Error(err)
	}

	if list[0] != "ааа" {
		t.Error("list is unsorted")
	}
}

func TestReverseSort(t *testing.T) {
	r = true
	t.Cleanup(func() { r = false })

	s1 := "ааа"
	s2 := "ббб"

	list := []string{s1, s2}

	list, err := sortList(list)
	if err != nil {
		t.Error(err)
	}

	if list[0] != "ббб" {
		t.Error("list is unsorted")
	}
}

func TestNoDublicateSort(t *testing.T) {
	u = true
	t.Cleanup(func() { u = false })

	s1 := "ббб"
	s2 := "ааа"
	s3 := "ааа"

	list := []string{s1, s2, s3}

	list, err := sortList(list)
	if err != nil {
		t.Error(err)
	}

	if list[0] != "ааа" {
		t.Error("list is unsorted")
	}

	if len(list) == 3 {
		t.Error("dublicate found")
	}
}

func TestReverseNoDublicateSort(t *testing.T) {
	u = true
	r = true
	t.Cleanup(func() {
		u = false
		r = false
	})

	s1 := "ааа"
	s2 := "ббб"
	s3 := "ббб"

	list := []string{s1, s2, s3}

	list, err := sortList(list)
	if err != nil {
		t.Error(err)
	}

	if list[0] != "ббб" {
		t.Error("list is unsorted")
	}

	if len(list) == 3 {
		t.Error("dublicate found")
	}
}

func TestNumSort(t *testing.T) {
	n = true
	t.Cleanup(func() { n = false })

	s1 := "10"
	s2 := "9"
	s3 := "90"
	s4 := "11"

	list := []string{s1, s2, s3, s4}

	list, err := sortList(list)
	if err != nil {
		t.Error(err)
	}

	if list[0] != "9" {
		t.Error("list is unsorted")
	}
}

func TestSuffixSort(t *testing.T) {
	h = true
	t.Cleanup(func() { h = false })

	s1 := "4"
	s2 := "4.2"
	s3 := "4.1"
	s4 := "4.1"
	s5 := "3.8"
	s6 := "3.7"
	s7 := "2"
	s8 := "5"
	s9 := "1"

	list := []string{s1, s2, s3, s4, s5, s6, s7, s8, s9}

	list, err := sortList(list)
	if err != nil {
		t.Error(err)
	}

	if list[0] != "1" || list[1] != "2" || list[6] != "4.1" || list[7] != "4.2" {
		t.Error("list is unsorted")
	}
}

func TestMonthSort(t *testing.T) {
	M = true
	t.Cleanup(func() { M = false })

	s1 := "Февраль"
	s2 := "Январь"

	list := []string{s1, s2}

	list, err := sortList(list)
	if err != nil {
		t.Error(err)
	}

	if list[0] != "Январь" {
		t.Error("list is unsorted")
	}
}

func TestMonthReverseSort(t *testing.T) {
	M = true
	r = true
	t.Cleanup(func() {
		M = false
		r = false
	})

	s1 := "Январь"
	s2 := "Февраль"

	list := []string{s1, s2}

	list, err := sortList(list)
	if err != nil {
		t.Error(err)
	}

	if list[0] != "Февраль" {
		t.Error("list is unsorted")
	}
}

func TestMonthColumnSort(t *testing.T) {
	M = true
	k = "Февраль"
	t.Cleanup(func() {
		M = false
		k = ""
	})

	s1 := "тест Февраль"
	s2 := "тест Январь"

	list := []string{s1, s2}

	list, err := sortList(list)
	if err != nil {
		t.Error(err)
	}

	if list[0] != "тест Январь" {
		t.Error("list is unsorted")
	}
}

func TestCheckSort(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	M = true
	c = true

	t.Cleanup(func() {
		M = false
		c = false
		log.SetOutput(os.Stderr)
	})

	s1 := "Январь"
	s2 := "Февраль"

	list := []string{s1, s2}

	_, err := sortList(list)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(buf.String(), "данные уже отсортированы") {
		t.Error("invalid output")
	}
}
