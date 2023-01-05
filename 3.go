package main

import "fmt"

func main(){
	var x float64
	fmt.Scan(&x)
	if int (x) % 3 == 0 {
		fmt.Print("Fizz")
	}
	if int (x) % 5 == 0 {
		fmt.Print("Bazz")
	}
	
}