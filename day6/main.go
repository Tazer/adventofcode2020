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

	sum := parseGroups(lines)

	log.Printf("Version: %d ,Result: %d , Result v2: %d🎄", *version, sum.getTotalDistinctAnswers(), sum.getTotalAnswersEveryone())

}

func parseGroups(input []string) summary {
	s := summary{
		groups: []group{},
	}

	g := group{}

	for _, i := range input {
		if i == "" {
			s.groups = append(s.groups, g)
			g = group{}
			continue
		}

		p := person{}
		for _, c := range i {
			p.answers = append(p.answers, string(c))
		}
		g.persons = append(g.persons, p)
	}
	s.groups = append(s.groups, g)

	return s
}

type (
	summary struct {
		groups []group
	}
	group struct {
		persons []person
	}
	person struct {
		answers []string
	}
)

func (g *group) getDistinctAnswers() int {
	m := map[string]bool{}

	for _, p := range g.persons {
		for _, a := range p.answers {
			if _, ok := m[a]; !ok {
				m[a] = true
			}
		}
	}
	return len(m)
}

func (g *group) getAnswersEveryone() int {
	m := map[string]int{}

	for _, p := range g.persons {
		for _, a := range p.answers {
			m[a]++
		}
	}

	answers := 0

	for _, v := range m {
		if v == len(g.persons) {
			answers++
		}
	}

	return answers
}

func (s *summary) getTotalDistinctAnswers() int {
	total := 0
	for _, g := range s.groups {
		total += g.getDistinctAnswers()
	}
	return total
}

func (s *summary) getTotalAnswersEveryone() int {
	total := 0
	for _, g := range s.groups {
		total += g.getAnswersEveryone()
	}
	return total
}
