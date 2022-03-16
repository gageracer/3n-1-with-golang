package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	howMany := 10
	face := "heads"
	tossed := 0
	tries := 0
	side := ""
	showDetails := true
	coin := []string{
		"heads",
		"tails",
	}
	args := os.Args[1:]
	if len(args) == 1 {
		tmp, _ := strconv.Atoi(args[0])
		howMany = tmp
	}
	if len(args) == 2 {
		tmp, _ := strconv.Atoi(args[0])
		howMany = tmp
		if args[1] == "heads" || args[1] == "tails" {
			face = args[1]
		} else if args[1] == "h" {
			showDetails = false
		}
	}
	if len(args) == 3 {
		tmp, _ := strconv.Atoi(args[0])
		howMany = tmp
		if args[1] == "heads" || args[1] == "tails" {
			face = args[1]
		} else if args[1] == "h" {
			showDetails = false
		}
		if args[2] == "heads" || args[2] == "tails" {
			face = args[2]
		} else if args[2] == "h" {
			showDetails = false
		}
	}
Loop:
	for {
		rand.Seed(time.Now().UnixNano())
		// flip the coin
		side = coin[rand.Intn(len(coin))]
		tries++
		if showDetails {
			fmt.Println(tries, ":Flipped the coin and you get : ", side)
		}
		if side == face {
			tossed++
			if tossed > howMany/2 {
				fmt.Print("\033[H\033[2J")
				fmt.Print(tossed)
			}
			if tossed == howMany {
				break Loop
			}
		} else {
			tossed = 0
		}
	}
	// You get 30 tails after 428746222 tosses.
	fmt.Println("\nYou get", howMany, side, "after", tries, "tosses.")
}
