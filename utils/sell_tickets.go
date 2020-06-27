package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket = 10
var wg sync.WaitGroup

func main(){
	fmt.Println("main goroutine start")
	wg.Add(4)

	go saleTickets("售票口 1")
	go saleTickets("售票口 2")
	go saleTickets("售票口 3")
	go saleTickets("售票口 4")

	/*
	Q: 主线程如何等待子线程全部执行结束？
	A: 1. 主线程进入等待状态 time.Sleep
	   2. 使用 同步等待组 sync WaitGroup
	   3. 使用 通道 channel
	*/
	//time.Sleep(5*time.Second)
	wg.Wait()
	fmt.Println("main ... over ...")
}

func saleTickets(name string) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())

	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name, " 售出：", ticket)
			ticket--
		} else {
			fmt.Println(name, "售罄，票卖完了...")
			break
		}
	}
}