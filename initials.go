// Initials project initials.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func scan() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}

func setInitials(s string) {

	start := 0

	for s[start] == ' ' {
		start++
	}

	fmt.Printf("%c", unicode.ToUpper(rune(s[start])))

	length := len(s)

	for i := start + 1; i < length; i++ {
		for s[i] == ' ' {
			i++
			if i < length && s[i] != ' ' {
				fmt.Printf("%c", unicode.ToUpper(rune(s[i])))
			}
		}
	}

	fmt.Printf("\n")

}

func main() {
	setInitials(scan())
}
