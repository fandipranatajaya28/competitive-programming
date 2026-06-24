package main

import "fmt"

const MOD int64 = 1_000_000_007

func zigZagArrays(n int, l int, r int) int {
	m := r - l + 1

	// Length 1:
	// Any value in [l, r] is valid.
	if n == 1 {
		return m
	}

	// If there is only one possible value and n > 1,
	// adjacent values must be different, so answer is 0.
	if m == 1 {
		return 0
	}

	size := 2 * m

	// Vector layout:
	//
	// vec[0 ... m-1]       = up values
	// vec[m ... 2*m-1]     = down values
	//
	// Base case is length 2.
	//
	// up[i] = number of previous values smaller than i
	// down[i] = number of previous values greater than i
	vec := make([]int64, size)

	for i := 0; i < m; i++ {
		vec[i] = int64(i)           // up[i]
		vec[m+i] = int64(m - 1 - i) // down[i]
	}

	// If n == 2, directly sum base vector.
	if n == 2 {
		return int(sumVector(vec))
	}

	// Build transition matrix.
	//
	// Matrix transforms:
	// old vec -> new vec
	//
	// newUp[i] = sum oldDown[j] for j < i
	// newDown[i] = sum oldUp[j] for j > i
	matrix := make([][]int64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int64, size)
	}

	for i := 0; i < m; i++ {
		// Build row for newUp[i].
		// It depends on oldDown[0 ... i-1].
		for j := 0; j < i; j++ {
			matrix[i][m+j] = 1
		}

		// Build row for newDown[i].
		// It depends on oldUp[i+1 ... m-1].
		for j := i + 1; j < m; j++ {
			matrix[m+i][j] = 1
		}
	}

	// We already have DP for length 2.
	// Need to apply transition n-2 more times.
	resultVec := matrixPowerMultiply(matrix, vec, n-2)

	return int(sumVector(resultVec))
}

// matrixPowerMultiply calculates:
//
// matrix^power * vec
//
// We do not need to calculate matrix^power alone first.
// We multiply the vector whenever the current power bit is 1.
func matrixPowerMultiply(matrix [][]int64, vec []int64, power int) []int64 {
	result := make([]int64, len(vec))
	copy(result, vec)

	// We need identity behavior for result vector.
	// Easier approach:
	// Use applied = false.
	// When first active bit appears, result = matrix * vec.
	applied := false

	for power > 0 {
		if power%2 == 1 {
			if !applied {
				result = multiplyMatrixVector(matrix, vec)
				applied = true
			} else {
				result = multiplyMatrixVector(matrix, result)
			}
		}

		matrix = multiplyMatrixMatrix(matrix, matrix)
		power /= 2
	}

	// If power was 0 originally, return original vec.
	// In this problem we call it only when n > 2,
	// so power >= 1, but this keeps the helper safe.
	if !applied {
		return vec
	}

	return result
}

// multiplyMatrixVector returns matrix * vec.
func multiplyMatrixVector(matrix [][]int64, vec []int64) []int64 {
	n := len(vec)
	res := make([]int64, n)

	for i := 0; i < n; i++ {
		var sum int64 = 0

		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 || vec[j] == 0 {
				continue
			}

			sum = (sum + matrix[i][j]*vec[j]) % MOD
		}

		res[i] = sum
	}

	return res
}

// multiplyMatrixMatrix returns a * b.
func multiplyMatrixMatrix(a [][]int64, b [][]int64) [][]int64 {
	n := len(a)
	res := make([][]int64, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int64, n)
	}

	for i := 0; i < n; i++ {
		for k := 0; k < n; k++ {
			if a[i][k] == 0 {
				continue
			}

			for j := 0; j < n; j++ {
				if b[k][j] == 0 {
					continue
				}

				res[i][j] = (res[i][j] + a[i][k]*b[k][j]) % MOD
			}
		}
	}

	return res
}

func sumVector(vec []int64) int64 {
	var sum int64 = 0

	for _, val := range vec {
		sum = (sum + val) % MOD
	}

	return sum
}

func main() {
	fmt.Println(zigZagArrays(3, 1, 3))
}
