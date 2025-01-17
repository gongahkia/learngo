// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Declare string
//
//  1. Declare a string variable
//
//  2. Print that variable
//
// EXPECTED OUTPUT
//  ""
// ---------------------------------------------------------

import (
    "fmt"
)

func main() {

	var s = ""
	fmt.Printf("s (%T): %q\n", s, s)

    // some peripheral comments 
        // %T prints the type of the value
        // %q prints an empty string
}
