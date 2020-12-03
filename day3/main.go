package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {

	var version = flag.Int("version", 1, "first or second part of the assignment")

	flag.Parse()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		l := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, l)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	grid := parseGrid(lines)

	log.Printf("Version: %d ,Result: %d 🎄", *version, treesHit(grid, 1, 3))

	tress2 := treesHit(grid, 1, 1)
	tress := treesHit(grid, 1, 3)
	tress3 := treesHit(grid, 1, 5)
	tress4 := treesHit(grid, 1, 7)
	tress5 := treesHit(grid, 2, 1)
	multi := tress * tress2 * tress3 * tress4 * tress5
	log.Printf("Version: %d ,Result: %d 🎄", 2, multi)
}

func parseGrid(rows []string) map[int]map[int]bool {
	grid := map[int]map[int]bool{}

	for i, r := range rows {
		grid[i] = map[int]bool{}
		for i2, c := range r {
			tree := false
			if c == '#' {
				tree = true
			}
			grid[i][i2] = tree
		}
	}

	return grid
}

func treesHit(grid map[int]map[int]bool, down, right int) int {
	top := 0
	left := 0
	trees := 0
	for {
		top += down
		left += right

		if _, ok := grid[top]; !ok {
			return trees
		}

		row := grid[top]

		if left >= len(row) {
			left = left - (len(row))
		}

		col := row[left]

		if col {
			trees++
		}

	}

}
