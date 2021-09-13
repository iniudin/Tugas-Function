package main

import (
	"fmt"
	"unicode"
)

func countDigit(word string, length int) int {
	count := 0
	if length == 0 {
		return count
	} else {
		if unicode.IsDigit(rune(word[length-1])) {
			count++
			return count + countDigit(word, length-1)
		}
		return countDigit(word, length-1)
	}

}

func main() {
	var inputString string
	fmt.Println("Input Kata: ")
	fmt.Scan(&inputString)
	fmt.Printf("Terhitung %d digit pada kata\n", countDigit(inputString, len(inputString)))
}
