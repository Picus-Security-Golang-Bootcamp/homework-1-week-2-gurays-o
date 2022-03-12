package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	var moviesList = [10]string{
		"The Shawshank Redemption",
		"The Godfather",
		"The Dark Knight",
		"The Godfather: Part II",
		"12 Angry Men",
		"Schindler's List",
		"The Lord of the Rings: The Return of the King",
		"Pulp Fiction",
		"The Lord of the Rings: The Fellowship of the Ring",
		"The Good, the Bad and the Ugly",
	}

	userInput := os.Args[1:]

	userCommand := userInput[0]

	switch userCommand {
	case "list":
		for _, movie := range moviesList {
			fmt.Println(movie)
		}
	case "search":
		searchQuery := strings.TrimSpace(strings.Join(userInput[1:], " "))
		isFound := false
		for i, movie := range moviesList {
			if strings.ToLower(searchQuery) == strings.ToLower(movie) {
				fmt.Println("Title: " + movie)
				fmt.Printf("Position in the list: %d\n", i+1)
				isFound = true
				break
			}
		}
		if isFound == false {
			fmt.Println("Movie not found.")
		}
	default:
		fmt.Println("You did not provide a valid input.")
	}
}
