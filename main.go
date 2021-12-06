package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type expense struct {
	name string
	amt float64
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage: sumsum <file> <sum>")
	}

	sum, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatalln(err)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	var expenses []expense
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		expenseStr := strings.Split(s, ",")
		if len(expenseStr) != 2 {
			log.Fatalf("invalid line: %s\n", s)
		}

		amt, err := strconv.ParseFloat(strings.TrimSpace(expenseStr[1]), 64)
		if err != nil {
			log.Fatalln(err.Error())
		}

		expenses = append(expenses, expense{
			name: strings.TrimSpace(expenseStr[0]),
			amt:  amt,
		})
	}

	err = printSumComponents(sum, expenses)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func printSumComponents(sum float64, expenses []expense) error {
	expenseMap := map[float64][]expense{}

	fmt.Printf("Pair(s) found:\n")

	for _, exp := range expenses {
		_, ok := expenseMap[exp.amt]
		if ok {
			expenseMap[exp.amt] = append(expenseMap[exp.amt], exp)
		} else {
			expenseList := []expense{exp}
			expenseMap[exp.amt] = expenseList
		}
	}

	amtMap := map[float64]bool{}

	for amt, exp := range expenseMap {
		amtMap[amt] = true
		if amt == sum/2.0 {
			exps := expenseMap[amt]
			if len(exps) > 1 {
				for i := 0; i < len(exps)-1; i++ {
					for j := i+1; j < len(exps); j++ {
						fmt.Printf("[(%s:%0.2f), (%s:%0.2f)]\n", exps[i].name, exps[i].amt, exps[j].name, exps[j].amt)
					}
				}
			}
		} else if _, ok := amtMap[sum-amt]; ok {
			exp2 := expenseMap[sum-amt]
			for _, e1 := range exp {
				for _, e2 := range exp2 {
					fmt.Printf("[(%s:%0.2f), (%s:%0.2f)]\n", e1.name, e1.amt, e2.name, e2.amt)
				}
			}
			fmt.Printf("\n")
		}
	}

	return nil
}
