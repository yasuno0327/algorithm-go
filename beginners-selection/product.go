package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	arr := make([]int, 2)
	for i := 0; i < 2; i++ {
		sc.Scan()
		v := sc.Text()
		arr[i], _ = strconv.Atoi(v)
	}

	result := 1
	for _, v := range arr {
		result *= v
	}

	switch result % 2 {
	case 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}
}
