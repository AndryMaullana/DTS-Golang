package main

import "fmt"

func main() {
	// Operator aritmatika
	/*
		di Go ada beberapa operator aritmatika dan mungkin hampir sama di bahasa pemrograman yang lain.
		berikut beberapa simbol operator aritmatika
			a. + -> untuk penjumlahan
			b. - -> untuk pengurangan
			c. * -> untuk perkalian
			d. / -> untuk pembagian
			e. % -> untuk modulos ( sisa bagi )
			f. ++ -> increment
			g. -- -> decrement
	*/
	// contoh pengunaan operator aritmatika
	var value int = (2 + 2) * 3
	fmt.Println(value)
	// contoh penggunaan operator aritmatikan pada variable
	var firstOperan int = 2 + 2
	var secondOperan = 2 - 2
	var thirdOperan = 2 * 2
	var fourOperan = 2 / 2
	var fiveOperan = 2 % 2
	// var sixOperan = ++1
	// var sevenOperan = --1

	fmt.Println(firstOperan)
	fmt.Println(secondOperan)
	fmt.Println(thirdOperan)
	fmt.Println(fourOperan)
	fmt.Println(fiveOperan)
}
