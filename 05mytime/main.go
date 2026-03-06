package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("My time ")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}
