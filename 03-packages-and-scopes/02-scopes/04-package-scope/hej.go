package main

import "fmt";

func hej(){

	fmt.Println("hej");
	fmt.Println("since this is also in the main package, it can also be called from the main function");
	fmt.Println("and if it isn't redeclared in the same package scope, then everything's alright and it can run error-free");

}
