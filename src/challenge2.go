package main

import "fmt"

func main() {
	for i := 0; i <= 4; i++ {
		fmt.Println("Nilai i = ", i)
	}
	for j := 0; j <= 10; j++ {
		// if j == 5 {
		// teks := []rune{'C', 'A', 'M', 'P', 'B', 'O'}
		// for indeks := 0; indeks < len(teks); indeks++ {
		// 	char := teks[indeks]
		// 	fmt.Printf("Caracter %U '%s' starts at byte position %d \n", char, string(char), indeks*2)
		// }
		// } else {
		// 	fmt.Println("Nilai j = ", j)
		// }

		if j == 5 {
			Caracter := []rune{'\u0421', '\u0410', '\u0428', '\u0410', '\u0420', '\u0412', '\u041E'}
			for pos, char := range Caracter {
				fmt.Printf("character %U '%c' starts at byte position %d\n", char, char, pos*2)
			}
		} else {
			fmt.Println("Nilai j : ", j)
		}
	}
}
