package main

import (
	"fmt"
	"time"
)

const (
	SECONDS = 60
)

//display time now
func HelloNow() {
	var Time = time.Now()
	fmt.Println(Time.Format("25-06-2002"))
	fmt.Println("Time now :", Time)
}
func main() {
	HelloNow()
}
