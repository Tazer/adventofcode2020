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

	log.Print(lines[0])

	rules := parseRules(lines)

	res := findpossiblebags("shiny gold", rules)
	log.Printf("Version: %d ,Result: %d Result2: %d 🎄", *version, len(res), calculateNumberOfBags("shiny gold", rules))

}

func calculateNumberOfBags(bag string, rules []rule) int {
	r := findRule(bag, rules)
	bags := 0
	for _, b := range r.bags {
		bags += b.count
		r2 := findRule(b.name, rules)
		if len(r2.bags) > 0 {
			bags += b.count * calculateNumberOfBags(b.name, rules)
		}
	}

	return bags
}

func findRule(bag string, rules []rule) *rule {
	for _, r := range rules {
		if r.bag == bag {
			return &r
		}
	}
	return nil
}

func parseRules(lines []string) []rule {
	rs := []rule{}

	for _, l := range lines {
		r := rule{}

		cl := strings.ReplaceAll(l, "bags", "")
		cl = strings.ReplaceAll(cl, "bag", "")
		cl = strings.ReplaceAll(cl, ".", "")
		cl = strings.ReplaceAll(cl, "no other", "")

		arr := regexp.MustCompile("(contain)|,").Split(cl, -1)

		r.bag = strings.TrimSpace(arr[0])

		for _, b := range arr[1:] {
			b = strings.TrimSpace(b)
			if b == "" {
				continue
			}
			bArr := strings.Split(b, " ")

			count, _ := strconv.Atoi(bArr[0])

			name := b[strings.Index(b, " "):]

			bag := bag{
				name:  strings.TrimSpace(name),
				count: count,
			}
			r.bags = append(r.bags, bag)
		}
		rs = append(rs, r)
	}

	return rs
}

func findpossiblebags(bag string, rules []rule) map[string]int {

	bags := map[string]int{}

	for _, r := range rules {
		if r.bag == bag && r.bag != "shiny gold" {
			bags[r.bag] = 1
		}

		if len(r.bags) > 0 {
			bagFound := false
			for _, b := range r.bags {
				if b.name == bag {
					bags[r.bag] = b.count
					bagFound = true
				}

			}
			if bagFound {
				res2 := findpossiblebags(r.bag, rules)

				for k, v := range res2 {
					bags[k] = v
				}
			}
		}

	}

	return bags
}

type rule struct {
	bag  string
	bags []bag
}

type bag struct {
	name  string
	count int
}
