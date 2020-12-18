package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	mapBytes, err := ioutil.ReadFile("./map.txt")
	if err != nil {
		log.Fatalf("Couldn't read map: %s", err)
	}
	mapStr := string(mapBytes)
	// []string ~= [][]char
	mapLines := strings.Split(mapStr, "\n")
	colNum := 1
	rowNum := 1
	numTrees := 0
	for {
		// go right 3, down 1 and see if we hit an octothorp
		// yes, an octothorp
		// that's what it's called
		// Jennifer says
		// and she's right
		// there's even a website
		// the end
		line := mapLines[rowNum-1]
		// if the last line is empty, we are at the end of the file
		// (my VSCode markdown plugin automatically appends a newline to all files)
		if line == "" {
			break
		}
		if (colNum + 3) > len(line) {
			remainingLength := len(line) - colNum
			colNum = 3 - remainingLength
		} else {
			colNum += 3
		}
		// now go down one. if the next row doesn't exist,
		// we're done
		if (rowNum + 1) >= len(mapLines) {
			break
		} else {
			rowNum++
		}
		if mapLines[rowNum-1][colNum-1] == '#' {
			numTrees++
		}
	}
	log.Printf("Num trees: %d", numTrees)
}
