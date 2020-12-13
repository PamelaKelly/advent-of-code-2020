package four

import (
	"io/ioutil"
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
	Fields map[string]string
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
		Fields: map[string]string{},
	}
	allfields := strings.FieldsFunc(raw, Sep)
	for _, field := range allfields {
		f := strings.Split(field, ":")
		passport.Fields[f[0]] = f[1]
	}
	return passport
}

// Valid ...
func (p Passport) Valid(required []string) bool {
	// Invalid if any field in list of required fields is missing
	for _, req := range required {
		if _, ok := p.Fields[req]; !ok {
			return false
		}
	}
	return true
}

// Sep - to use multiple separators
// Reference: https://stackoverflow.com/questions/39862613/how-to-split-a-string-by-multiple-delimiters
func Sep(s rune) bool {
	return s == ' ' || s == '\n'
}
