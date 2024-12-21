package main

import (
	"testing"
)

func TestSearchDefault(t *testing.T) {
	s := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	m := *Search(&s)

	words1, ok1 := m["пятак"]
	words2, ok2 := m["листок"]

	if !ok1 || !ok2 {
		t.Fatal("key not found")
	}

	if words1 == nil || words2 == nil {
		t.Fatal("slice not found")
	}

	if words1[0] != "пятак" || words1[1] != "пятка" || words1[2] != "тяпка" {
		t.Fatal("wrong result")
	}

	if words2[0] != "листок" || words2[1] != "слиток" || words2[2] != "столик" {
		t.Fatal("wrong result")
	}
}

func TestSearchDublicates(t *testing.T) {
	s := []string{"аб", "аб", "ба"}

	m := *Search(&s)

	words, ok := m["аб"]

	if !ok || words == nil {
		t.Fatal("key not found")
	}

	if len(words) > 2 {
		t.Fatal("unexpected dublicate")
	}

}

func TestSearchSort(t *testing.T) {
	s := []string{"ба", "аб"}

	m := *Search(&s)

	words, ok := m["ба"]

	if !ok || words == nil {
		t.Fatal("key not found")
	}

	if words[0] != "аб" {
		t.Fatal("wrong order")
	}
}
