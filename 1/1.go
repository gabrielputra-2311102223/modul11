package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Membuat reader untuk membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan data suara (pisahkan dengan spasi, akhiri dengan 0):")

	// Membaca input dari pengguna
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Split input menjadi slice
	data := strings.Split(input, " ")

	// Inisialisasi variabel untuk menghitung suara sah dan tidak sah
	var validVotes, invalidVotes int
	votesCount := make(map[int]int)

	// Loop untuk memproses setiap suara
	for _, v := range data {
		// Konversi string ke integer
		num, err := strconv.Atoi(v)
		if err != nil || num == 0 {
			break
		}
		if num >= 1 && num <= 20 {
			votesCount[num]++
			validVotes++
		} else {
			invalidVotes++
		}
	}

	// Hitung total suara
	totalVotes := validVotes + invalidVotes

	// Output hasil
	fmt.Printf("Suara masuk: %d\n", totalVotes)
	fmt.Printf("Suara sah: %d\n", validVotes)

	for candidate, count := range votesCount {
		fmt.Printf("%d: %d\n", candidate, count)
	}
}
