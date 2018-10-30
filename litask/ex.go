package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	nString := os.Args[1]
	n, _ := strconv.ParseInt(nString, 10, 64)
	fmt.Println(strconv.FormatInt(n, 2))
	ex2(strconv.FormatInt(n, 2))
}

func ex2(n string) {
	length := len(n)
	result := big.NewInt(1)
	for _, v := range n {
		if length <= 0 {
			break
		}
		i := string(v)
		if i == "1" {
			exp := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(length-1)), nil)
			expRes := new(big.Int).Exp(big.NewInt(2), exp, nil)
			result.Mul(result, expRes)
		}
		length--
	}
	fmt.Println(result)
}
