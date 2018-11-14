package main

import(
	"fmt"
	"time"
)

const layout = "2006-01-02"

func main() {
	pocky := time.Date(2018,10,15,0,0,0,0,time.Local)
	nowdate :=time.Now()

	var pockylist []int64
	for nowdate.Unix() > pocky.Unix() {
		pockylist = append(pockylist,pocky.Unix())
		pocky = pocky.AddDate(0, 0, 28-1)
	}
	pockylist = append(pockylist,pocky.Unix())
	fmt.Printf(printOut(pockylist))
}

func printOut(list[]int64) string{
	beforeTime := time.Unix(list[len(list)-2],0)
	beforeEndTime := beforeTime.AddDate(0,0,4-1)

	nextTime := time.Unix(list[len(list)-1],0)
	nextEndTime := nextTime.AddDate(0,0,4-1)
	
	str := "before :"+beforeTime.Format(layout)+ "   last  :" + beforeEndTime.Format(layout)+"\n"
	str += "next   :"+nextTime.Format(layout)+   "   last  :" + nextEndTime.Format(layout)+"\n"
	return str
}
