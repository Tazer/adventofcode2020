package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input          string
	PasswordPolicy passwordPolicy
	Valid          bool
}

func TestDay2Version1(t *testing.T) {
	cases := []testCase{
		{Input: "1-3 a: abcde",
			PasswordPolicy: passwordPolicy{Password: "abcde", Char: "a", Min: 1, Max: 3}, Valid: true},
		{Input: "1-3 b: cdefg",
			PasswordPolicy: passwordPolicy{Password: "cdefg", Char: "b", Min: 1, Max: 3}, Valid: false},
		{Input: "2-9 c: ccccccccc",
			PasswordPolicy: passwordPolicy{Password: "ccccccccc", Char: "c", Min: 2, Max: 9}, Valid: true},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			p := parseInput(c.Input)
			log.Printf("Password policy %+v", p)
			assert.Equal(t, c.PasswordPolicy.Char, p.Char)
			assert.Equal(t, c.PasswordPolicy.Min, p.Min)
			assert.Equal(t, c.PasswordPolicy.Max, p.Max)
			assert.Equal(t, c.Valid, p.Valid())
		})

	}
}

func TestDay2Version2(t *testing.T) {
	cases := []testCase{
		{Input: "1-3 a: abcde",
			PasswordPolicy: passwordPolicy{Password: "abcde", Char: "a", Min: 1, Max: 3}, Valid: true},
		{Input: "1-3 b: cdefg",
			PasswordPolicy: passwordPolicy{Password: "cdefg", Char: "b", Min: 1, Max: 3}, Valid: false},
		{Input: "2-9 c: ccccccccc",
			PasswordPolicy: passwordPolicy{Password: "ccccccccc", Char: "c", Min: 2, Max: 9}, Valid: false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			p := parseInput(c.Input)
			log.Printf("Password policy %+v", p)
			assert.Equal(t, c.PasswordPolicy.Char, p.Char)
			assert.Equal(t, c.PasswordPolicy.Min, p.Min)
			assert.Equal(t, c.PasswordPolicy.Max, p.Max)
			assert.Equal(t, c.Valid, p.ValidVersion2())
		})

	}
}
