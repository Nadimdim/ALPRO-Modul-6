package main

import "fmt"

func main() {
	var S string
	var A, B int

	fmt.Scan(&S)
	fmt.Scan(&A, &B)
	fmt.Println("kata =", S)
	fmt.Println("jumlah =", A+B)
}
