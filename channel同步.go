package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func output(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println("  ---  ")
}
func person1() {
	output("hello")
	ch <- 66
}
func person2() {
	<-ch
	output("nihao,woshi xxx")

}
func main() {

	go person1()
	go person2()
	for {

	}
}
