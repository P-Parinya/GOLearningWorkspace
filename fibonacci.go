package main

import (
	"fmt"
)

// calculate and return values
func innerCal(i int) int{
	f:=fibonacci()
	if i == 0 || i==1{
		return i
	}
	return f(i-1)+f(i-2)
}

// func fibonacci to call inner calculator
func fibonacci () func (i int) int {
	return innerCal
}

// main func
func main (){
	f:=fibonacci()
	for i :=0; i<10 ; i++ {
		fmt.Println(f(i))
	}
}