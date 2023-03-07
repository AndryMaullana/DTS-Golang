package main

import "fmt"

func main() {
	// Fungsi print hampir sama kegunaanya dengan println yang membedakaya adalah fungsi print tidak menghasilkan garis baru di akhir output nya.
	// contoh perbedaan antara println() dan print()
	fmt.Println("Airell Jordan") // <- penggunaan Println
	// Output yang dihasilkan : Airell Jordan
	fmt.Println("Airell", "Jordan") // <- Penggunaan Println
	// Output yang dihasilkan : Airell Jordan

	fmt.Print("Airell Jordan\n") // penggunaan menggunakan Print()
	// Output yang dihasilkan jika menggunakan tanda \n : Airell Jordan ( dengan garis baru dibawahnya )
	// Output yang dihasilkan tanpa tanda \n : Airell Jordan ( tanpa ada garis baru sehingga output yang di bawah nya akan bergabung )
	fmt.Print("Airell", " Jordan\n")     // penggunaan menggunakan Print()
	fmt.Print("Airell", " ", "Jordan\n") // penggunaan menggunakan Print()
}
