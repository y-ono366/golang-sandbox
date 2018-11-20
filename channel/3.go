package main

import (
	"fmt"
)

func listenChannel(myCh chan int) {
	fmt.Println("listenChannel")
	// myChからの送信を受け取る
	for out := range myCh {
		fmt.Println(fmt.Sprintf("out from channel => %d",out))
	}
	fmt.Println("end listenChannel")
}
func main() {
	fmt.Println("This is func main")

	// intを送受信どちらもできるchannel
	myChan := make(chan int)
	fmt.Println(myChan)

	// 並行処理goroutineを立ち上げる
	go listenChannel(myChan)
	myChan <- 100
	fmt.Println("main exited now")
}
