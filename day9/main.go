package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
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

	lines := []int{}

	for scanner.Scan() {
		l := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		iL, _ := strconv.Atoi(l)

		lines = append(lines, iL)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	res := getInvalidNumber(lines, 25)

	log.Printf("Version: %d ,Result: %d , Result v2: %d 🎄", *version, res, findWeakness(lines, res))

}

func getInvalidNumber(input []int, preambleSize int) int {
	for index, i := range input[preambleSize:] {
		match := false
		for _, p := range input[index : preambleSize+index] {
			for _, p2 := range input[index : preambleSize+index] {
				if p != p2 && p+p2 == i {
					match = true
				}
			}
		}
		if !match {
			return i
		}
	}

	return 0
}

func findWeakness(input []int, weakNumber int) int {
	numbers := []int{}
	found := false
	for i, p := range input {
		numbers = append(numbers, p)
		for _, p2 := range input[i+1:] {
			numbers = append(numbers, p2)
			s := checkSum(numbers)
			if s == weakNumber {
				found = true
				break
			}
		}
		if found {
			break
		}
		numbers = []int{}
	}

	if !found {
		return 0
	}

	return lowestAndHigestSum(numbers)
}

func lowestAndHigestSum(input []int) int {
	sort.Ints(input)

	low := input[0]
	high := input[len(input)-1]
	return low + high
}

func checkSum(input []int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
}
