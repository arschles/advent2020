package main

import (
	"errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Zanuss: BODMAS and PEMDAS - mnemonics for order of operations

func main() {
	fileBytes, err := ioutil.ReadFile("../expenses.txt")
	if err != nil {
		log.Fatalf("Error reading file (%s)", err)
	}

	expensesStr := strings.Split(string(fileBytes), "\n")

	expenses := make([]int, len(expensesStr))
	for i := 0; i < len(expensesStr); i++ {
		expenseStr := expensesStr[i]
		expenseInt, err := strconv.Atoi(expenseStr)
		if err != nil {
			log.Printf("Expense %d is not an integer (%s)", i, err)
			continue
		}
		expenses[i] = expenseInt
	}
	if len(expenses) < 2 {
		log.Fatalf("Less than 2 expenses, no possible combination")
	}
	for i := 0; i < len(expenses); i++ {
		for j := i + 1; j < len(expenses); j++ {
			eltI := expenses[i]
			eltJ := expenses[j]
			sumIJ := eltI + eltJ
			rest := expenses[j+1:]

			maybeOther, err := find2020(sumIJ, rest)
			if err == errNotFound {
				// Jennifer: if rest is length 1, we don't have any more things to check. If we've
				// gotten to this point it means we haven't found a solution so we fail
				if len(rest) == 2 {
					log.Println("Hit the end of the list")
				}
				// otherwise, if we didn't find a solution here, move on and try the next one
				continue
			}

			log.Printf("Found solution: %d + %d  + %d = 2020", eltI, eltJ, maybeOther)
			log.Printf("%d * %d * %d = %d", eltI, eltJ, maybeOther, eltI*eltJ*maybeOther)
		}
	}

	// [1, 2, 3, 4]
	// find(1+2, [3, 4])
	// find(2+3, [4])

	// missing? (1+3, [2, 4])
	// (2+4, [1, 3])
	///
	// [1, 2, 3]

	// 1
	// 	2
	// 	3
	// 2
	// 	1
	// 	3
	// 3
	// 	1
	// 	2
}

var errNotFound = errors.New("Combination not found")

func find2020(cur int, rest []int) (int, error) {
	for _, expense := range rest {
		if cur+expense == 2020 {
			return expense, nil
		}
	}
	return 0, errNotFound
}
