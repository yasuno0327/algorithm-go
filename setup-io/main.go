package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// light read
func lightRead(a *int) {
	fmt.Println("数値を入力してください。")
	fmt.Scan(a)
	fmt.Println(*a)
}

// Read new line delimited
func delimitedLine() {
	var sc = bufio.NewScanner(os.Stdin) // func NewScanner(r io.Reader) *Scanner
	var s, t string
	fmt.Println("1行目の文字列を入力してください。")
	if sc.Scan() { // func (s *Scanner) Scan() bool: true => success, false => fail
		s = sc.Text()
	}
	fmt.Println("2行目の文字列を入力してください。")
	if sc.Scan() {
		t = sc.Text()
	}
	fmt.Println("入力した文字列")
	fmt.Println("1行目" + s)
	fmt.Println("2行目" + t)
}

// Read 1 line and return
func nextLine() string {
	var sc = bufio.NewScanner(os.Stdin)
	fmt.Println("文字列を入力してください。")
	sc.Scan()
	return sc.Text()
}

// Read string with delimited space
func delimitedSpace() []int {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		arr[i], _ = strconv.Atoi(sc.Text())
	}
	return arr
}

func nextInt() {

}

func main() {
	// var a int
	// lightRead(&a)
	// delimitedLine()
	// st := nextLine()
	// fmt.Println(st)
	arr := delimitedSpace()
	for v := range arr {
		fmt.Println(v)
	}
}
