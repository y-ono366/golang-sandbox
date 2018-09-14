package main

import(
    "fmt"
    "sync"
    "time"
)
func main(){
    var s, e time.Time
    s = time.Now()
    var wg sync.WaitGroup
    for i:=0; i<20; i++{
        wg.Add(1)
        go func(num int){
            defer wg.Done()
            process(num)
        }(i)
    }
    wg.Wait()
    e = time.Now()
    fmt.Printf("処理完了 : %v Seconds\n", (e.Sub(s)).Seconds())
}
func process(num int){
    fmt.Printf("%d 番目の処理開始\n", num)
    time.Sleep(1 * time.Second)
    fmt.Printf("%d 番目の処理終了\n", num)
}
