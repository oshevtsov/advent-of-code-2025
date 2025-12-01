package main

import "testing"

func TestProcessDocumentPart1(t *testing.T) {
	got := ProcessDocumentPart1("../inputs/test.txt")
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
