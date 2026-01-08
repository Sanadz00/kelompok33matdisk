package main

import (
	"fmt"
	"math/rand"
)

// --- SOAL 3: Perkalian dan Trace Matriks ---

func solveSoal3(n int) {
	fmt.Printf("Matrices generated (%dx%d)...\n", n, n)

	matrixA := make([][]int, n)
	matrixB := make([][]int, n)
	for i := range matrixA {
		matrixA[i] = make([]int, n)
		matrixB[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrixA[i][j] = rand.Intn(10) + 1
			matrixB[i][j] = rand.Intn(10) + 1
		}
	}

	fmt.Printf("Matrix A: %v\n", matrixA)
	fmt.Printf("Matrix B: %v\n", matrixB)

	resultR := make([][]int, n)
	for i := range resultR {
		resultR[i] = make([]int, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				resultR[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}
	fmt.Printf("Hasil Matriks R: %v\n", resultR)

	trace := 0
	for i := 0; i < n; i++ {
		trace += resultR[i][i]
	}
	fmt.Printf("Trace: %d\n\n", trace)
}

// --- SOAL 4: Transformasi Baris ---

func solveSoal4(n int) {
	matrixM := make([][]int, n)
	for i := range matrixM {
		matrixM[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrixM[i][j] = rand.Intn(20) + 1
		}
	}

	fmt.Printf("Matrix M (Generated): %v\n", matrixM)
	fmt.Printf("Menukar Baris 0 dan %d...\n", n-1)

	matrixM[0], matrixM[n-1] = matrixM[n-1], matrixM[0]

	fmt.Printf("Matrix M Terkini: %v\n", matrixM)

	maxVal := matrixM[0][0]
	posX, posY := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrixM[i][j] > maxVal {
				maxVal = matrixM[i][j]
				posX, posY = i, j
			}
		}
	}
	fmt.Printf("Nilai Maksimum: %d ditemukan di Posisi (%d,%d)\n", maxVal, posX, posY)
}
