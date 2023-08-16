package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 1. buatlah sebuah program untuk menghitung luas trapesium
func luasTrapesium(alasA int, alasB int, tinggi int) {
	rumus := (alasA + alasB) * tinggi / 2
	fmt.Println(rumus, "cm")
}

// 2. buatlah sebuah program untuk menentukan apakah sebuah bilang adalah bilangan ganjil atau genap
func tentukanBilangan(angka int) {
	rumus := angka % 2
	if rumus == 1 {
		fmt.Println(angka, "Merupakan Bilangan Ganjil")
	} else {
		fmt.Println(angka, "Merupakan BIlangan Genap")
	}
}

/*3. - buatlah sebuah program untuk menentukan grade dari sebuah nilai, dengan ketentuan sebagai berikut:
  - Nilai 80 - 100: A
  - Nilai 65 - 79: B
  - Nilai 50 - 64: C
  - Nilai 35 - 49: D
  - Nilai 0 - 34: E
  - Nilai kurang dari 0 atau lebih dari 100 maka tampilkan 'Nilai Invalid' */
func tentukanNilai(angka int) {
	if angka >= 80 {
		fmt.Println("Nilai Kamu Adalah A")
	} else if angka >= 65 {
		fmt.Println("Nilai Kamu ADalah B")
	} else if angka >= 50 {
		fmt.Println("Nilai Kamu ADalah C")
	} else if angka >= 35 {
		fmt.Println("Nilai Kamu ADalah D")
	} else if angka >= 0 {
		fmt.Println("Nilai Kamu ADalah E")
	} else {
		fmt.Println("Invalid")
	}
}

/* 4. buatlah sebuah program yang mencetak angka dari 1 sampai 100 dan untuk kelipatan '3' cetak "Fizz" sebagai ganti angka, dan untuk kelipatan '5' cetak "Buzz”.
Sebagai contoh; 1 2 fizz 4 buzz fizz 7 8 fizz buzz …….*/
func cetakFizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

/* 5. Buatlah sebuah program untuk mencetak segitiga asterik*/
func segitigaAsterisk(angka int) {
	for i := 1; i <= angka; i++ {
		for j := 1; j <= angka-i; j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}

// 6. Buatlah sebuah program untuk mencetak faktor bilangan dari sebuah angka
func faktorBilangan(angka int) {
	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Println(i)
		}
	}
}


// 7. Buatlah sebuah program untuk memeriksa apakah sebuah kata adalah palindrome atau bukan, serta coba terapkan scanner untuk menangkap inputan dari console
func cekPalindrom() {
	fmt.Print("Masukkan sebuah kata: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	s := strings.ToLower(input)
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			fmt.Println("Kata ini adalah bukan Palindrom")
			return 
		}
	}
	fmt.Println("Kata ini adalah Palindrom")
}

func main() {
	luasTrapesium(4, 8, 2)
	tentukanBilangan(6)
	tentukanNilai(63)
	cetakFizzBuzz()
	segitigaAsterisk(10)
	faktorBilangan(8)
	cekPalindrom()
}