package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
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

	boardingpasses := parseBoardingpasses(lines)

	log.Printf("Version: %d ,Result: %d 🎄", *version, getHighestSeatID(boardingpasses))

	log.Printf("My seat is %d", getMySeatID(boardingpasses))
}

func parseBoardingpasses(input []string) []boardingpass {
	bp := []boardingpass{}

	for _, i := range input {
		b := newBoardingpass(i)
		bp = append(bp, b)
	}
	return bp
}

func getMySeatID(bpasses []boardingpass) int {
	seatIDs := []int{}

	for _, b := range bpasses {
		seatIDs = append(seatIDs, b.getSeatID())
	}

	sort.Ints(seatIDs)

	start := 0
	for i, s := range seatIDs {
		if i == 0 {
			start = s
		} else {
			start++
			if s != start {
				return s - 1
			}
		}
	}

	return -1
}

func getHighestSeatID(bpasses []boardingpass) int {
	seatID := 0

	for _, b := range bpasses {
		if b.getSeatID() > seatID {
			seatID = b.getSeatID()
		}
	}
	return seatID
}

func newBoardingpass(input string) boardingpass {
	b := boardingpass{}
	rowMin := 0
	rowMax := 127

	row := 0

	for index, i := range input[:7] {
		row = rowMin + (rowMax-rowMin)/2
		if i == 'F' {
			rowMax = row
			if index == len(input[:7])-1 {
				row = rowMax
			}
		}

		if i == 'B' {
			rowMin = row + 1
			if index == len(input[:7])-1 {
				row = rowMin
			}
		}
	}

	columnMin := 0
	columnMax := 7

	column := 0

	for index, i := range input[7:] {
		column = columnMin + (columnMax-columnMin)/2
		if i == 'R' {
			columnMin = column + 1
			if index == len(input[:7])-1 {
				column = columnMin
			}
		}

		if i == 'L' {
			columnMax = column
			if index == len(input[:7])-1 {
				column = columnMax
			}
		}
	}

	column = columnMin + (columnMax-columnMin)/2

	b.row = row
	b.column = column
	return b
}

type boardingpass struct {
	row    int
	column int
}

func (b *boardingpass) getSeatID() int {
	return b.row*8 + b.column
}
