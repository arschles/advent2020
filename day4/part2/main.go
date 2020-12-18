package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type field struct {
	field, value string
}

func isInRange(val, min, max int) bool {
	return val >= min && val <= max
}

func (f field) isValid() bool {
	switch f.field {
	case "byr":
		val, err := strconv.Atoi(f.value)
		if err != nil {
			return false
		}
		return isInRange(val, 1920, 2002)
	case "iyr":
		val, err := strconv.Atoi(f.value)
		if err != nil {
			return false
		}
		return isInRange(val, 2010, 2020)
	case "eyr":
		val, err := strconv.Atoi(f.value)
		if err != nil {
			return false
		}
		return isInRange(val, 2020, 2030)
	case "hgt":
		if strings.HasSuffix(f.value, "cm") {
			cmVal, err := strconv.Atoi(strings.TrimRight(f.value, "cm"))
			if err != nil {
				return false
			}
			return isInRange(cmVal, 150, 193)
		} else if strings.HasSuffix(f.value, "in") {
			inVal, err := strconv.Atoi(strings.TrimRight(f.value, "in"))
			if err != nil {
				return false
			}
			return isInRange(inVal, 59, 76)
		} else {
			return false
		}
	case "hcl":
		if f.value[0] != '#' {
			return false
		}
		hexStr := f.value[1:]
		// parsing to hex will make sure the hair color is valid
		_, err := strconv.ParseInt(hexStr, 16, 64)
		if err != nil {
			return false
		}
	case "ecl":
		colors := map[string]bool{
			"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true,
		}
		if _, ok := colors[f.value]; !ok {
			return false
		}
	case "pid":
		if len(f.value) != 9 {
			return false
		}
		_, err := strconv.Atoi(f.value)
		if err != nil {
			return false
		}
	default:
		return false
	}

	return true
}

func main() {
	fileBytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
	fileSplit := strings.Split(string(fileBytes), "\n\n")
	numValid := 0
	for _, passport := range fileSplit {
		fieldsStr := strings.Join(strings.Split(passport, "\n"), " ")
		fieldsStr = strings.TrimRight(fieldsStr, " ")
		keyValuePairs := strings.Split(fieldsStr, " ")
		counters := map[string]int{}
		fields := map[string]field{}
		for _, keyValuePair := range keyValuePairs {
			keyAndValue := strings.Split(keyValuePair, ":")
			counters[keyAndValue[0]]++
			fields[keyAndValue[0]] = field{
				field: keyAndValue[0],
				value: keyAndValue[1],
			}
		}

		numInvalid := 0
		for _, requiredField := range requiredFields {
			if counters[requiredField] != 1 {
				numInvalid++
			}
			if !fields[requiredField].isValid() {
				numInvalid++
			}
		}
		if numInvalid == 0 {
			numValid++
		}
	}
	log.Printf("Valid passports: %d", numValid)
}
