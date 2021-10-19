package main

import (
	"fmt"
)

func main() {
	title := "Golang is the best"

	fmt.Println("Kasus Index Genap")

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("Letter : ", string(letter))
		} else {
			fmt.Println("Index tidak genap")
		}
	}

	fmt.Println("Kasus Huruf Vokal")

	for _, letter := range title {
		if string(letter) == "a" {
			fmt.Println("Letter : ", string(letter))
		} else if string(letter) == "i" {
			fmt.Println("Letter : ", string(letter))
		} else if string(letter) == "u" {
			fmt.Println("Letter : ", string(letter))
		} else if string(letter) == "e" {
			fmt.Println("Letter : ", string(letter))
		} else if string(letter) == "o" {
			fmt.Println("Letter : ", string(letter))
		} else {
			fmt.Println("Letter bukan huruf vokal")
		}
	}
}
