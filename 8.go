package main

import "fmt"

func main(){
	var r, t int
	fmt.Scan(&r, &t)
	fmt.Println(hitung_volume(r, t))
}
func hitung_volume(r int, t int)float64{
	var pi float64 =3.14
	return float64(r)*float64(r)*float64(pi)*float64(t)
}