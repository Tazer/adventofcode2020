package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay2Version1(t *testing.T) {

	lines := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	}

	passports := parsePassports(lines)

	p := passports[0]

	assert.Equal(t, "860033327", p.PassportID)

	valid := validPassports(passports, 1)

	assert.Equal(t, 2, valid)

}

func TestYearParser(t *testing.T) {
	p := passport{BirthDate: "2002"}

	v := validYear(p.BirthDate, 1920, 2002)

	assert.Equal(t, true, v)

	p2 := passport{BirthDate: "2003"}

	v2 := validYear(p2.BirthDate, 1920, 2002)

	assert.Equal(t, false, v2)
}

func TestHeight(t *testing.T) {
	p := passport{Height: "60in"}

	v := validHeight(p.Height)

	assert.Equal(t, true, v)

	p2 := passport{Height: "190cm"}

	v2 := validHeight(p2.Height)

	assert.Equal(t, true, v2)

	p3 := passport{Height: "190in"}

	v3 := validHeight(p3.Height)

	assert.Equal(t, false, v3)

	p4 := passport{Height: "190"}

	v4 := validHeight(p4.Height)

	assert.Equal(t, false, v4)
}

func TestHeightFailed(t *testing.T) {
	p := passport{Height: "200in"}

	v := validHeight(p.Height)
	assert.Equal(t, false, v)
}

func TestHairColor(t *testing.T) {
	p := passport{HairColor: "#000fff"}

	v := validHairColor(p.HairColor)
	assert.Equal(t, true, v)

	p2 := passport{HairColor: "#000ffF"}

	v2 := validHairColor(p2.HairColor)
	assert.Equal(t, false, v2)
}

func TestPassportID(t *testing.T) {
	p := passport{PassportID: "000000001"}

	v := validPassportID(p.PassportID)
	assert.Equal(t, true, v)

	p2 := passport{PassportID: "0123456789"}

	v2 := validPassportID(p2.PassportID)
	assert.Equal(t, false, v2)
}

func TestDay2Version2Invalid(t *testing.T) {

	lines := []string{
		"eyr:1972 cid:100",
		"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
		"",
		"iyr:2019",
		"hcl:#602927 eyr:1967 hgt:170cm",
		"ecl:grn pid:012533040 byr:1946",
		"",
		"hcl:dab227 iyr:2012",
		"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
		"",
		"hgt:59cm ecl:zzz",
		"eyr:2038 hcl:74454a iyr:2023",
		"pid:3556412378 byr:2007",
	}

	passports := parsePassports(lines)

	p := passports[0]

	assert.Equal(t, "186cm", p.PassportID)

	valid := validPassports(passports, 2)

	assert.Equal(t, 0, valid)

}

func TestDay2Version2Valid(t *testing.T) {

	lines := []string{
		"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
		"hcl:#623a2f",
		"",
		"eyr:2029 ecl:blu cid:129 byr:1989",
		"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
		"",
		"hcl:#888785",
		"hgt:164cm byr:2001 iyr:2015 cid:88",
		"pid:545766238 ecl:hzl",
		"eyr:2022",
		"",
		"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
	}

	passports := parsePassports(lines)

	p := passports[0]

	assert.Equal(t, "087499704", p.PassportID)

	valid := validPassports(passports, 2)

	assert.Equal(t, 4, valid)

}
