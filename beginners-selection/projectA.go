package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	i := ReadNewInt()
	j, k := ReadNewIntDelimitedSpace()
	total := i + j + k
	s := ReadNewLine()
	fmt.Println(strconv.Itoa(total) + " " + s)
}

func ReadNewInt() int {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	i, err := strconv.Atoi(text)
	if err != nil {
		panic(err)
	}
	return i
}

func ReadNewIntDelimitedSpace() (int, int) {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	text1 := sc.Text()
	sc.Scan()
	text2 := sc.Text()
	j, _ := strconv.Atoi(text1)
	k, _ := strconv.Atoi(text2)
	return j, k
}

func ReadNewLine() string {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}
