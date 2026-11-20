package main

import (
	"fmt"
	"math"
	"os"
)

func ReadMatrixFromTxt(filename string) (int, [][]float64, []float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, nil, err
	}
	defer file.Close()

	var n int
	_, err = fmt.Fscanf(file, "%d", &n)
	if err != nil {
		return 0, nil, nil, err
	}

	A := make([][]float64, n)
	f := make([]float64, n)

	for i := 0; i < n; i++ {
		A[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(file, "%f", &A[i][j])
		}
		fmt.Fscanf(file, "%f", &f[i])
	}
	return n, A, f, nil
}

func CopyMatrix(A [][]float64) [][]float64 {
	n := len(A)
	copied := make([][]float64, n)
	for i := 0; i < n; i++ {
		copied[i] = make([]float64, n)
		copy(copied[i], A[i])
	}
	return copied
}

func CopyFreeValuesVector(f []float64) []float64 {
	copied := make([]float64, len(f))
	copy(copied, f)
	return copied
}

func DownRelaxationMethod(n int, origA [][]float64, origf []float64, eps float64, w float64) ([]float64, int) {
	A := CopyMatrix(origA)
	f := CopyFreeValuesVector(origf)
	x := CopyFreeValuesVector(origf)
	k := 0

	for {
		k++
		xIter := make([]float64, n)
		for i := 0; i < n; i++ {
			xIter[i] = (1.0 - w) * x[i]
			for j := 0; j < i; j++ {
				xIter[i] -= w * A[i][j] / A[i][i] * xIter[j]
			}
			for j := i + 1; j < n; j++ {
				xIter[i] -= w * A[i][j] / A[i][i] * x[j]
			}
			xIter[i] += w * f[i] / A[i][i]
		}
		diff := make([]float64, n)
		for i := 0; i < n; i++ {
			diff[i] = xIter[i] - x[i]
		}
		normDiff := ChebyshevNormVector(diff)
		copy(x, xIter)
		if normDiff < eps {
			break
		}
	}
	return x, k
}

func ResidualVector(A [][]float64, x []float64, f []float64) []float64 {
	n := len(A)
	r := make([]float64, n)

	for i := 0; i < n; i++ {
		ax := 0.0
		for j := 0; j < n; j++ {
			ax += A[i][j] * x[j]
		}
		r[i] = ax - f[i]
	}
	return r
}

func ChebyshevNormVector(v []float64) float64 {
	max := 0.0
	for _, value := range v {
		if math.Abs(value) > max {
			max = math.Abs(value)
		}
	}
	return max
}

func ChebyshevNormMatrix(A [][]float64) float64 {
	max := 0.0

	for i := 0; i < len(A); i++ {
		rowSum := 0.0

		for j := 0; j < len(A[i]); j++ {
			rowSum += math.Abs(A[i][j])
		}
		if rowSum > max {
			max = rowSum
		}
	}
	return max
}

func PrintExpVector(v []float64) {
	for _, value := range v {
		fmt.Printf("%12.5e\n", value)
	}
}

func PrintMatrix(A [][]float64) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			fmt.Printf("%10.6f ", A[i][j])
		}
		fmt.Println()
	}
}

func DiagonalDominanceCheck(A [][]float64, n int) bool {
	for i := 0; i < n; i++ {
		rowSum := 0.0
		for j := 0; j < n; j++ {
			if i != j {
				rowSum += math.Abs(A[i][j])
			}
		}
		if math.Abs(A[i][i]) <= rowSum {
			fmt.Println(A[i][i], " <= ", rowSum)
			return false
		}
		fmt.Println(A[i][i], " > ", rowSum)
	}
	return true
}

func main() {
	eps := 1e-5
	w := 0.5
	n, A, f, _ := ReadMatrixFromTxt("../data.txt")
	fmt.Println("Diagonal dominance check:", DiagonalDominanceCheck(A, n))
	x, k := DownRelaxationMethod(n, A, f, eps, w)
	fmt.Printf("%1.5f\n ", x)
	fmt.Printf("Iterations: %d\n", k)
	resV := ResidualVector(A, x, f)
	fmt.Printf("Residual vector:\n")
	PrintExpVector(resV)
	fmt.Printf("Chebyshev norm of residual vector: %1.5e\n", ChebyshevNormVector(resV))
}
