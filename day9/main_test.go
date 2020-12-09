package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Version1(t *testing.T) {
	lines := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	res := getInvalidNumber(lines, 5)

	assert.Equal(t, 127, res)

	res2 := findWeakness(lines, 127)

	assert.Equal(t, 62, res2)

}
