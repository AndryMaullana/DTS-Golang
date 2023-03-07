package main

import "fmt"

func main() {
	// IOTA
	/*
		IOTA pada golang di representasikan sebagai rangkaian peningkatan sebuah variable konstan ( increment ). IOTA memiliki nilai default yaitu 0
	*/
	// contoh peggunaan IOTA pada variable const
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)
	/*
		IOTA dapat digunakan menggunakan sebagai perhitungan dengan operator matematika.
	*/
	// contoh penggunaan IOTA dengan operasi matematika
	const (
		c11 = iota + 2
		c22
		c33
	)
	const (
		c44 = iota * 2
		c55
		c66
	)
	fmt.Println(c11, c44)
	const (
		c77 = iota - 2
		c88
		c99
	)

}
