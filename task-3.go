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

	return "no movies"
}

func getMoviesWithRange(num int) string {
	data := []int{1, 7, 3, 4, 8, 9}

	movies := 0

	for _, item1 := range data { // 2, 3
		for _, item2 := range data { // 3, 4
			movies = item1 + item2 // 10

			if movies == num {
				output := fmt.Sprintf("%v dan %v", item1, item2)
				return output
			}
		}
	}

	return "no movies"
}
