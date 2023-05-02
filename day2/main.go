package main

import (
	"fmt"
)

func main() {
	fmt.Print("Hello World")
	// Contoh comment dalam bahasa go
	// fundamental pertama : Variable
	// diisi dalam RAM

	// titik koma (;) itu optional di Go.
	// kalau di save pake extensi Go bakal ilang

	// go ada unsigned integer
	// go ada int8, int16, int32, int64
	// juga ada float32, float64
	// plus complex32, complex64
	// bool dan string

	var name string = "Alterra"
	var alamat string
	alamat = "malang"

	// golang itu ada zero value pada setiap variable
	// jadi tidak ada isu null
	// ada case nil ; nanti di struktur data

	// Buat variabel short declaration dengan
	// namavar := value
	// harus ada valuenya
	alamat2 := "jakarta"

	fmt.Print(name)
	fmt.Print(alamat)
	fmt.Print(alamat2)

	// operator => processor
	// bisa langsung data ataupun variabel.

	alas := 10
	tinggi := 20
	// kalau luas := 0.5 * alas * tinggi
	// gak sesuai tipe datanya.
	// jadi harus disesuaikan tipedata()
	luas := 0.5 * float32(alas) * float32(tinggi)

	fmt.Print(luas)

	// seputar bitwise, ingat pergeseran
	// bilangan biner (bit)

	// const
	const (
		pi float64 = 3.14
		pi2
		age = 10
	)

	// **struktur kontrol**
	// kondisi
	// kondisi tidak usah di kurung ()

	var a, b int32
	if a > b {
		fmt.Println("Grand")
	} else if a == b {
		fmt.Println("Nah")
	} else {
		fmt.Println("Bummer")
	}

	var c, d, e int32
	if c == d && c == e {
		fmt.Println("All Equals")
	} else if c != d || c != e || d != e {
		fmt.Println("Some Not Equals")
	} else {
		fmt.Println("Not Possible to be Here")
	}

	// loops
	// for untuk iterasi/inkremen
	// i++ itu i = i + 1

	for i := 1; i <= 10; i++ {
		fmt.Println("WeTons")
	}

	// gak ada while di Go.
	// gunakan for

}
