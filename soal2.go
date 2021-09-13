package main

import (
	"fmt"
	"unicode"
)

func countLower(word string, length int) int {
	count := 0
	if length == 0 {
		return count
	} else {
		if unicode.IsLower(rune(word[length-1])) {
			count++
			return count + countLower(word, length-1)
		}
		return countLower(word, length-1)
	}

}

func main() {
	var inputString string
	fmt.Println("Input Kata: ")
	fmt.Scan(&inputString)
	fmt.Printf("Terhitung %d huruf kecil\n", countLower(inputString, len(inputString)))
}
