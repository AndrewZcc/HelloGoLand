package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("系统逻辑CPU数量：", runtime.NumCPU()) // 4, 双核四线程
}
