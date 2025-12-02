package main

import "testing"

func TestProcessDocumentPart2(t *testing.T) {
	got := ProcessDocumentPart2("../inputs/test.txt")
	want := 6

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
