package main

import "testing"

func TestHumanAge(t *testing.T) {
	humanAge := humanAge(3)
	if humanAge != 21 {
		t.Fatal("No match, got", humanAge)
	}
}
