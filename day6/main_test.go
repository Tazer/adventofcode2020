package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Version1(t *testing.T) {
	lines := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}

	summary := parseGroups(lines)

	a := summary.groups[0].persons[0].answers[0]

	assert.Equal(t, "a", a)

	g1 := summary.groups[0].getDistinctAnswers()

	assert.Equal(t, 3, g1)

	g2 := summary.groups[1].getDistinctAnswers()

	assert.Equal(t, 3, g2)

	g3 := summary.groups[2].getDistinctAnswers()

	assert.Equal(t, 3, g3)

	g4 := summary.groups[3].getDistinctAnswers()

	assert.Equal(t, 1, g4)

	g5 := summary.groups[4].getDistinctAnswers()

	assert.Equal(t, 1, g5)

	sum := summary.getTotalDistinctAnswers()

	assert.Equal(t, 11, sum)
}
