package main

import(
	"fmt"
)

func main(){
	var arr = []string{"A","S","A","P"}
	// for key,value := range arr {
	// 	fmt.Println(key)
	// 	fmt.Println(value)
	// }
	var sum int
	for i:=0; i< 1000000;i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(arr)
}
