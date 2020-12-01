package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {

	sample := []int{1721, 979, 366, 299, 675, 1456}
	expected := 514579

	result := calculateAccouting(sample)

	assert.Equal(t, result, expected)

}

func TestDay1Version2(t *testing.T) {

	sample := []int{1721, 979, 366, 299, 675, 1456}
	expected := 241861950

	result := calculateAccoutingVersion2(sample)

	assert.Equal(t, result, expected)

}
