package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Program Start ===")

	// SOAL 1 & 2 (Himpunan)
	SolveHimpunan()
	fmt.Println("--------------------------")

	// SOAL 3 & 4 (Matriks)
	fmt.Println("=== Materi II: Matriks ===")
	N := 2
	solveSoal3(N)
	fmt.Println("--------------------------")
	solveSoal4(N)
	fmt.Println("--------------------------")

	// SOAL 5 & 6 (Deret)
	SolveDeret()
}
