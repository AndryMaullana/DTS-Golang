package main

import "fmt"

func main() {
	// Kondisional
	/*
		kondisional merupakan suatu alur bagaimana seharusnya program berjalan
	*/
	// contoh penggunaan kondisional menggunakan if else
	var currentYear = 2023                   // membuat variable untuk patokan pada kondisi
	if age := currentYear - 2000; age > 17 { //di dalam if dapat membuat variable yang dimana variable tersebut hanya bisa diakses oleh block if tersebut
		// pada if di suatu block kondisi dapat dipisah denngan tanda ;
		fmt.Println("Kamu Sudah Boleh Membuat Kartu Sim") // kosidisi dimana jika bernilai false/ else
	} else {
		fmt.Println("Kamu Belum Boleh Membuat Kartu Sim") // kondisi dimana jika bernilai true
	}

}
