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
		cur := expenses[i]
		// Jennifer: we only need to test out the numbers _after_ cur because we already tested
		// combinations that include numbers before cur
		rest := expenses[i+1:]
		maybeOther, err := find2020(cur, rest)
		if err == errNotFound {
			// Jennifer: if rest is length 1, we don't have any more things to check. If we've
			// gotten to this point it means we haven't found a solution so we fail
			if len(rest) == 1 {
				log.Println("Hit the end of the list")
			}
			// otherwise, if we didn't find a solution here, move on and try the next one
			continue
		}
		log.Printf("Found solution: %d + %d = 2020", cur, maybeOther)
		log.Printf("%d * %d = %d", cur, maybeOther, cur*maybeOther)
	}

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
