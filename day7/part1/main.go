package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	answerLines, err := ioutil.ReadFile("./test_input.txt")
	if err != nil {
		log.Fatalf("Error reading file (%s)", err)
	}

	for _, answerLine := range strings.Split(string(answerLines), "\n") {

	}

}

func parseLine(line string) AThingWhoKnows {
	// we need to parse line from this format:
	//
	//	$ADJECTIVE $COLOR BAGS CONTAIN $NUMBER $ADJECTIVE $COLOR BAGS [, $ADJECTIVE $COLOR BAG[S]]...
	//
	// into a data structure that can traverse from the initial $ADJECTIVE $COLOR
	// into one or more of the bags that can contain that bag
	//
	// in other words, we want to map the _contained_ bag to all of its potential
	// _container_ bags. this will look like a graph
	//
	// we also want to detect any loops in this graph
}
