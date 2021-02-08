package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	answerLines, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Error reading file (%s)", err)
	}

	// this is a set type of characters for a group. we don't
	// double count characters, so we just want to know if the
	// group answered something at least once. here's an example:
	//
	//	{
	//		a, b, f, z
	//	}
	//
	// In this example, the group answered with 'a', 'b', 'f' and 'z'
	answerSet := map[rune]struct{}{}
	answerSum := 0
	for _, answerLine := range strings.Split(string(answerLines), "\n") {
		// if there's an empty line, then this group's answers are complete.
		// count them up and reset the answer counter
		if answerLine == "" {
			answerSum += len(answerSet)
			answerSet = map[rune]struct{}{}
		} else {
			// go through all the answers on this line
			// and record them in the answerSet
			for _, answerChar := range answerLine {
				// here is where we add the character to the set.
				// this is effectively a no-op if someone in the group
				// has already answered this question
				answerSet[answerChar] = struct{}{}
			}
		}
	}
	log.Printf("Sum of all the groups: %d", answerSum)
}
