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

func GaussZeidelMethod(n int, origA [][]float64, origf []float64, eps float64) ([]float64, int) {
	A := CopyMatrix(origA)
	f := CopyFreeValuesVector(origf)
	B := make([][]float64, n)
	for i := 0; i < n; i++ {
		B[i] = make([]float64, n)
	}
	b := make([]float64, n)
	x := make([]float64, n)
	x = CopyFreeValuesVector(f)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				B[i][j] = 0
			} else {
				B[i][j] = -A[i][j] / A[i][i]
			}
		}

		b[i] = f[i] / A[i][i]
	}

	fmt.Println("Chebyshev norm of matrix B:", ChebyshevNormMatrix(B))
	//PrintMatrix(B)
	//PrintExpVector(b)

	k := 0
	for {
		k++
		xIter := make([]float64, n)
		for i := 0; i < n; i++ {
			for j := 0; j < i; j++ {
				xIter[i] += B[i][j] * xIter[j]
			}
			for j := i; j < n; j++ {
				xIter[i] += B[i][j] * x[j]
			}
			xIter[i] += b[i]
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

func main() {
	n, A, f, _ := ReadMatrixFromTxt("../data.txt")
	//PrintMatrix(A)
	//PrintExpVector(f)
	eps := 1e-5
	x, k := GaussZeidelMethod(n, A, f, eps)
	fmt.Printf("%1.5f\n ", x)
	fmt.Printf("Iterations: %d\n", k)
	resV := ResidualVector(A, x, f)
	fmt.Printf("Residual vector:\n")
	PrintExpVector(resV)
	fmt.Printf("Chebyshev norm of residual vector: %1.5e\n", ChebyshevNormVector(resV))
}
