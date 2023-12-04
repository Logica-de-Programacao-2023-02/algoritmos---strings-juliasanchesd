package main

import (
	"fmt"
	"strings"
)

func sa(s1, s2 string) bool {
	s1 = strings.ToLower(strings.ReplaceAll(s1, " ", ""))
	s2 = strings.ToLower(strings.ReplaceAll(s2, " ", ""))

	if len(s1) != len(s2) {
		return false
	}

	countMap := make(map[rune]int)
	for _, char := range s1 {
		countMap[char]++
	}

	for _, char := range s2 {
		if countMap[char] <= 0 {
			return false
		}
		countMap[char]--
	}

	return true
}

func main() {
	var s1, s2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&s1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&s2)

	if sa(s1, s2) {
		fmt.Println("As strings são anagramas.")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}
