package four

import "testing"

func TestValid(t *testing.T) {
	tests := []struct {
		name     string
		passport Passport
		expected bool
	}{
		{
			name: "invalid path - height metric not correct",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "190hm",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - height value not valid",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "!50cm",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - missing iyr",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "60in",
					"eyr": "2021",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - missing eyr",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "60in",
					"iyr": "2011",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "valid path - all fields present and valid, height in cm",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "190cm",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: true,
		},
		{
			name: "valid path - all fields present and valid, height in inches",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "70in",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: true,
		},
		{
			name: "invalid path - byr too high",
			passport: Passport{
				map[string]string{
					"byr": "2003",
					"hgt": "60in",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - byr too low",
			passport: Passport{
				map[string]string{
					"byr": "1902",
					"hgt": "60in",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - iyr too high",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "70in",
					"iyr": "2030",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - iyr too low",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "70in",
					"iyr": "2009",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - eyr too high",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "70in",
					"iyr": "2005",
					"eyr": "2040",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - eyr too low",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "70in",
					"iyr": "2010",
					"eyr": "2019",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - hgt inches too tall",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "190in",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - hgt inches too short",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "10in",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - hgt cm too short",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "20cm",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - hgt cm too tall",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "500cm",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - hcl incorrect",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "160cm",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "123abc",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - hcl incorrect alternative",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "160cm",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abz",
					"ecl": "brn",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - ecl incorrect",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "160cm",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "wat",
					"pid": "000000001",
				},
			},
			expected: false,
		},
		{
			name: "invalid path - pid incorrect",
			passport: Passport{
				map[string]string{
					"byr": "2002",
					"hgt": "160cm",
					"iyr": "2010",
					"eyr": "2020",
					"hcl": "#123abc",
					"ecl": "grn",
					"pid": "0123456789",
				},
			},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.passport.Valid(requiredFields)
			if actual != tt.expected {
				t.Errorf("expected %t but got %t", tt.expected, actual)
			}
		})
	}
}

func TestFormatPassport(t *testing.T) {
	tests := []struct {
		name     string
		rawInput string
		expected Passport
	}{
		{
			name:     "happy path",
			rawInput: "iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
			expected: Passport{
				FieldsJSON: map[string]string{
					"iyr": "2010",
					"hgt": "158cm",
					"hcl": "#b6652a",
					"ecl": "blu",
					"byr": "1944",
					"eyr": "2021",
					"pid": "093154719",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FormatPassport(tt.rawInput)
			for wantKey, wantVal := range tt.expected.FieldsJSON {
				if _, ok := actual.FieldsJSON[wantKey]; !ok {
					t.Errorf("expected key %s to be present in formatted passport but didn't get any", wantKey)
				}
				gotVal := actual.FieldsJSON[wantKey]
				if wantVal != gotVal {
					t.Errorf("expected key:value %s:%s but got key:value %s:%s", wantKey, wantVal, wantKey, gotVal)
				}
			}
		})
	}
}
