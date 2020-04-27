package main

import "fmt"

func main(){
	fmt.Println("Selisih tanggal")
	fmt.Println("-------------------------------------------")

	januari := 31
	bulan1 := "januari"
	bulan2 := "februari"
	tgl1 := 3 
	tgl2 := 10

	if bulan1 != bulan2{
		fmt.Println(januari - tgl1 + tgl2)
	}else if bulan1 == bulan2{
		fmt.Println(tgl2-tgl1)
	}
	
}