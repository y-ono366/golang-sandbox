package main

import(
	"fmt"
	"time"
)

func main() {
	pocky := time.Date(2018,10,15,0,0,0,0,time.Local)
	nowdate :=time.Now()

	var pockylist []string
	for nowdate.Unix() > pocky.Unix() {
		pockylist = append(pockylist,pocky.Format("2006/1/2"))
		pocky = pocky.AddDate(0, 0, 28-1)
	}
	pockylist = append(pockylist,pocky.Format("2006/1/2"))

	fmt.Printf("before  "+ pockylist[len(pockylist)-2] + " next  "+pockylist[len(pockylist)-1])
}
