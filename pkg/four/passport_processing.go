package four

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	pathToData = "input/four.txt"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

// Passport ...
type Passport struct {
	FieldsJSON map[string]string
}

// Run ...
func Run() (int, error) {
	numValidPassports := 0
	passports, err := ParseInput(pathToData)
	if err != nil {
		return -1, err
	}
	for _, passport := range passports {
		if passport.Valid(requiredFields) {
			numValidPassports++
		}
	}
	return numValidPassports, nil
}

// ParseInput ...
func ParseInput(filepath string) ([]Passport, error) {
	allPassports := []Passport{}
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return allPassports, err
	}

	rawPassports := strings.Split(string(data), "\n\n")
	for _, p := range rawPassports {
		passport := FormatPassport(p)
		allPassports = append(allPassports, passport)
	}

	return allPassports, nil
}

// FormatPassport ...
func FormatPassport(raw string) Passport {
	passport := Passport{
		FieldsJSON: map[string]string{},
	}
	allfields := strings.FieldsFunc(raw, Sep)
	for _, field := range allfields {
		f := strings.Split(field, ":")
		passport.FieldsJSON[f[0]] = f[1]
	}
	return passport
}

// Valid ...
func (p Passport) Valid(required []string) bool {
	// Invalid if any field in list of required fields is missing
	for _, req := range required {
		if _, ok := p.FieldsJSON[req]; !ok {
			return false
		}
	}
	// Validate each field once confirming each is present
	for _, req := range required {
		if val, ok := p.FieldsJSON[req]; ok {
			switch req {
			case "byr":
				year, err := strconv.Atoi(val)
				if err != nil || year < 1920 || year > 2002 {
					return false
				}
			case "iyr":
				year, err := strconv.Atoi(val)
				if err != nil || year < 2010 || year > 2020 {
					return false
				}
			case "eyr":
				year, err := strconv.Atoi(val)
				if err != nil || year < 2020 || year > 2030 {
					return false
				}
			case "hgt":
				// The last two characters have to represent the metrics
				metric := val[len(val)-2:]
				matched, err := regexp.MatchString("cm|in", metric)
				if err != nil || !matched {
					return false
				}
				height, err := strconv.Atoi(val[:len(val)-2]) // subslice inclusive?
				if err != nil {
					return false
				}
				if metric == "cm" {
					if height < 150 || height > 193 {
						return false
					}
				} else if metric == "in" {
					if height < 59 || height > 76 {
						return false
					}
				}
			case "ecl":
				matched, err := regexp.MatchString("amb|blu|brn|gry|grn|hzl|oth", val)
				if err != nil || !matched {
					return false
				}
			case "hcl":
				matched, err := regexp.MatchString("^#[a-f0-9]{6}$", val)
				if err != nil || !matched {
					return false
				}
			case "pid":
				matched, err := regexp.MatchString("^[0-9]{9}$", val)
				if err != nil || !matched {
					return false
				}
			case "cid":
				// whoops we forgot to check this one
			}
		}
	}
	return true
}

// Sep - to use multiple separators
// Reference: https://stackoverflow.com/questions/39862613/how-to-split-a-string-by-multiple-delimiters
func Sep(s rune) bool {
	return s == ' ' || s == '\n'
}
