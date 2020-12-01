package main

import (
	"bufio"
	"flag"
	"log"
	"os"
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

	inputs := []int{}

	for scanner.Scan() {
		i1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, i1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if *version == 1 {
		log.Printf("Result: %d 🎄", calculateAccouting(inputs))
	} else {
		log.Printf("Result: %d 🎄", calculateAccoutingVersion2(inputs))
	}
}

func calculateAccouting(input []int) int {

	for _, i := range input {
		for _, i2 := range input {
			if i+i2 == 2020 {
				return i * i2
			}
		}
	}
	return 0
}

func calculateAccoutingVersion2(input []int) int {

	for _, i := range input {
		for _, i2 := range input {
			for _, i3 := range input {
				if i+i2+i3 == 2020 {
					return i * i2 * i3
				}
			}
		}
	}
	return 0
}
