// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Refactor Feet to Meter
//
//  Remember the feet to meters program?
//  Now, it's time to refactor it.
//  Define your own Feet and Meters types.
//
//  Follow the steps inside the code.
// ---------------------------------------------------------

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	type (
		Feet   float64
		Meters float64
	);
	var (
		feet   Feet
		meters Meters
	);
	arg := os.Args[1];
	val, _ := strconv.ParseFloat(arg, 64);
	feet = Feet(val);
	meters = Meters(feet * 0.3048);
	fmt.Printf("%g feet is %g meters.\n", feet, meters);
}
