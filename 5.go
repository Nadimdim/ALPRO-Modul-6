package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	for str != "selesai" {
		fmt.Println(str)
		fmt.Scan(&str)
	}
}
