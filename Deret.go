package main

import (
	"fmt"
	"math"
)

func SolveDeret() {
	// --- SOAL 5: RELASI REKURENS ITERATIF ---
	fmt.Println("=== SOAL 5: RELASI REKURENS ITERATIF ===")
	c1, c2, n1 := 1, 3, 11

	a0, a1 := 1, 1

	fmt.Printf("INPUT: C1=%d, C2=%d, N=%d\n", c1, c2, n1)
	fmt.Print("Proses Perhitungan:\n")
	fmt.Printf("Suku 0: %d | Suku 1: %d | ", a0, a1)

	prev2 := a0
	prev1 := a1
	var current int

	for i := 2; i <= n1; i++ {
		current = (c1 * prev1) + (c2 * prev2)
		fmt.Printf("Suku %d: %d", i, current)
		if i < n1 {
			fmt.Print(" | ")
		}
		prev2 = prev1
		prev1 = current
	}
	fmt.Printf("\nHASIL AKHIR Suku ke-%d: %d\n\n", n1, current)

	// --- SOAL 6: ANALISIS KEDEKATAN DERET GEOMETRI ---
	fmt.Println("=== SOAL 6: ANALISIS KEDEKATAN DERET GEOMETRI ===")
	a, r, n2 := 8.0, 0.2, 15.0

	sn := a * (1 - math.Pow(r, n2)) / (1 - r)

	sinf := a / (1 - r)

	kedekatan := (sn / sinf) * 100

	fmt.Printf("Input Paket: a=%.0f, r=%.1f, N=%.0f\n", a, r, n2)
	fmt.Printf("Sum Berhingga S(%.0f): %.2f\n", n2, sn)
	fmt.Printf("Sum Tak Hingga S(inf): %.2f\n", sinf)
	fmt.Printf("Persentase Kedekatan: %.2f%%\n", kedekatan)
}
