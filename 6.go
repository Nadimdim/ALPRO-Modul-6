package main

import "fmt"

func main(){
	var nilai int
	var rata2 float64
	var jumlah int = 0
	var n int = 0
	fmt.Scan(&nilai)

	for nilai != -1 {
		n = n + 1
		jumlah = jumlah + nilai
		fmt.Scan(&nilai)
	}
	if n == 0 {
		rata2 = 0.0
	}else {
		rata2 = float64 (jumlah) / float64 (n)
	}
	fmt.Println(rata2)
}