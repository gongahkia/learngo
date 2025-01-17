// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #2
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  10.5
// ---------------------------------------------------------

import (
	"fmt"
)

func main() {
	a, b := float64(10), 5.5
	a = b
	fmt.Println(a + b)
}
