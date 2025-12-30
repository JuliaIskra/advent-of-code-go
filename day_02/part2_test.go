package day_02

import "testing"
import "github.com/stretchr/testify/assert"

func TestPart2(t *testing.T) {
	res, _ := Part2("test-input.txt")
	assert.Equal(t, 4174379265, res)
}
