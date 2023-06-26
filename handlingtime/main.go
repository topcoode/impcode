package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now()
	fmt.Println(presentTime)
	//2022-11-21 15:26:31.2616293 +0530 IST m=+0.004698901
	fmt.Print(presentTime.Format("09/12/2002"))
	createdaDate := time.Date(2020, time.December, 10, 02, 25, 45, 555, time.UTC)
	fmt.Println(createdaDate)
	//2022-11-21 15:36:38.2419557 +0530 IST m=+0.006795301
	//09/1121/213252020-12-10 02:25:45.000000555 +0000 UT
}
