package main

import (
	"fmt"
)

// calculate and return values
func innerCal(number int) int{
	f:=fibonacci()
	if number == 0 || number == 1{
		return number
	}
	return f(number-1)+f(number-2)
}

// func fibonacci to call inner calculator
func fibonacci () func (number int) int {
	return innerCal
}

// main func
func main (){
	f:=fibonacci()
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(f(i))
	}
}