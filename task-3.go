package main

import (
	"fmt"
)

func getMovies(num int) string {
	data := []int{1, 7, 3, 4, 8, 9}

	// movie := 0

	// for i := 0; i < len(data); i++ {
	// 	if num == data[i] {
	// 		movie = data[i]
	// 		// fmt.Println("1")
	// 	}
	// }

	// if movie > 0 {

	// 	t := strconv.Itoa(movie)
	// 	return t
	// }

	movies := 0

	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			movies = data[i] + data[j]
			if movies == num {
				num1 := data[i]
				num2 := data[j]
				output := fmt.Sprint(num1, " dan ", num2)
				return output
			}

		}
	}

	// t := strconv.Itoa(movies)
	return "no movies"
}
