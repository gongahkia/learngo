// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Multiple
//
//  1. Declare two variables using
//     multiple variable declaration statement
//
//  2. The first variable's name should be active
//  3. The second variable's name should be delta
//
//  4. Print them all
//
// HINT
//  You should declare a bool and an int variable
//
// EXPECTED OUTPUT
//  false 0
// ---------------------------------------------------------

func main() {
	var (
	    active bool
	    delta int
	);
    active = false;
    delta = 0;
    // note that value assignment here is technically unneccesary
	fmt.Println(active, delta);
}
