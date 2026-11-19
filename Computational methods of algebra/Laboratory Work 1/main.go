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

func GaussDivSchemeMethod(A_original [][]float64, f_original []float64) ([]float64, float64) {
	n := len(A_original)
	A := CopyMatrix(A_original)
	f := CopyFreeValuesVector(f_original)

	B := make([][]float64, n)
	for i := 0; i < n; i++ {
		B[i] = make([]float64, n)
	}
	g := make([]float64, n)
	detA := 1.0

	for k := 0; k < n; k++ {
		leader := A[k][k]
		detA *= leader

		for j := k + 1; j < n; j++ {
			B[k][j] = A[k][j] / leader
		}
		g[k] = f[k] / leader

		for i := k + 1; i < n; i++ {
			for j := k + 1; j < n; j++ {
				A[i][j] -= A[i][k] * B[k][j]
			}
			f[i] -= A[i][k] * g[k]
		}
	}

	//PrintMatrix(B)

	x := make([]float64, n)
	x[n-1] = g[n-1]

	for i := n - 2; i >= 0; i-- {
		sum := 0.0
		for j := i + 1; j < n; j++ {
			sum += B[i][j] * x[j]
		}
		x[i] = g[i] - sum
	}

	return x, detA
}

func InverseMatrix(A_original [][]float64) [][]float64 {
	n := len(A_original)
	AInv := make([][]float64, n)
	for i := 0; i < n; i++ {
		AInv[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		e := make([]float64, n)
		e[i] = 1.0

		x, _ := GaussDivSchemeMethod(A_original, e)

		for row := 0; row < n; row++ {
			AInv[row][i] = x[row]
		}
	}

	return AInv
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

func ResidualMatrix(A [][]float64, InvA [][]float64) [][]float64 {
	n := len(A)
	R := make([][]float64, n)

	for i := 0; i < n; i++ {
		R[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			sum := 0.0
			for k := 0; k < n; k++ {
				sum += InvA[i][k] * A[k][j]
			}
			if i == j {
				R[i][j] = sum - 1.0
			} else {
				R[i][j] = sum
			}
		}
	}

	return R
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

func PrintExpMatrix(A [][]float64) {
	for _, row := range A {
		for _, val := range row {
			fmt.Printf("%13.5e ", val)
		}
		fmt.Println()
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
	_, A, v, _ := ReadMatrixFromTxt("../data.txt")
	//PrintMatrix(A)
	//PrintVector(v)
	x, detA := GaussDivSchemeMethod(A, v)
	fmt.Printf("%5f\n", x)
	fmt.Println("Determinant:", detA)

	resA := ResidualMatrix(A, InverseMatrix(A))
	fmt.Println("\nResidual Matrix:")
	PrintExpMatrix(resA)

	resV := ResidualVector(A, x, v)
	fmt.Println("\nResidual Vector:")
	PrintExpVector(resV)

	fmt.Println("\nChebyshev Norm of Residual Matrix:", ChebyshevNormMatrix(resA))
	fmt.Println("Chebyshev Norm of Residual Vector:", ChebyshevNormVector(resV))
	fmt.Println("Comparison of Chebyshev Norms:", ChebyshevNormMatrix(resA) < ChebyshevNormVector(resV))

	normA := ChebyshevNormMatrix(A)
	normInvA := ChebyshevNormMatrix(InverseMatrix(A))
	condA := normA * normInvA
	fmt.Printf("\nThe number of causality: %.5e\n", condA)

	//PrintMatrix(A)
}
