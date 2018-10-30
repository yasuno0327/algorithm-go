package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	n := new(big.Int)
	nString := os.Args[1]
	ex2(nString)
	n, _ = n.SetString(os.Args[1], 10)
	mString := os.Args[2]
	m, _ := strconv.ParseInt(mString, 10, 64)

	nRes := exponential(n)
	mRes := factorial(m)

	printLargest(nRes, mRes)
}

func ex2(n string) {
	length := len(n)
	for _, v := range n {
		i := string(v)
		length--
	}
}

func printLargest(n *big.Int, m *big.Int) {
	cmpRes := n.Cmp(m)
	switch cmpRes {
	case -1:
		fmt.Println("Factorial")
	case 1:
		fmt.Println("Exponential")
	}
}

func exponential(n *big.Int) *big.Int {
	return new(big.Int).Exp(big.NewInt(2), n, nil)
}

func factorial(m int64) *big.Int {
	return new(big.Int).MulRange(1, m)
}
