package main

import (
	"fmt"
	"strings"
)

func main() {
	var x string
	fmt.Print("Escreva uma string: ")
	fmt.Scan(&x)

	count := 1
	for _, c := range x {
		if strings.ToUpper(string(c)) == string(c) {
			count++
		}
	}
}
