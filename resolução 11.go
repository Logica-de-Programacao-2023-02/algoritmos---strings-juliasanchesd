package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Escreva uma string: ")
	scanner.Scan()
	s1 := scanner.Text()

	r := strings.ReplaceAll(s1, "a", "")
	r = strings.ReplaceAll(r, "e", "")
	r = strings.ReplaceAll(r, "i", "")
	r = strings.ReplaceAll(r, "o", "")
	r = strings.ReplaceAll(r, "u", "")
	r = strings.ReplaceAll(r, "A", "")
	r = strings.ReplaceAll(r, "E", "")
	r = strings.ReplaceAll(r, "I", "")
	r = strings.ReplaceAll(r, "O", "")
	r = strings.ReplaceAll(r, "U", "")

	fmt.Print("Sua string sem vogais fica: ", r)
}
