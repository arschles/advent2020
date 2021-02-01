package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	seatSpecLines, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Error reading file (%s)", err)
	}

	seatIDs := map[int]struct{}{}
	for _, seatSpecLine := range strings.Split(string(seatSpecLines), "\n") {
		if seatSpecLine == "" {
			continue
		}

		firstSeven := seatSpecLine[0:7]
		startRow := 0
		endRow := 127
		for _, char := range firstSeven {
			startRow, endRow = getRangeOfRows(startRow, endRow, char)
		}

		// at this point, startRow and endRow should be the same
		row := startRow
		lastThree := seatSpecLine[7:]
		startCol := 0
		endCol := 7
		for _, char := range lastThree {
			startCol, endCol = getRangeOfCols(startCol, endCol, char)
		}
		// similar to rows, startCol is gonna be the same as endCol here
		col := startCol

		seatID := (row * 8) + col
		seatIDs[seatID] = struct{}{}
	}

	for seatID := range seatIDs {
		if _, ok := seatIDs[seatID+1]; !ok {
			log.Printf("Found your seat: %d", seatID+1)
		} else if _, ok := seatIDs[seatID-1]; !ok {
			log.Printf("Found your seat: %d", seatID-1)
		}
	}
	// NOTE: Jennifer is smart. You'll get 4 printouts from the above.
	// Take the one that is most "in the middle" and recognize that the low or high
	// seat IDs mean that you're near the front or back of the plane, and
	// hence you don't have a seat on this plane because it's smaller. have fun
	// sitting on the wing or just plain flying

}

//
func getRangeOfRows(startRow, endRow int, letter rune) (int, int) {
	startOffset := (endRow - startRow) / 2
	if letter == 'F' {
		return startRow, (startRow + startOffset)
	}
	return (startRow + 1 + startOffset), endRow
}

func getRangeOfCols(startCol, endCol int, letter rune) (int, int) {
	startOffset := (endCol - startCol) / 2
	if letter == 'L' {
		return startCol, (startCol + startOffset)
	}
	return (startCol + 1 + startOffset), endCol
}
