package main

import (
	"fmt"
	"math/rand"
)

//
// BubbleSortBasic
// Sapuan kiri → kanan (mendorong nilai terbesar ke kanan).
// Selalu melakukan (n-1) pass, dan setiap pass membandingkan sepanjang array.
// Setelah pass ke-1, a[n-1] adalah nilai terbesar; pass ke-2, a[n-2] terbesar ke-2; dst.
//
func BubbleSortBasic(a []int) {
	n := len(a)
	if n < 2 {
		return
	}
	for pass := 0; pass < n-1; pass++ {
		for j := 1; j < n; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

//
// BubbleSortReducedComparisons
// Sapuan kiri → kanan dengan pengurangan perbandingan tiap pass.
// Batas kanan menyusut: pass-1 bandingkan j=1..n-1, pass-2 j=1..n-2, dst.
// Setelah pass ke-(pass+1), elemen terbesar ke-(pass+1) berada di indeks n-1-pass.
//
func BubbleSortReducedComparisons(a []int) {
	n := len(a)
	if n < 2 {
		return
	}
	for pass := 0; pass < n-1; pass++ {
		for j := 1; j < n-pass; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

//
// BubbleSortEarlyExit
// Sapuan kiri → kanan dengan pengurangan perbandingan dan early-exit.
// Jika pada suatu pass tidak ada swap, array sudah terurut dan proses berhenti.
// Best-case menjadi O(n), worst/average-case tetap O(n^2).
//
func BubbleSortEarlyExit(a []int) {
	n := len(a)
	if n < 2 {
		return
	}
	for pass := 0; pass < n-1; pass++ {
		swapped := false
		for j := 1; j < n-pass; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

//
// cloneInts
// Membuat salinan slice agar setiap varian sort diuji pada data yang sama.
//
func cloneInts(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func main() {
	data := []int{9, 1, 5, 3, 8, 2, 7, 4, 6, 0}
	fmt.Println("Original                       :", data)

	a := cloneInts(data)
	BubbleSortBasic(a)
	fmt.Println("Basic (max→kanan, full)        :", a)

	b := cloneInts(data)
	BubbleSortReducedComparisons(b)
	fmt.Println("ReducedCmp (max→kanan)         :", b)

	c := cloneInts(data)
	BubbleSortEarlyExit(c)
	fmt.Println("EarlyExit (max→kanan)          :", c)

	almost := []int{0, 1, 2, 4, 3, 5, 6, 7, 8, 9}
	fmt.Println("\nAlmost sorted                  :", almost)
	BubbleSortEarlyExit(almost)
	fmt.Println("EarlyExit result               :", almost)

	rng := rand.New(rand.NewSource(42))
	random := make([]int, 10)
	for i := range random {
		random[i] = rng.Intn(100)
	}
	fmt.Println("\nRandom before                  :", random)
	BubbleSortEarlyExit(random)
	fmt.Println("Random sorted (max→kanan)      :", random)
}
