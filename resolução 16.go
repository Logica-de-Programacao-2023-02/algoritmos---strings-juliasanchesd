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
	fmt.Print("Fale outra string: ")
	scanner.Scan()
	s2 := scanner.Text()

	if strings.Contains(s1, s2) {
		fmt.Printf("A string %s é uma substring da string %s.", s2, s1)
	} else {
		fmt.Printf("A string %s não é uma substring da string %s.", s2, s1)
	}
}
