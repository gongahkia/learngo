// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Count the Arguments
//
//  Print the count of the command-line arguments
//
// INPUT
//  bilbo balbo bungo
//
// EXPECTED OUTPUT
//  There are 3 names.
// ---------------------------------------------------------

func main() {
	// UNCOMMENT & FIX THIS CODE
	// count := len(strings.Split(os.Args, " ")) - 1; // there is no need to manually split the strings, the os.args[] handsles this simple tokenisation for you
	count := len(os.Args) - 1

	// UNCOMMENT IT & THEN DO NOT TOUCH THIS CODE
	fmt.Printf("There are %d names.\n", count)
}
