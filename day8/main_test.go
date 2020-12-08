package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Version1(t *testing.T) {
	lines := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	instructions := parseIntstructions(lines)

	assert.Equal(t, "nop", instructions[0].op)
	assert.Equal(t, "acc", instructions[1].op)
	assert.Equal(t, 1, instructions[1].value)

	res, _ := runProgram(instructions)

	assert.Equal(t, 5, res)

	res2, _ := executeInstructions(instructions, instructions, 0)

	assert.Equal(t, 8, res2)

}
