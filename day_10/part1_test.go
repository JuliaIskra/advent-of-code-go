package day_10

import "testing"
import "github.com/stretchr/testify/assert"

func TestPart1(t *testing.T) {
	res, _ := Part1("test-input.txt")
	assert.Equal(t, 7, res)
}
