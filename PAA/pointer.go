package main

import (
	"fmt"
)

func main() {
	// int型のポインタ変数
	var pointer *int
	// int型変数
	var n int = 100

	// &（アドレス演算子）を使って、nのアドレスを代入
	pointer = &n

	fmt.Println("nのアドレス：", &n)
	fmt.Println("pointerの値：", pointer)

	fmt.Println("nの値：", n)
	// *(間接参照演算子）を利用して、ポインタの中身を取得
	fmt.Println("pointerの中身：", *pointer)
}
