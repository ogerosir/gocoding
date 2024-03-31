package main

import (
	"fmt"
)

var myChannel = make(chan string)

func main() {

	var a uint = 7
	var b uint = 8
	fmt.Println("Goroutine doing something in the background...")
	go getSum(a, b)
	fmt.Println("Main function executing")
	msg := <-myChannel
	fmt.Println(msg)
	// CPUCOUNT := runtime.NumCPU()
	// fmt.Println(CPUCOUNT)
}
func getSum(a uint, b uint) {
	sum := a + b
	var rtn = fmt.Sprintf("The sum of %v and %v is: %v\n", a, b, sum)
	myChannel <- rtn
}
