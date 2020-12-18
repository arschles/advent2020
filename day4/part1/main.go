package main

import (
	"io/ioutil"
	"log"
	"strings"
)

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
		keyValuePairs := strings.Split(fieldsStr, " ")
		counters := map[string]int{}
		for _, keyValuePair := range keyValuePairs {
			keyAndValue := strings.Split(keyValuePair, ":")
			counters[keyAndValue[0]]++
		}

		numInvalid := 0
		for _, requiredField := range requiredFields {
			if counters[requiredField] != 1 {
				numInvalid++
			}
		}
		if numInvalid == 0 {
			numValid++
		}
	}
	log.Printf("Valid passports: %d", numValid)
}
