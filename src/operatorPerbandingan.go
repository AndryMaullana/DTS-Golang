package main

import "fmt"

func main() {
	// Operator Perbandingan
	/*
		Operator perbandingan pada golang dapat mengembalikan 2 nilai bool yaitu false dan true
		berikut beberapa simbol operator perbandingan
			a. == <- cek apakah 2 operan bernilai sama atau type sama (A==B) is FALSE
			b. != <- cek apakah salah satu operan tidak sama dengan (A!=B) is True
			c. > <- simbol lebih dari value salah satu operan (A>B) is False
			d. < <- simbol kurang dari
			e. >=<- simbo lebih dari sama dengan
			f. <=<- simbol kurang dari sama dengan
	*/
	// contoh penggunaan operator perbandingan
	fmt.Println(true == true)
	fmt.Println(true != true)
	fmt.Println(3 > 5)
	fmt.Println(3 < 5)
	fmt.Println(3 >= 5)
	fmt.Println(3 <= 5)
	// contoh penggunaan operator perbandingan pada variable
	var firstConstion bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11
	// cetak variable
	fmt.Println(firstConstion)
	fmt.Println(secondCondition)
	fmt.Println(thirdCondition)
	fmt.Println(fourthCondition)
}
