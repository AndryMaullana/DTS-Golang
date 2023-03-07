package main

import "fmt"

func main() {
	// ada beberapa cara dalam penulisan variable di golang
	/*
		dalam pembuatan variable di golang perlu menggunakan penulisan yang tepat agar setiap komponen mudah di baca dan di pahami. dan penulisan variable yang effisien da banyak digunakan adalah teknik camelCase

		camelCase adalah konvensi penamaan variable di mana huruf pertama daris setiap kata majemuk akan di kapitalisasi kecuali untuk huruf pertama
		contoh	:
	*/

	var (
		favoriteColor = "green"     //-> perhatikan huruf diawal ditulis dengan huruf kecil
		favoriteShape = "rectangle" //-> sedangkan huruf "C" dan "S" ditulis dengan kapital
	)
	// cara mengakses sebuah variable dengan menggunakan println
	fmt.Println(favoriteColor)
	fmt.Println(favoriteShape)

	/*
		ada dua cara untuk pendeklarasian variable di go yang pertama dengan menyebutkan type data dan yang kedua tanpa menyebutkan type data nya
	*/
	//  Declaring variable dengan type data
	var name string = "Andri Maulana" // Melakukan deklarasi penamaan , type data, dan value
	var age int = 23                  // Melakukan deklarasi penamaan , type data, dan value
	fmt.Println("ini adalah namanya ==>", name)
	fmt.Println("ini adalah umurnya ==>", age)

	// Deklarasi variable dengan tipe data.
	// melakukan resaign sebuah value pada variable.
	var name2 string // melakukan deklarasi nama dan type data nya terlebih dahulu.
	var age2 int     // melakukan deklarasi nama dan type data nya terlebih dahulu.
	name2 = "Andri"  // Melakukan asign value terhadap variable yang sudah di deklarasikan diatas.
	age2 = 32        // Melakukan asign value terhadap variable yang sudah di deklarasikan diatas.
	fmt.Println("ini adalah re asign variable nama2 ==>", name2)
	fmt.Println("Ini adalah re asign variabke age2 ==>", age2)

	// Variable without data type
	// Cara penulisan variable ini biasa disebut variable inference yang dimana saat kita melakukan deklarasi variable tanpa type data maka pada sistem golang akan secara otomatis mendeteksi tipe data apa yang kita gunakan tergantung value atau nilai yang dimasukan di variable tersebut.
	// contoh variable inference
	var name3 = "Andra"
	var age3 = 34
	var isTrue = true
	fmt.Printf("%T, %T, %T", name3, age3, isTrue)

	// Variable without data type and var init
	// Short declaration, dalam teknik ini kita tidak perlu menggunakan lagi keyword var dalam pendeklarsian variable nya jadi bisa digantikan dengan tanda : sesudah melakukan penamaan variable nya
	// Contoh teknik short declaration
	name4 := "Andri Maulana"
	age4 := 32
	fmt.Printf("%T, %T", name4, age4)
	fmt.Println("")

	// Multiple declaration
	// selain hanya satu di golang dapat melakukan banya deklarasi variable dalama 1 baris dengan syarat type data yang akan di deklarasi harus satu type
	// contoh multiple declarasi dalam satu baris
	var studen, studen1, studen2, studen3 string = "satu", "2", "tiga", "empat"
	fmt.Println(studen, studen1, studen2, studen3)
	// namun jika ingin melakukan multiple declaration dengan berbeda nilai type data nya maka bisa menggunakan type inference atau short declaration
	// contoh multiple declaration dengan type inference
	var a, b, c, d = "1", 2, "3", true
	fmt.Printf("type nilai variable a adalah = %T\ntype nilai variable b adalah = %T\ntype nilai variable c adalah = %T\ntype nilai variable d adalah = %T\n", a, b, c, d)
	fmt.Println("===========================================================")
	// contoh multiple declaration dengan short declaration
	e, f, g, h := "1", 2, "3", true
	fmt.Printf("type nilai variable e adalah = %T\ntype nilai variable f adalah = %T\ntype nilai variable g adalah = %T\ntype nilai variable h adalah = %T\n", e, f, g, h)

	// Undercore variable
	// di Golang jika sudah melakukan deklarasi variable mau tidak mau variable tersebut harus digunakan alias tidak boleh nganggur. Jika tetap dipaksakan tidak digunakan maka program golang akan menyatakan error. Nah cara mengatasi hal seperti itu bisa melakukan denga cara menuliskan underscore pada variable yang sekiranya tidak akan digunakan
	// contohnya
	var variabelNganggur1, variabelNganggur2 = "test", "test"
	_, _ = variabelNganggur1, variabelNganggur2

	// Tambahan
	/* Penggunaan flag pada fmt.Printf, ada beberapa flag pada printf dan salah satunya sudah di gunakan di atas. berikut flag pada printf diantaranya :
	%T => untuk mengetahui jenis tipe data pada sebuah variable
	%s => untuk memformat pada nilai suatu string pada variable
	%d => untuk memformat pada nilai suatu int pada variable (nilai yg ditampung : 0 - 9)
	*/
	// contoh penggunaan flag %s dan %d
	nama := "Andri Maulana"
	umur := 23

	fmt.Printf("Hallo perkenalkan nama saya %s, saya sekarang berumur %d tahun, senang mengenal kalian!!", nama, umur)
	/*
		Output nya akan sepert ini :
		Hallo perkenalkan nama saya Andri Maulana (diambil dari variable nama), saya sekarang berumur 23(diambil dari variable umur) tahun, senang mengenal kalian!!
	*/
}
