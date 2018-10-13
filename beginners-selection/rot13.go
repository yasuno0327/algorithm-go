package main

import (
	"fmt"
	"os"
)

func main() {
	run(os.Args[1:])
}

func run(args []string) {
	for i := range args {
		changeRot13(args[i])
	}
}

func changeRot13(arg string) {
	s := []byte(arg)
	for i := range s {
		s[i] = rot13(s[i])
	}
	fmt.Printf("%s\n", s)
}

func rot13(char byte) byte {
	switch {
	case (char >= 'A' && char <= 'Z'):
		return (char-'A'+13)%26 + 'A'
	case (char >= 'a' && char <= 'z'):
		return (char-'z'+13)%26 + 'a'
	default:
		return char
	}
}
