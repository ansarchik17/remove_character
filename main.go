package main

import "fmt"

func main() {
	fmt.Println(RemoveFirstAndLast("Ansar"))
	fmt.Println(RemoveFirstAndLast("An"))
}

func RemoveFirstAndLast(word string) string {
	if len(word) <= 2 {
		return ""
	}
	return word[1 : len(word)-1]
}
