package main

import (
	"fmt"
	"time"
)

func comparing_time() {
	currentTime := time.Now()

	mytime := currentTime.Add(time.Minute * 3)

	firstTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)

	fmt.Println("firstTime", firstTime)

	secondTime := time.Date(2021, 12, 25, 16, 40, 55, 200, time.UTC)

	fmt.Println("secondTime", secondTime)

	fmt.Println("current_time : ", currentTime.Format("2006.01.02 15:04:05"))

	fmt.Println("my_time : ", mytime.Format("2006.01.02 15:04:05"))

	g1 := currentTime.Before(mytime)

	g2 := currentTime.After(mytime)

	fmt.Println("g1", g1)

	fmt.Println("g2", g2)

}
