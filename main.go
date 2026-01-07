package main

import (
	"fmt"
	"math/rand"
	"time"

	"matdisk2025-uts-kelompok8/deret"
	"matdisk2025-uts-kelompok8/himpunan"
	"matdisk2025-uts-kelompok8/matrix"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// ===== SOAL 1 =====
	N := 60
	limit := 20

	A := randomSlice(3, limit)
	B := randomSlice(3, limit)
	C := randomSlice(3, limit)

	fmt.Println("Generating Sets... (N=60)")
	fmt.Println("A:", A, "| B:", B, "| C:", C)

	R := himpunan.HitungR(A, B, C, N)
	fmt.Println("Hasil Operasi R:", R)

	filter := himpunan.FilterGenapKurangDari(R, N)
	fmt.Printf("Hasil Filter (Genap < %d): %v\n", N/4, filter)
	fmt.Println("Total Elemen:", len(filter))

	// ===== SOAL 2 =====
	fmt.Println("\nSoal 2: Analisis Pasangan Subset")

	M := 7
	K := 11
	S := randomSlice(M, 30)
	fmt.Println("Set S:", S, "| Target K:", K)

	pairs := himpunan.CariPasanganSubset(S, K)

	for i, p := range pairs {
		fmt.Printf("%d. {%d, %d} (Sum=%d)\n", i+1, p.X, p.Y, p.Sum)
	}
	fmt.Println("Total:", len(pairs), "Pasangan")

	// ===== SOAL 3 =====
	fmt.Println("\nSoal 3: Perkalian dan Trace Matriks")
	fmt.Println("Matrices generated (2x2)...")

	// Masukan matrix A dan B
	a := [][]int{
		{rand.Intn(limit), rand.Intn(limit)},
		{rand.Intn(limit), rand.Intn(limit)},
	}
	b := [][]int{
		{rand.Intn(limit), rand.Intn(limit)},
		{rand.Intn(limit), rand.Intn(limit)},
	}
	fmt.Printf("Matrix A: %v\nMatrix B: %v\n", a, b)

	//  R = A x B
	r := matrix.Multiply(a, b)
	fmt.Printf("Hasil Matrix R: %v\n", r)

	// trace R
	tr := matrix.Trace(r)
	fmt.Printf("Trace R: %d\n", tr)

	// ===== SOAL 4 =====
	fmt.Println("\nSoal 4: Transformasi Baris")

	// Matriks M
	n4 := 5

	m := make([][]int, n4)
	for i := 0; i < n4; i++ {
		m[i] = make([]int, n4)
		for j := 0; j < n4; j++ {
			m[i][j] = rand.Intn(limit) + 1
		}
	}

	fmt.Println("Matrix M (Generated):", m)

	//  N-1
	matrix.SwapRows(m, 0, len(m)-1)
	fmt.Println("Menukar baris 0 dan", len(m)-1)
	fmt.Println("Matrix M Terkini:", m)

	//  nilai maksimum
	max, row, col := matrix.MaxValue(m)
	fmt.Printf("Nilai Maksimum %d ditemukan di posisi (%d, %d)\n", max, row, col)

	// ===== SOAL 5 =====
	fmt.Println("\nSoal 5: Relasi Rekurens Iteratif")

	c1 := 3
	c2 := -2
	n := 12

	// Deret
	hasilDeret, hasilAkhir := deret.Deret(c1, c2, n)

	fmt.Printf("\nINPUT: C1=%d, C2=%d, N=%d\n", c1, c2, n)
	fmt.Println("Proses Perhitungan:")

	// Tampilkan semua suku
	for i, v := range hasilDeret {
		if i == len(hasilDeret)-1 {
			fmt.Printf("Suku %d: %d", i, v)
		} else {
			fmt.Printf("Suku %d: %d | ", i, v)
		}
	}
	fmt.Println()

	fmt.Printf("HASIL AKHIR Suku ke-%d: %d\n", n, hasilAkhir)

	// soal 6 Analisis Kedekatan Deret Geometri
	fmt.Println("\nSoal 6: Analisis Kedekatan Deret Geometri")
	var aGeo, rGeo float64
	var nGeo int

	aGeo = 4
	rGeo = 0.6
	nGeo = 10

	// Panggil fungsi GeometriKedekatan dari package deret
	sn, sinf, persen := deret.GeometriKedekatan(aGeo, rGeo, nGeo)

	fmt.Printf("\nInput Paket: a=%.2f, r=%.2f, N=%d\n", aGeo, rGeo, nGeo)
	fmt.Printf("Sum Berhingga S(%d): %.2f\n", nGeo, sn)
	fmt.Printf("Sum Tak Hingga S(inf): %.2f\n", sinf)
	fmt.Printf("Persentase Kedekatan: %.2f%%\n", persen)

}

func randomSlice(n, limit int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(limit) + 1
	}
	return arr
}
