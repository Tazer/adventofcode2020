package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Version1(t *testing.T) {

	lines := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	grid := parseGrid(lines)

	assert.Equal(t, true, grid[0][3])
	assert.Equal(t, true, grid[10][1])

	tress2 := treesHit(grid, 1, 1)

	assert.Equal(t, 2, tress2)

	tress := treesHit(grid, 1, 3)

	assert.Equal(t, 7, tress)

	tress3 := treesHit(grid, 1, 5)

	assert.Equal(t, 3, tress3)

	tress4 := treesHit(grid, 1, 7)

	assert.Equal(t, 4, tress4)

	tress5 := treesHit(grid, 2, 1)

	assert.Equal(t, 2, tress5)

	multi := tress * tress2 * tress3 * tress4 * tress5

	assert.Equal(t, 336, multi)

}

func TestDay2Version2(t *testing.T) {

}
