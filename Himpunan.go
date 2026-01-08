package main

import (
	"fmt"
	"math/rand"
)

// --- FUNGSI HELPER (OPERASI HIMPUNAN MANUAL) ---

func exists(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func symmetricDifference(A, B []int) []int {
	res := []int{}
	for _, v := range A {
		if !exists(B, v) {
			res = append(res, v)
		}
	}

	for _, v := range B {
		if !exists(A, v) {
			res = append(res, v)
		}
	}
	return res
}

func difference(set, C []int) []int {
	res := []int{}
	for _, v := range set {
		if !exists(C, v) {
			res = append(res, v)
		}
	}
	return res
}

func intersection(A, C []int) []int {
	res := []int{}
	for _, v := range A {
		if exists(C, v) {
			res = append(res, v)
		}
	}
	return res
}

func union(s1, s2 []int) []int {
	res := append([]int{}, s1...)
	for _, v := range s2 {
		if !exists(res, v) {
			res = append(res, v)
		}
	}
	return res
}

func SolveHimpunan() {
	N := 140
	limit := N

	generateSet := func() []int {
		s := make([]int, 5)
		for i := range s {
			s[i] = rand.Intn(limit) + 1
		}
		return s
	}

	A := generateSet()
	B := generateSet()
	C := generateSet()

	step1 := symmetricDifference(A, B)
	step2 := difference(step1, C)
	step3 := intersection(A, C)
	R := union(step2, step3)

	fmt.Println("--- SOAL 1 ---")
	fmt.Printf("Generating Sets... (N=%d)\n", N)
	fmt.Printf("A: %v | B: %v | C: %v\n", A, B, C)
	fmt.Printf("Hasil Operasi R : %v\n", R)

	filterVal := N / 4
	filteredR := []int{}
	for _, v := range R {
		if v%2 == 0 && v < filterVal {
			filteredR = append(filteredR, v)
		}
	}
	fmt.Printf("Hasil Filter (Genap < %d): %v\n", filterVal, filteredR)
	fmt.Printf("Total Elemen: %d\n\n", len(filteredR))

	M := 7
	K := 15
	S := make([]int, M)
	for i := range S {
		S[i] = rand.Intn(20) + 1
	}

	fmt.Println("--- SOAL 2 ---")
	fmt.Printf("Set S: %v | Target K: %d\n", S, K)
	fmt.Println("Subset 2-Elemen (Sum > 15):")

	count := 0
	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			sum := S[i] + S[j]
			if sum > K {
				count++
				fmt.Printf("%d. {%d, %d} (Sum=%d)\n", count, S[i], S[j], sum)
			}
		}
	}
	fmt.Printf("Total: %d Pasangan\n", count)
}
