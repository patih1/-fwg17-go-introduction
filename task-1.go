package main

import "fmt"

func printSegitiga(num int) {
	for i := 0; i < num; i++ {

		for k := 0; k < i; k++ {
			fmt.Printf(" ")
		}

		for j := i + 1; j < num; j++ {
			fmt.Printf("**")
		}

		fmt.Println("*")
	}
}
