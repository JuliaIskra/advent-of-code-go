package day_01

import "testing"
import "github.com/stretchr/testify/assert"

func TestPart2(t *testing.T) {
	res, _ := Part2("test-input.txt")
	assert.Equal(t, 6, res)
}

// This is an example test that intentionally fails.
//func TestCountZeros(t *testing.T) {
//	p, zeros := CountZeros(0, 10)
//	// Intentionally wrong expectation to demonstrate a failing test
//	if p != 10 || zeros != 1 {
//		t.Fatalf("Invalid result (intentional fail): got p=%d zeros=%d, want p=10 zeros=1", p, zeros)
//	}
//}
