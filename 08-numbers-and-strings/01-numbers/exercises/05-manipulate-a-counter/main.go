// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Manipulate a Counter
//
//  1. Write the simplest line of code to increase
//     the counter variable by 1.
//
//  2. Write the simplest line of code to decrease
//     the counter variable by 1.
//
//  3. Write the simplest line of code to increase
//     the counter variable by 5.
//
//  4. Write the simplest line of code to multiply
//     the counter variable by 10.
//
//  5. Write the simplest line of code to divide
//     the counter variable by 5.
//
// EXPECTED OUTPUT
//  10
// ---------------------------------------------------------

import (
	"fmt"
)

func main() {
	var counter int;
	counter++;
	counter--;
	counter += 5;
	counter *= 10;
	counter /= 5;
	fmt.Println(counter);
}
