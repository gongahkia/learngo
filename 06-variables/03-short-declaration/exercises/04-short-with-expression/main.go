// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Short With Expression
//
// 	1. Short declare a variable named `sum`
//
//  2. Initialize it with an expression by adding 27 and 3.5
//
// EXPECTED OUTPUT
//  30.5
// ---------------------------------------------------------

import (
	"fmt"
)

func main() {
	// ADD YOUR DECLARATION HERE

    sum := 27 + 3.5;

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(sum)

	// due to float and integer addition, go will automatically recognise sum as a float datetype
}
