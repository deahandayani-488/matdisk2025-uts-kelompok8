package deret

import "math"

// Deret menghitung relasi rekurens iteratif
// a(n) = c1 * a(n-1) + c2 * a(n-2)
// dengan a(0) = 1, a(1) = 1
func Deret(c1, c2, n int) ([]int, int) {
	if n < 0 {
		return nil, 0
	}
	// Jika n = 0 cukup 1 elemen
	if n == 0 {
		return []int{1}, 1
	}

	// Buat slice untuk menyimpan semua suku
	a := make([]int, n+1)

	// Minimal butuh a0 dan a1
	a[0] = 1
	a[1] = 1

	for i := 2; i <= n; i++ {
		a[i] = c1*a[i-1] + c2*a[i-2]
	}
	return a, a[n]
}

// Analisis Kedekatan Deret Geometri
// Menghitung Sn (sum berhingga), Sinf (sum tak hingga), dan persentase kedekatan
func GeometriKedekatan(a, r float64, n int) (sn, sinf, persen float64) {
	if n <= 0 {
		return 0, 0, 0
	}

	// Kasus r == 1: deret tidak konvergen, Sn = a * n
	if r == 1 {
		sn = a * float64(n)
		sinf = 0 // tidak ada sum tak hingga untuk r=1
		persen = 0
		return
	}

	// Hitung Sum Berhingga: Sn = a * (1 - r^n) / (1 - r)
	sn = a * (1 - math.Pow(r, float64(n))) / (1 - r)

	// Hitung Sum Tak Hingga: Sinf = a / (1 - r) hanya jika |r| < 1
	if math.Abs(r) < 1 {
		sinf = a / (1 - r)
		// Hitung persentase kedekatan
		if sinf != 0 {
			persen = (sn / sinf) * 100
		}
	} else {
		// Jika |r| >= 1, deret divergen, tidak ada sum tak hingga
		sinf = 0
		persen = 0
	}

	return
}
