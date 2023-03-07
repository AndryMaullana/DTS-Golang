package main

import "fmt"

func main() {
	/*
		Constan merupakan jenis variable yang dimana setiap value atau nilai nya tidak dapat berubah alias constan.
		Di go penulisan constan cukup mudah dan hampir sama dengan penulisan variable, yang membedakan hanya penamaan keyword nya saja
	*/
	// contoh pendeklarasian constanta
	const nilai_constan = 90
	fmt.Printf("Nilai constanta : %d\n", nilai_constan)
	fmt.Println("====================================")
	/*
		sama seperti variable, const juga dapat melakukan multiple deklarasi
	*/
	// contoh
	const consta1, constan2, constan3 = "konstan1", "konstan2", "konstan3"
	fmt.Printf("constan1 : %s\nconstant2 : %s\nconstant3 : %s\n", consta1, constan2, constan3)
	fmt.Println("===================================================")
}
