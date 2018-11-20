package main

import (
	"fmt"
	"time"
)

func task1(result chan string) {
	// 重い処理を想定
	time.Sleep(time.Second * 2)
	fmt.Println("task1 finished!")
	
	result <- "task1 result"
}

func task2() {
	fmt.Println("task2 finished!")
}

// メイン関数
func main() {
	// chan型で受け渡すデータの型はstring
	result := make(chan string)
	
	// goを付けて並行処理にする
	go task1(result)
	go task2()
	
	// resultの中に何も入ってなければ、入ってくるまで待つ仕様になっている
	fmt.Println(<-result)
	
	// goroutineが終わる前に main 関数自体が終わる為、待ち時間をつける。
	// task2待ち
	time.Sleep(time.Second*1)
}
