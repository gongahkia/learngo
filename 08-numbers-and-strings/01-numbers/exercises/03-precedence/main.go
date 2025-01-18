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
// EXERCISE: Precedence
//
//  Change the expressions to produce the expected outputs
//
// RESTRICTION
//  Use parentheses to change the precedence
// ---------------------------------------------------------

func main() {
	fmt.Println(10 + 5 - (5 - 10));
	fmt.Println(-10 + 0.5 - (1 + 5.5));
	fmt.Println(5 + 10*(2-5));
	fmt.Println(0.5 * (2 - 1));
	fmt.Println((3+1)/2*10 + 4);
	fmt.Println(10 / 2 * (10 % 7));
	fmt.Println(100 / (5.0 / 2));
}
