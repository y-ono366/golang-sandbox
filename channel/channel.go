package main

import (
	"fmt"
)

func listenChannel(myCh chan int) {
	// myChからの送信を受け取る
	for out := range myCh {
		fmt.Println(fmt.Sprintf("out from channel => %d",out))
	}
}
func main() {
	fmt.Println("This is func main")

	// intを送受信どちらもできるchannel
	myChan := make(chan int)

	// 並行処理goroutineを立ち上げる
	go listenChannel(myChan)

	myChan <- 100 // 100を流し込む
	fmt.Printf("main exited now")
}

//Channel は goroutine 間でのメッセージパッシングをするためのもの
//メッセージの型を指定できる
//first class value であり、引数や戻り値にも使える
//send/receive でブロックする
//buffer で、一度に扱えるメッセージ量を指定できる
