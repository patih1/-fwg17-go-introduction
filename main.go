package main

import "fmt"

// func bite() {

// 	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	var res []byte
// 	for i := 0; i < len(charset); i++ {
// 		res = append(res, charset[i])
// 	}
// 	fmt.Print(res)
// }

func main() {
	printSegitiga(5)
	fmt.Println(getMoviesWithRange(7))
	fmt.Println(generatePassword("fazztrack", "strong"))
	// fmt.Println(getMovies(7))
	// bite()
}
