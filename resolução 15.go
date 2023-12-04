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
	s2 := scanner.Text()
	fmt.Print("Fale outra string: ")
	scanner.Scan()
	s1 := scanner.Text()

	r1 := strings.ReplaceAll(s1, "a", "*")
	r2 := strings.ReplaceAll(s2, "a", "*")
	r1 = strings.ReplaceAll(r1, "e", "*")
	r2 = strings.ReplaceAll(r2, "e", "*")
	r1 = strings.ReplaceAll(r1, "i", "*")
	r2 = strings.ReplaceAll(r2, "i", "*")
	r1 = strings.ReplaceAll(r1, "o", "*")
	r2 = strings.ReplaceAll(r2, "o", "*")
	r1 = strings.ReplaceAll(r1, "u", "*")
	r2 = strings.ReplaceAll(r2, "u", "*")

	fmt.Printf("Suas strings ficaram assim: %s , %s \n  ", r1, r2)

}
