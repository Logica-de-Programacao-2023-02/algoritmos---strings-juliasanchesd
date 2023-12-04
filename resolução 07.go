package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Me fale uma string: ")
	scanner.Scan()
	f1 := scanner.Text()

	tn := false
	for _, char := range f1 {
		if unicode.IsDigit(char) {
			tn = true
			break
		}
	}

	if tn {
		fmt.Printf("Seu string %s têm número.", f1)
	} else {
		fmt.Printf("Sua string %s não contem números.", f1)
	}

}
