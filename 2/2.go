package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Input Data
	var input string
	fmt.Print("Masukkan data suara: ")
	fmt.Scanln(&input)

	// Split Input String
	data := strings.Split(input, " ")

	// Initialize Variables
	var suaraMasuk, suaraSah int
	var ketuaRT, wakilKetua int
	var calonSuara map[int]int = make(map[int]int)

	// Process Input Data
	for _, v := range data {
		// Check if Input is Valid
		if v == "0" {
			break
		}

		// Convert String to Integer
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error converting input to integer:", err)
			os.Exit(1)
		}

		// Check if Input is Valid
		if num >= 1 && num <= 20 {
			// Increment Suara Sah Count
			suaraSah++
			// Increment Calon Suara
			calonSuara[num]++
		}

		// Increment Suara Masuk Count
		suaraMasuk++
	}

	// Determine Ketua RT and Wakil Ketua
	ketuaRT = -1
	wakilKetua = -1
	maxSuara := 0
	secondMaxSuara := 0

	// Find Maximum Suara
	for calon, suara := range calonSuara {
		if suara > maxSuara {
			wakilKetua = ketuaRT
			secondMaxSuara = maxSuara
			ketuaRT = calon
			maxSuara = suara
		} else if suara > secondMaxSuara && suara != maxSuara {
			wakilKetua = calon
			secondMaxSuara = suara
		}
	}

	// Output
	fmt.Println("Suara masuk:", suaraMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketuaRT)
	fmt.Println("Wakil ketua:", wakilKetua)
}
