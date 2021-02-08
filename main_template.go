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
