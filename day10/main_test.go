package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Version1(t *testing.T) {
	lines := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}

	cs := getChargers(lines)

	assert.Equal(t, 7*5, cs.getDiffNumber())

	combs := getCombinitionsPossible(lines)

	testCombos := []string{
		"1456710111215161922",
		"14567101215161922",
		"145710111215161922",
		"1457101215161922",
		"146710111215161922",
		"1467101215161922",
		"14710111215161922",
		"147101215161922",
	}

	for _, tc := range testCombos {
		found := false
		for c := range combs {
			if tc == c {
				found = true
			}
		}
		assert.Equal(t, true, found, tc)
	}

	assert.Equal(t, 8, len(combs))

}

func TestDay2Version1_2(t *testing.T) {
	lines := []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}

	cs := getChargers(lines)

	assert.Equal(t, 220, cs.getDiffNumber())

	combs := getCombinitionsPossible(lines)

	assert.Equal(t, 19208, len(combs))

}
