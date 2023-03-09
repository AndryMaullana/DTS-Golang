package main

import (
	"fmt"
	"strings"
	// "strings"
)

func main() {
	// var slicess = []string{"array", "slice", "map"}
	// _ = slicess
	// fmt.Printf("%#v\n", slicess)
	// fmt.Println(strings.Repeat("#", 25))

	// var fruits = make([]string, 3)
	// fruits[0] = "Semangka"
	// fruits[1] = "Apel"
	// fruits[2] = "Anggur"
	// _ = fruits
	// fmt.Printf("%#v\n", fruits)

	// var fruits1 = []string{"Apel", "Mangga", "Anggur"}
	// var fruits2 = []string{"Semangka", "Pear", "Rambutan"}

	// nn := copy(fruits1, fruits2)
	// fmt.Println("fruits ==> ", fruits1)
	// fmt.Println("fruits2 ==> ", fruits2)
	// fmt.Println("nn ==>", nn)
	// fruits1 = append(fruits1, fruits2...)
	// fmt.Printf("%#v\n", fruits1)

	var buah = []string{"Apel", "Mangga", "Anggur", "Nanas", "Strawberry"}

	var buah1 = buah[1:4]
	var buah2 = buah[0:]
	var buah3 = buah[:3]
	var buah4 = buah[:]

	fmt.Println("Buah", buah)
	fmt.Println("Buah 1 =>", buah1)
	fmt.Println("Buah 2 =>", buah2)
	fmt.Println("Buah 3 =>", buah3)
	fmt.Println("Buah 4 =>", buah4)

	fruits := []string{"Apel", "Semangka", "Anggur", "durian", "Melon"}
	fruits1 := append(fruits[1:3], "Rambutan")
	fmt.Printf("%#v\n", fruits1)
	fmt.Printf("%#v\n", fruits)
	fruits2 := append(fruits1[1:2], "watermelon")
	fmt.Printf("%#v\n", fruits2)
	fmt.Printf("%#v\n", fruits1)
	fmt.Printf("%#v\n", fruits)

	// Slice Cap Function
	fmt.Println(strings.Repeat("#", 30))
	var benda = []string{"pisau", "pen", "meja", "mouse", "buku"}
	fmt.Println("Cap Slice Benda => ", cap(benda))
	fmt.Println("Len Slice Benda => ", len(benda))

	var benda1 = benda[1:3]
	fmt.Println("Cap Slice Benda1 =>", cap(benda1))
	fmt.Println("Len Slice Benda1 =>", len(benda1))
	fmt.Printf("%#v\n", benda1)

	var benda2 = benda[2:5]
	fmt.Println("Cap Slice Benda2 =>", cap(benda2))
	fmt.Println("Len Slice Benda2 =>", len(benda2))
	fmt.Printf("%#v\n", benda2)

	// Slice new Backing Array
	fmt.Println(strings.Repeat("#", 25))
	cars := []string{"ford", "honda", "Audi", "range over"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)
	// fmt.Printf("%#v\n", newCars)
	cars[0] = "Nissan"
	fmt.Println("cars :", cars)
	fmt.Println("newCars :", newCars)
}
