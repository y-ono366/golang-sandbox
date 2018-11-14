package main

import(
	"fmt"
	"time"
)

func main() {
	const layout = "2006-01-02"

	pocky := time.Date(2018,10,15,0,0,0,0,time.Local)
	nowdate :=time.Now()

	var pockylist []int64
	for nowdate.Unix() > pocky.Unix() {
		pockylist = append(pockylist,pocky.Unix())
		pocky = pocky.AddDate(0, 0, 28-1)
	}
	pockylist = append(pockylist,pocky.Unix())

	beforeTime := time.Unix(pockylist[len(pockylist)-2],0)
	beforeEndTime := beforeTime.AddDate(0,0,4-1)
	nextTime := time.Unix(pockylist[len(pockylist)-1],0)
	nextEndTime := nextTime.AddDate(0,0,4-1)

	fmt.Printf("before :"+beforeTime.Format(layout)+ "   last  :" + beforeEndTime.Format(layout)+"\n")
	fmt.Printf("next   :"+nextTime.Format(layout)+   "   last  :" + nextEndTime.Format(layout))
}
