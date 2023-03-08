package main

import "fmt"

func main() {
	// Looping First Way

	for i := 0; i < 10; i++ {
		fmt.Println("Hello world")
	}
	fmt.Println("============================================")
	// Looping second way
	var number = 0   // deklarasi variabel diletakan diawal
	for number < 3 { // di cara kedua ini menyelipkan kondisi pada keyword for
		fmt.Println("Angka", number) //eksekusi program akan di looping sampai kondisi bernilai true
		number++                     // melakukan counter pada variable yang di loop
	}
	// Looping third way
	fmt.Println("==============================================")
	var n = 0 // deklarasi variable diawal
	for {     // for tanpa kondisi ini akan terus bernilai true yang artinya looping akan terus di jalankan
		fmt.Println("Number", n) // melakukan eksekusi looping dengan mencetak number di variable di atas
		n++                      // counter variable n
		if n == 3 {              // jika nilai n sudah mencapa 3
			break // maka eksekusi looping akan berhenti dengan menggunakan keyword break
		}
	}
	fmt.Println("===============================================")
	// for dengan penggunaan keyword continue
	for i := 1; i <= 10; i++ { // melakukan deklarasi di block for
		if i%2 == 1 { // jika nilai i habis bagi 2 sisa 1
			continue // maka program akan di paksa di paksa melakukan perulangan lagi
		}
		if i > 8 { // kondisi jika nilai i lebih besar dari 8
			break // program akan di paksa berhenti
		}
		fmt.Println("nilai genap :", i) // akan mencetak angka genap
	}
	// For Nested atau bersarang
	fmt.Println("====================================================")
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			for k := j; k < 5; k++ {
				fmt.Print("*")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println("=======================================")
outerloop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perluangan ke -", i+1)
		for j := i; j < 3; j++ {
			if i == 2 {
				break outerloop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}
