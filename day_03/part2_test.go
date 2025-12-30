package day_03

import "testing"
import "github.com/stretchr/testify/assert"

func TestPart2(t *testing.T) {
	res, _ := Part2("test-input.txt")
	assert.Equal(t, 3121910778619, res)
}
