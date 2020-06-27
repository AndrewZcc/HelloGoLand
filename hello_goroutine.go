package main

import "fmt"

/** 协程 goroutine 学习
 * 主goroutine 和 子 goroutine
 */
func main() {
	fmt.Println("Hello Go!")

	// 启动一个子Goroutine
	go printNum()

	for i:=1; i<=100; i++ {
		fmt.Printf("\t主Goroutine中打印字母：A, %d\n", i)
	}

	fmt.Println("Main Goroutine Over! ...")
}

func printNum() {
	for i:=1; i<=100; i++ {
		fmt.Printf("子Goroutine打印数字: %d\n", i)
	}

	fmt.Println("Sub Goroutine PrintNum over ...")
}