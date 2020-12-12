package two

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	pathToData = "input/two.txt"
)

// Entry ...
type Entry struct {
	Min      int
	Max      int
	Letter   string
	Password string
}

// Valid ...
func (r Entry) Valid() bool {
	occurrences := 0
	for _, c := range r.Password {
		if string(c) == r.Letter {
			occurrences++
		}
	}
	if occurrences < r.Min || occurrences > r.Max {
		return false
	}
	return true
}

// Run ...
func Run() (int, error) {
	entries, err := ParseInput(pathToData)
	if err != nil {
		return -1, err
	}
	return CountValidPasswords(entries), nil
}

// CountValidPasswords ...
func CountValidPasswords(entries []Entry) int {
	validPasswords := 0
	for _, entry := range entries {
		if entry.Valid() {
			validPasswords++
		}
	}
	return validPasswords
}

// ParseInput ...
func ParseInput(filepath string) ([]Entry, error) {
	entries := []Entry{}
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	if err != nil {
		return nil, err
	}
	for _, line := range lines {
		if line != "" {
			items := strings.Split(line, " ")
			min, err := strconv.Atoi(strings.Split(items[0], "-")[0])
			if err != nil {
				return nil, err
			}
			max, err := strconv.Atoi(strings.Split(items[0], "-")[1])
			if err != nil {
				return nil, err
			}
			newEntry := Entry{
				Min:      min,
				Max:      max,
				Letter:   strings.TrimRight(items[1], ":"),
				Password: items[2],
			}
			entries = append(entries, newEntry)
		}
	}
	return entries, nil
}
