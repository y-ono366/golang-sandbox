package main

import(
	"fmt"
	"time"
)

func main() {
	pocky := time.Date(2018,11,11,0,0,0,0,time.Local)
	nowdate :=time.Now()

	var pockylist []string

	for nowdate.Unix() > pocky.Unix() {
		pockylist = append(pockylist,pocky.Format("2006/1/2"))
		pocky = pocky.AddDate(0, 0, 28)
	}
	pockylist = append(pockylist,pocky.Format("2006/1/2"))

	fmt.Println(pockylist)
}







