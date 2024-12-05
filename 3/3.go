package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	// I.S. terdefinisi integer n, dan sejumlah n data sudah siap pada piranti masukan.
	// F.S. Array data berisi n (< NMAX) bilangan
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	// mengembalikan posisi k dalam array data dengan elemen. Posisi dimulai dari posisi 0. Jika tidak ada kembalikan -1
	for i := 0; i < n; i++ {
		if data[i] == k {
			return i
		}
	}
	return -1
}

func main() {
	// buatlah kode utama yang membaca baris pertama (n dan k). Kemudian data diisi oleh prosedur isiArray(), dan pencarian oleh fungsi posisi(n, k), dan setelah itu output dicetak.
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)
	pos := posisi(n, k)
	if pos != -1 {
		fmt.Println(pos)
	} else {
		fmt.Println("TIDAK ADA")
	}
}
