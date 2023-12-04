package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Fale uma string: ")
	scanner.Scan()
	s1 := scanner.Text()

	s1 = strings.ReplaceAll(s1, " ", "")
	sN := true

	for _, caractere := range s1 {
		if caractere < '0' || caractere > '9' {
			sN = false
			break
		}
	}

	if sN {
		fmt.Println("A string contém apenas números (0 a 9).")
	} else {
		fmt.Println("A string não contém apenas números (0 a 9).")
	}
}
