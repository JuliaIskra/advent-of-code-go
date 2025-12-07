//go:build ignore
// +build ignore

package day_01

import "testing"

func TestCountZeros(t *testing.T) {
	var p, zeros int
	p, zeros = CountZeros(0, 10)
	if p != 10 || zeros != 0 {
		t.Fatalf("Invalid result")
	}
}
