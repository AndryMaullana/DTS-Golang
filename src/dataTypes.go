package main

import "fmt"

func main() {
	/*
		ada 4 tipe data di go diantaranya :
		a. Basic Type : number, string, boolean
		b. Aggregate Type : array and struct
		c. Reference Type : slice, pointer, map, function, channel
		d. Interface Type : interface
	*/
	// NUMBER
	/*
		Number di go dibagi menjadi 2 yaitu
		a. int (integer bilangan non decimal/floating-point)
		b. float (bilangan decimal/floating foint)
	*/
	// Int ( Integer )
	/*
		integer merupakan tipe data numerik non decimal atau non floating-point.
		Di go sendiri tipe data ini dibagi menjadi beberapa sub-kategori diantaranya:
		a. int (bilangan cacah yang hanya menerima nilai positif)
		b. uint (bilangan fositif maupun negatif)

		kemudian int dan uint di pecah lagi menjadi beberapa sub kategori menjadi :
		- int (merepresentasikan int32 dan int64 )
		- uint8
		unsign data sebesar 8-bit (0-255)
		-uint16
		unsign data sebesar 16-bit (0-655355)
		-uint32
		unsign data sebesar 32-bit (0-4294967295)
		-uint64
		unsign data sebesar 64-bit (0-18446744073709551615)
		-int8
		signed data sebesar 8-bit (-128 to 127)
		-int16
		signed data sebesar 16-bit (-32768 to 326767)
		-int32
		signed data sebesar 32-bit (-2147483648 to 2147483647)
		-int64
		signed data sebesar 64-bit (-92233720368547755808 to 92233720368547755807)
	*/
	// Contoh penggunaan tipe data int dan uint pada sebuah variable
	// Int
	var first int = 89
	var second int = -12
	fmt.Printf("tipe data first: %T\n", first)
	fmt.Printf("tipe data second: %T\n", second)
	fmt.Println("================================")
	// uint
	var first1 uint8 = 89
	var second1 int8 = -12
	fmt.Printf("tipe data first: %T\n", first1)
	fmt.Printf("tipe data second: %T\n", second1)
	fmt.Println("================================")

	// Number ( Float )
	/*
		secara umum di go float di bagi menjadi 2 sub yaitu
		- float 32
		-float 64
	*/
	// contoh penggunaan tipe data float dalam variable
	var decimalNumber float32 = 6.36
	fmt.Printf("Decimal number : %f\n", decimalNumber)   // verb %f akan mengambil 6 nilai di belakang decimal
	fmt.Printf("Decimal number : %.3f\n", decimalNumber) // verb %.nf akan mengambil 3 nilai di belakang decimal
	fmt.Println("======================================")

	// Tipe data boolean
	/*
		di go atau di bahasa pemrograman apapun memberikan 2 nilai yaitu true dan false
	*/
	// contoh implentasinya pada sebuah variabel dan mengkases nya menggunakan fmt.Printf()
	var true bool = true
	var false bool = false
	fmt.Printf("is true ? : %t\n", true)   //verb %t untuk format boolean menjadi string
	fmt.Printf("is false ? : %t\n", false) //verb %t untuk format boolean menjadi string
	fmt.Println("============================================")

	// String
	// karakter penulisan tipe data string adalah menggunakan tanda petik dua ("")
	// contoh penerapanya pada variabel
	var String string = "string"
	fmt.Printf("Ini adalah tipe data : %s\n", String)
	fmt.Println("======================================")
	// selain menggunakan tanda petik dua ("") string bisa bisa menggunakan tanda backtik (``) , nah keistimewaan menggunakan becktik adalah ketika ada escape sepert garis baru aka dianggap string
	// contoh string menggunakan backtik
	var stringBacktik string = `ini adalah tipe data string
	menggunakan tanda backtik dimana
	disini bebas menggunakan tab, new line tanpa

	menggunakan escape`
	fmt.Println(stringBacktik)

}
