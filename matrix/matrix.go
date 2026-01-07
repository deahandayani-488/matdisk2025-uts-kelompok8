package matrix

// Multiply melakukan perkalian dua matriks
func Multiply(a, b [][]int) [][]int {
	n := len(a)
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum += a[i][k] * b[k][j]
			}
			r[i][j] = sum
		}
	}
	return r
}

// Trace menghitung trace matriks (jumlah diagonal utama)
func Trace(m [][]int) int {
	t := 0
	for i := 0; i < len(m); i++ {
		t += m[i][i]
	}
	return t
}

// soal nomor 4 Transformasi Baris

// Tukar baris 0 dengan baris N
func SwapRows(m [][]int, r1, r2 int) [][]int {
	m[r1], m[r2] = m[r2], m[r1]
	return m
}

func MaxValue(m [][]int) (max, row, col int) {
	max = m[0][0]
	row = 0
	col = 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] > max {
				max = m[i][j]
				row = i
				col = j
			}
		}
	}
	return
}
