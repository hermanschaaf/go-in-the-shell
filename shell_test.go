package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestCommand(t *testing.T) {
	out, err := Command(`echo "hi"`).Output()
	expected := "hi"
	if err != nil {
		fmt.Println(err)
	}
	output := strings.TrimRight(string(out), " \t\n")
	if output != expected {
		t.Errorf("Error: %s != %s", output, expected)
	}
}
