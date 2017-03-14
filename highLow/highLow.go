// Samples used in a small go tutorial
//
// Copyright (C) 2017 framp at linux-tips-and-tricks dot de
//
// Samples for go syntax - simple high/low to demonstrate some go syntax constructs

package main

import ( // used packages
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// constant definitions
const ( // upper and lower bounds
	highLimit = 100 // integer constant
	lowLimit  = 1
)

// main program
func main() {

	// variable definitions
	var (
		high        = highLimit // type int derived from constant
		low         = lowLimit
		actualValue int                           // ints are initialized with 0
		err         error                         // pointers are initialized with nil
		guessValue  = rand.Intn(high-low+1) + low //
		guesses     int                           // ints are initialized with 0
	)

	for { // endless loop
		reader := bufio.NewReader(os.Stdin)                                                // read from console
		fmt.Printf("Guesses: %d - Low: %d - High: %d -> Your guess: ", guesses, low, high) // prompt for input
		text, _ := reader.ReadString('\n')                                                 // returns two values, _ is an unnamed variable thus the error returned is ignored

		if actualValue, err = strconv.Atoi(strings.TrimSpace(text)); err != nil { // if statement, atoi returns two results - second result is used to notify about errors
			fmt.Printf("Invalid number %s entered\n", text)
			continue // continue loop
		}

		if actualValue < low || actualValue > high { // logical comparisons
			fmt.Printf("Number %d is out of bounds\n", actualValue)
			continue // continue loop
		}

		guesses++ // increment operator

		if actualValue == guessValue {
			fmt.Printf("Congratulations: Number %d guessed with %d guesses\n", guessValue, guesses)
			break
		}

		switch actualValue > guessValue { // switch
		case true:
			high = actualValue // no break needed because this is standard
		case false:
			low = actualValue
		}
	}
}
