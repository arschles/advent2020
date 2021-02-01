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

	maxSeatID := 0
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
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	log.Printf("MAX SEAT ID = %d", maxSeatID)

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
