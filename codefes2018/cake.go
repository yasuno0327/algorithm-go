package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text()) // 分割数N
	pieces := make([][]int, n)
	for i := range pieces { // 配列初期化
		pieces[i] = make([]int, n)
	}

	// 探索 1 => 食べられないピース 2 => スプレーかけたピース
	for i := range pieces {
		for j := range pieces[i] {
			if pieces[i][j] == 1 || pieces[i][j] == 2 { // 1 2だったら次の操作
				break
			}
			pieces[i][j] = 2 // スプレーをかける

			if i >= 1 { // 食べられないように色付け
				pieces[i-1][j] = 1
				pieces[i+1][j] = 1
			}
			if j >= 1 {
				pieces[i][j-1] = 1
				pieces[i][j+1] = 1
			}
			if i >= 1 && j >= 1 {
				pieces[i-1][j-1] = 1
				pieces[i-1][j] = 1
				pieces[i-1][j+1] = 1
				pieces[i][j-1] = 1
				pieces[i][j] = 1
				pieces[i][j+1] = 1
				pieces[i+1][j-1] = 1
				pieces[i+1][j] = 1
				pieces[i+1][j+1] = 1
			}
		}
	}

	// 出力
	for i := range pieces {
		for j := range pieces[i] {
			if pieces[i][j] == 1 {
				fmt.Print(".")
			}
		}
	}
}
