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

	// this is a grouping of the answers for the group of people.
	// for each answer -- which is a key in the map -- there is a counter
	// that represents the number of people in the group that answered
	// yes for it.
	//
	// if the counter is equal to the number of people in the group,
	// that means that everyone in the group answered yes for it.
	// here is an example for a group of 2 people that answered the following:
	//
	//	ab
	//	ac
	//
	//	{
	//		a: 2,
	//		b: 1,
	//		c: 1,
	//	}
	//
	// In this example, both people answered yes for 'a', one of them answered
	// yes for 'b', and the other person answered yes for 'c'
	answerSet := map[rune]int{}
	answerSum := 0
	numPeopleInGroup := 0
	for _, answerLine := range strings.Split(string(answerLines), "\n") {
		// if there's an empty line, then this group's answers are complete.
		// count them up and reset the answer counter
		if answerLine == "" {
			log.Printf("End of group of %d people", numPeopleInGroup)
			for answer, numAnswerers := range answerSet {
				log.Printf("Num answerers for '%s': %d", string(answer), numAnswerers)
				if numAnswerers == numPeopleInGroup {
					answerSum++
				}
			}
			answerSet = map[rune]int{}
			numPeopleInGroup = 0
		} else {
			log.Printf("Found another person in the group")
			numPeopleInGroup++
			// go through all the answers on this line
			// and record them in the answerSet
			for _, answerChar := range answerLine {
				log.Printf("This person answered yes for '%s'", string(answerChar))
				// here is where we increment the character in the set,
				// to indicate that one more person in the group answered
				// yes for it
				answerSet[answerChar]++
			}
		}
	}
	log.Printf("Sum of all the groups: %d", answerSum)
}
