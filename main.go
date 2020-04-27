package main

import "fmt"

func main(){
	// Kurs dolar ke rupiah
	dolar := 5
	rupiah := 15524 * dolar
	fmt.Println("Kurs dolar ke rupiah")
	fmt.Println("-------------------------------------------")
	fmt.Println("Rp", rupiah)
	fmt.Println("===========================================")

	// Luas lingkaran
	phi := 3.14
	r := 5.2
	luas := phi * r * r

	fmt.Println("Luas lingkaran")
	fmt.Println("-------------------------------------------")
	fmt.Println(luas)
	fmt.Println("===========================================")

	// FIZZ dan BUZZ
	fmt.Println("FizzBuzz")
	fmt.Println("-------------------------------------------")
	for i := 1; i <= 100; i++ {
        if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		}else if i % 5 == 0 {
			fmt.Println("Buzz")
		}else if i % 3 == 0 {
			fmt.Println("Fizz")
		}else{
			fmt.Println(i)
		}
    }

	fmt.Println("===========================================")

	// Jumlah hari diantara 2 tanggal
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