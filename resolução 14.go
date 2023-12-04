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

	for _, caractere := range s {
		if caractere == ' ' || caractere == '\t' {
			continue
		}

		numero := int(caractere - '0')

		if numero < int(ca) {
			c = true
			break
		}

		ca = caractere
	}
	if c {
		fmt.Println("A string é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica decrescente.")
	}
}
