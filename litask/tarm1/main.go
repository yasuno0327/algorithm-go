package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	analysisArgs(os.Args[1:])
}

func analysisArgs(args []string) {
	length := len(args) / 2
	i := 1
	ni := 0
	mi := 0
	narray := make([]int, length)
	marray := make([]int, length)
	for _, v := range args {
		if i%2 == 0 {
			m, _ := strconv.Atoi(v)
			marray[mi] = m
			mi++
		} else if i%2 != 0 {
			n, _ := strconv.Atoi(v)
			narray[ni] = n
			ni++
		}
		i++
	}

	for i := 0; i < len(narray); i++ {
		wg.Add(1) // 関数が実行される前に終了するのを防ぐ
		go func(n int, m int) {
			devide(n, m)
		}(narray[i], marray[i])
	}
	wg.Wait() // ゴルーチンが実行完了するまで待つ。
}

func devide(n int, m int) int {
	k := 1
	for k <= n {
		if m*k%n == 0 {
			break
		}
		k++
	}
	fmt.Println(k)
	wg.Done() // ゴルーチン実行完了
	return k
}
