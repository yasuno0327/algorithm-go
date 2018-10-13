package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	if len(text) > 3 {
		return
	}
	result := strings.Count(text, "1")
	fmt.Println(result)
}
