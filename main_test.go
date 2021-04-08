package main

import (
	"bytes"
	"testing"
)

func TestShouldReturnEmptyStringIfMultipleWords(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("multiple words"))
	got := ReadString(&stdin)
	want := ""

	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}

}

func TestShouldReturnStringIfValid(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("validtext"))

	got := ReadString(&stdin)
	want := "validtext"

	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}
