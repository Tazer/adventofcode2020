package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	inputs := []passwordPolicy{}

	for scanner.Scan() {
		p := parseInput(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Version: %d ,Result: %d 🎄", *version, checkPasswordPolicy(inputs, *version))
}

type passwordPolicy struct {
	Password string
	Char     string
	Min      int
	Max      int
}

func (p *passwordPolicy) Valid() bool {

	c := strings.Count(p.Password, p.Char)

	if c > p.Max {
		return false
	}

	if c < p.Min {
		return false
	}

	return true
}

func (p *passwordPolicy) ValidVersion2() bool {
	if p.Char == string(p.Password[p.Max-1]) && p.Char != string(p.Password[p.Min-1]) {
		return true
	}

	if p.Char == string(p.Password[p.Min-1]) && p.Char != string(p.Password[p.Max-1]) {
		return true
	}
	return false
}

func parseInput(input string) passwordPolicy {

	regex := *regexp.MustCompile(`([0-9]+)-([0-9]+) (\w): (\w+)`)
	res := regex.FindAllStringSubmatch(input, -1)

	min, _ := strconv.Atoi(res[0][1])
	max, _ := strconv.Atoi(res[0][2])

	return passwordPolicy{
		Min:      min,
		Max:      max,
		Char:     res[0][3],
		Password: res[0][4],
	}
}

func checkPasswordPolicy(input []passwordPolicy, version int) int {

	valid := 0
	for _, p := range input {
		if version == 2 {
			if p.ValidVersion2() {
				valid++
			}
		} else {
			if p.Valid() {
				valid++
			}
		}
	}

	return valid
}
