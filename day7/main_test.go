package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Version1(t *testing.T) {
	lines := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}

	rules := parseRules(lines)

	assert.Equal(t, "light red", rules[0].bag)
	assert.Equal(t, "bright white", rules[0].bags[0].name)
	assert.Equal(t, 1, rules[0].bags[0].count)

	assert.Equal(t, "muted yellow", rules[0].bags[1].name)
	assert.Equal(t, 2, rules[0].bags[1].count)

	i := "shiny gold"

	res := findpossiblebags(i, rules)

	assert.Equal(t, 4, len(res))

}

func TestDay7Version2(t *testing.T) {
	lines := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}

	rules := parseRules(lines)

	res := calculateNumberOfBags("shiny gold", rules)
	assert.Equal(t, 32, res)
}

func TestDay7Version2_2(t *testing.T) {
	lines := []string{
		"shiny gold bags contain 2 dark red bags.",
		"dark red bags contain 2 dark orange bags.",
		"dark orange bags contain 2 dark yellow bags.",
		"dark yellow bags contain 2 dark green bags.",
		"dark green bags contain 2 dark blue bags.",
		"dark blue bags contain 2 dark violet bags.",
		"dark violet bags contain no other bags.",
	}

	rules := parseRules(lines)

	res := calculateNumberOfBags("shiny gold", rules)
	assert.Equal(t, 126, res)
}
