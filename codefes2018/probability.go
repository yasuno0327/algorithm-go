package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	text := sc.Text()
	n, _ := strconv.Atoi(text)
	fmt.Println(100 - (100 / n))
}
