// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Multi Assign #2
//
//  1. Assign the correct values to the variables
//     to match to the EXPECTED OUTPUT below
//
//  2. Print the variables
//
// HINT
//  Use multiple Println calls to print each sentence.
//
// EXPECTED OUTPUT
//  Air is good on Mars
//  It's true
//  It is 19.5 degrees
// ---------------------------------------------------------

import (
	"strconv"
	"fmt"
)

func main() {
	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	// ADD YOUR CODE BELOW

	planet = "Mars"
	isTrue = true
	temp = 19.5

	fmt.Println("Air is good on" + planet);
	fmt.Println("It's " + strconv.FormatBool(isTrue))
    fmt.Printf("It is %.1f degrees\n", temp)
}
