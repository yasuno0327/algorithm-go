package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	// スペース区切りでa,bを入力
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	x, _ := strconv.Atoi(sc.Text())
	humans := make([][]int, n)
	for i := 0; i < n; i++ {
		humans[i] = make([]int, 2)
		for j := range humans[i] {
			sc.Scan()
			humans[i][j], _ = strconv.Atoi(sc.Text())
		}
	}

	// biの最大値を求める
	max := 0
	for i := 0; i < n; i++ {
		if humans[max][1] < humans[i][1] {
			max = i
		}
	}
	humans[max][0] = humans[max][0] + x // x分足す

	// 好感度の合計値を計算
	result := 0
	for i := range humans {
		total := 1
		for j := range humans[i] {
			total = total * humans[i][j]
		}
		result = result + total
	}
	fmt.Println(result)
}
