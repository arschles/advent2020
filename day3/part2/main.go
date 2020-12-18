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
	slopes := []struct {
		right, down int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	totalProduct := 1
	for _, slope := range slopes {
		numTrees := traverse(mapLines, slope.right, slope.down)
		totalProduct *= numTrees
	}
	log.Printf("Tree Multiplication: %d", totalProduct)
}

func traverse(lines []string, right, down int) int {
	colNum := 1
	rowNum := 1
	numTrees := 0
	for {
		// go right and down and see if we hit an octothorp
		// yes, an octothorp
		// that's what it's called
		// Jennifer says
		// and she's right
		// there's even a website
		// the end
		line := lines[rowNum-1]
		// if the last line is empty, we are at the end of the file
		// (my VSCode markdown plugin automatically appends a newline to all files)
		if line == "" {
			break
		}
		if (colNum + right) > len(line) {
			remainingLength := len(line) - colNum
			colNum = right - remainingLength
		} else {
			colNum += right
		}
		// now go down. if that row doesn't exist,
		// we're done
		if (rowNum + down) >= len(lines) {
			break
		} else {
			rowNum += down
		}
		if lines[rowNum-1][colNum-1] == '#' {
			numTrees++
		}
	}
	return numTrees
}
