package main

import (
	"strings"
	"testing"
)

func TestExpandPath(t *testing.T) {
	const Home = "~"

	path := "~/Code/go"
	got := ExpandPath(path)

	if strings.Contains(got, Home) {
		t.Errorf(`ExpandPath() for %v is invalid`, path)
	}
}
