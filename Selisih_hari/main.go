package main

import "fmt"

func main() {
	fmt.Println("Selisih tanggal")
	fmt.Println("-------------------------------------------")

	januari := 31
	bulan1 := 1 // 1 = januari
	bulan2 := 2 // 2 = februari
	tglBulan1 := 3
	tglBulan2 := 10

	if bulan1 != bulan2 {
		fmt.Println(januari - tglBulan1 + tglBulan2)
	} else {
		fmt.Println(tglBulan2 - tglBulan1)
	}

}
