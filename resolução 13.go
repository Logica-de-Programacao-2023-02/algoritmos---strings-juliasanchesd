package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Print("Digite uma string numérica separada por espaços: ")
	fmt.Scanln(&s)

	var ca rune

	c := false

	for _, char := range s {
		if char == ' ' || char == '\t' {
			continue
		}

		numero := int(char - '0')

		if numero > int(ca) {
			c = true
			break
		}

		ca = char
	}
	if c {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}
