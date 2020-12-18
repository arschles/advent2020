package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type passwordInfo struct {
	min, max int
	letter   string
	password string
}

func (pwi *passwordInfo) isValid() bool {
	if pwi == nil {
		return false
	}
	hasMin := string(pwi.password[pwi.min-1]) == pwi.letter
	hasMax := string(pwi.password[pwi.max-1]) == pwi.letter
	// 	    min  max
	//  min  T    F
	//  max  F    T
	return hasMin != hasMax
}

func main() {
	inputBytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %s", err)
	}
	// turn each line in the file into a []string
	lines := strings.Split(string(inputBytes), "\n")

	// iterate through each line
	numValid := 0
	for _, line := range lines {
		pwi, err := parsePasswordInfo(line)
		if err != nil {
			continue
		}
		if pwi.isValid() {
			numValid++
		}
	}

	log.Printf("num passwords valid: %d", numValid)
	// 1-3 a: whatever
	// <min>-<max> <letter>: <password>
}

func parsePasswordInfo(line string) (*passwordInfo, error) {
	minMax := strings.Split(line, "-")
	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		return nil, err
	}
	remainder := strings.Split(minMax[1], " ")
	max, err := strconv.Atoi(remainder[0])
	if err != nil {
		return nil, err
	}
	password := remainder[2]
	letter := strings.TrimRight(remainder[1], ":")
	return &passwordInfo{
		min:      min,
		max:      max,
		letter:   letter,
		password: password,
	}, nil
}
