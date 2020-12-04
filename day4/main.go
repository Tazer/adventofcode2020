package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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

	passports := parsePassports(lines)

	log.Printf("Version: %d ,Result: %d 🎄", *version, validPassports(passports, *version))
}

func validPassports(passports []passport, version int) int {
	valid := 0
	for _, p := range passports {
		if version == 2 {
			if p.valid2() {
				valid++
			}
		} else {
			if p.valid() {
				valid++
			}
		}
	}
	return valid
}

func parsePassports(rows []string) []passport {
	passports := []passport{}

	p := passport{}

	for _, l := range rows {

		if l == "" {
			passports = append(passports, p)
			p = passport{}
			continue
		}

		kvs := strings.Split(l, " ")

		for _, kv := range kvs {
			s := strings.Split(kv, ":")
			k := s[0]
			v := s[1]

			switch k {
			case "byr":
				p.BirthDate = v
			case "iyr":
				p.IssueYear = v
			case "eyr":
				p.ExpirationYear = v
			case "hgt":
				p.Height = v
			case "hcl":
				p.HairColor = v
			case "ecl":
				p.EyeColor = v
			case "pid":
				p.PassportID = v
			case "cid":
				p.CountryID = v

			}
		}
	}

	passports = append(passports, p)

	return passports
}

type passport struct {
	BirthDate      string `json:"byr"`
	IssueYear      string `json:"iyr"`
	ExpirationYear string `json:"eyr"`
	Height         string `json:"hgt"`
	HairColor      string `json:"hcl"`
	EyeColor       string `json:"ecl"`
	PassportID     string `json:"pid"`
	CountryID      string `json:"cid"`
}

func (p *passport) valid2() bool {

	if !p.valid() {
		return false
	}

	if !validYear(p.BirthDate, 1920, 2002) {
		return false
	}

	if !validYear(p.IssueYear, 2010, 2020) {
		return false
	}

	if !validYear(p.ExpirationYear, 2020, 2030) {
		return false
	}

	if !validHeight(p.Height) {
		return false
	}

	if !validHairColor(p.HairColor) {
		return false
	}

	if p.EyeColor != "amb" && p.EyeColor != "blu" && p.EyeColor != "brn" && p.EyeColor != "gry" && p.EyeColor != "grn" && p.EyeColor != "hzl" && p.EyeColor != "oth" {
		return false
	}

	if !validPassportID(p.PassportID) {
		return false
	}

	return true
}

func validPassportID(input string) bool {
	regex := *regexp.MustCompile(`^[0-9]{9}$`)

	return regex.MatchString(input)
}

func validHairColor(input string) bool {
	regex := *regexp.MustCompile(`^#[a-f0-9]{6}$`)

	return regex.MatchString(input)
}

func validHeight(input string) bool {
	m := string(input[len(input)-2:])
	v, _ := strconv.Atoi(input[0 : len(input)-2])

	log.Printf("m is %s and v is %d", m, v)

	if m == "cm" {
		if v > 193 {
			return false
		}
		if v < 150 {
			return false
		}
	} else if m == "in" {
		if v > 76 {
			return false
		}
		if v < 59 {
			return false
		}
	} else {
		return false
	}
	return true
}

func validYear(input string, min, max int) bool {
	d, err := time.Parse(time.RFC3339, input+"-01-01T00:00:00Z")

	if err != nil {
		log.Printf("err %v", err)
		return false
	}
	if d.Year() < min {
		return false
	}

	if d.Year() > max {
		return false
	}
	return true
}

func (p *passport) valid() bool {

	if p.BirthDate == "" {
		return false
	}

	if p.IssueYear == "" {
		return false
	}

	if p.ExpirationYear == "" {
		return false
	}

	if p.Height == "" {
		return false
	}

	if p.HairColor == "" {
		return false
	}

	if p.EyeColor == "" {
		return false
	}

	if p.PassportID == "" {
		return false
	}

	// if p.CountryID == "" {
	// 	return false
	// }
	return true

}
