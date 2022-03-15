package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {

	startArg, endArg := big.NewInt(1), big.NewInt(5)
	showResults := true
	matrix := false
	args := os.Args[1:]
	//fmt.Println(len(args))

	if len(args) > 0 {
		if args[0] == "no" || args[0] == "h" || args[0] == "n" {
			showResults = false
		} else if args[0] == "m" {
			matrix = true
			showResults = false
		} else {
			tmp, _ := strconv.ParseInt(args[0], 10, 64)
			startArg = big.NewInt(tmp)
			startArg = min(startArg, big.NewInt(0))
			endArg.Mul(startArg, big.NewInt(2))
		}
	}
	if len(args) > 1 {
		if args[1] == "no" || args[1] == "h" || args[1] == "n" {
			showResults = false
		} else if args[1] == "m" {
			matrix = true
			showResults = false
		} else {
			tmp, _ := strconv.ParseInt(args[1], 10, 64)
			endArg = big.NewInt(tmp)
			endArg = min(endArg, startArg)
		}
	}
	if len(args) > 2 {
		if args[2] == "no" || args[2] == "h" || args[2] == "n" {
			showResults = false
		} else if args[2] == "m" {
			matrix = true
			showResults = false
		}
		// fmt.Println(showResults)
	}

	for start := startArg; start.Cmp(endArg) < 0; start.Add(start, big.NewInt(1)) {
		round := 1
		num := big.NewInt(0)
		a := big.NewInt(9)
		for num.Add(start, num); num.Cmp(big.NewInt(1)) != 0; round++ {
			if a.Mod(num, big.NewInt(2)); a.Cmp(big.NewInt(0)) == 0 {
				num.Div(num, big.NewInt(2))
			} else {
				num.Mul(num, big.NewInt(3))
				num.Add(num, big.NewInt(1))
			}
			if matrix {
				fmt.Print(round, ":", num, "|")
			}
		}
		if showResults {
			fmt.Printf("%v took %v rounds to reach the loop.\n",
				start, round)
			fmt.Println("----------------------------------------")
		}
	}
	fmt.Printf("\nLast number is %v.\n", startArg)
}

func min(a, b *big.Int) *big.Int {
	if a.Cmp(b) == 1 {
		return a
	}
	result := big.NewInt(0)
	return result.Add(b, big.NewInt(1))
}
