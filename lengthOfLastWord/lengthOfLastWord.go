package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("a"))
	return
}

func lengthOfLastWord(s string) int {
	splitedWord := strings.Split(s, " ")
	for i := len(splitedWord) - 1; i >= 0; i-- {
		if len(splitedWord[i]) > 0 {
			return len(splitedWord[i])
		}
	}

	return 0
}
