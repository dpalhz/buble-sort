package main

import (
	"fmt"
	"math/rand"
)

// BubbleSortBasic
// Versi dasar: melakukan (n-1) pass, tiap pass membandingkan (n-1) elemen.
// Implementasi menggeser nilai kecil ke kiri (awal array).
// Setelah pass-1, elemen terkecil ada di indeks 0; setelah pass-2, dua terkecil di 0..1, dst.
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

// BubbleSortReducedComparisons
// Modifikasi #3: jumlah perbandingan berkurang tiap pass.
// Pass-1: (n-1), Pass-2: (n-2), ..., Pass-(n-1): 1.
// Karena setelah pass-ke-k, k elemen terkecil sudah fix di depan (0..k-1),
// loop dalam dimulai dari j = pass+1.
func BubbleSortReducedComparisons(a []int) {
	n := len(a)
	if n < 2 {
		return
	}
	for pass := 0; pass < n-1; pass++ {
		for j := pass + 1; j < n; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// BubbleSortEarlyExit
// Modifikasi #4 (optimal): ReducedComparisons + early-exit.
// Jika dalam satu pass tidak ada swap, array sudah terurut → hentikan.
// Best-case O(n) saat array sudah/hampir terurut.
func BubbleSortEarlyExit(a []int) {
	n := len(a)
	if n < 2 {
		return
	}
	for pass := 0; pass < n-1; pass++ {
		swapped := false
		for j := pass + 1; j < n; j++ {
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

// cloneInts: helper agar tiap varian diuji di data yang sama
func cloneInts(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func main() {
	// Dataset contoh (10 elemen)
	data := []int{9, 1, 5, 3, 8, 2, 7, 4, 6, 0}
	fmt.Println("Original        :", data)

	// Versi dasar
	a := cloneInts(data)
	BubbleSortBasic(a)
	fmt.Println("Basic           :", a)

	// Versi pengurangan perbandingan per pass
	b := cloneInts(data)
	BubbleSortReducedComparisons(b)
	fmt.Println("ReducedCmp      :", b)

	// Versi optimal (pengurangan + early exit)
	c := cloneInts(data)
	BubbleSortEarlyExit(c)
	fmt.Println("EarlyExit       :", c)

	// Contoh best-case: hampir terurut → harus cepat selesai dengan early-exit
	almostSorted := []int{0, 1, 2, 4, 3, 5, 6, 7, 8, 9}
	fmt.Println("\nAlmost sorted   :", almostSorted)
	BubbleSortEarlyExit(almostSorted)
	fmt.Println("EarlyExit result:", almostSorted)

	// Uji data acak (deterministik) dengan RNG lokal — TIDAK menggunakan rand.Seed global
	rng := rand.New(rand.NewSource(42)) // ganti 42 kalau mau urutan lain
	random := make([]int, 10)
	for i := range random {
		random[i] = rng.Intn(100)
	}
	fmt.Println("\nRandom          :", random)
	BubbleSortEarlyExit(random)
	fmt.Println("Sorted          :", random)
}
