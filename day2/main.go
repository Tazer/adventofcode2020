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
	if *version == 1 {
		log.Printf("Result: %d 🎄", checkPasswordPolicy(inputs))
	}
	// } else {
	// 	log.Printf("Result: %d 🎄", calculateAccoutingVersion2(inputs))
	// }
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

func checkPasswordPolicy(input []passwordPolicy) int {

	valid := 0
	for _, p := range input {
		if p.Valid() {
			valid++
		}
	}

	return valid
}
