package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input        string
	Boardingpass boardingpass
	SeatID       int
}

func TestDay2Version1(t *testing.T) {
	cases := []testCase{
		{Input: "FBFBBFFRLR",
			Boardingpass: boardingpass{row: 44, column: 5},
			SeatID:       357,
		},
		{Input: "BFFFBBFRRR",
			Boardingpass: boardingpass{row: 70, column: 7},
			SeatID:       567,
		},
		{Input: "FFFBBBFRRR",
			Boardingpass: boardingpass{row: 14, column: 7},
			SeatID:       119,
		},
		{Input: "BBFFBBFRLL",
			Boardingpass: boardingpass{row: 102, column: 4},
			SeatID:       820,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			b := newBoardingpass(c.Input)

			assert.Equal(t, c.Boardingpass.row, b.row)
			assert.Equal(t, c.Boardingpass.column, b.column)
			assert.Equal(t, c.SeatID, b.getSeatID())
		})
	}
}
