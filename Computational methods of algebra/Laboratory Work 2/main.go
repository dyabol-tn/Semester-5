package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadDiagonalsFromTxt(filename string) ([][]float64, []float64, []float64, []float64, []float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	defer file.Close()

	var A [][]float64
	var f []float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		nCols := len(parts)
		rowA := make([]float64, nCols-1)
		for j := 0; j < nCols-1; j++ {
			val, err := strconv.ParseFloat(parts[j], 64)
			if err != nil {
				return nil, nil, nil, nil, nil, err
			}
			rowA[j] = val
		}
		fVal, err := strconv.ParseFloat(parts[nCols-1], 64)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}

		A = append(A, rowA)
		f = append(f, fVal)
	}

	n := len(A)
	a := make([]float64, n)
	c := make([]float64, n)
	b := make([]float64, n)

	for i := 0; i < n; i++ {
		c[i] = A[i][i]

		if i > 0 {
			a[i] = -A[i][i-1]
		} else {
			a[i] = 0.0
		}

		if i < n-1 {
			b[i] = -A[i][i+1]
		} else {
			b[i] = 0.0
		}
	}

	return A, a, c, b, f, nil
}

func MethodVstrProgonki(origA, origB, origC, origF []float64) []float64 {
	n := len(origC) - 1
	a := CopyFreeValuesVector(origA)
	b := CopyFreeValuesVector(origB)
	c := CopyFreeValuesVector(origC)
	f := CopyFreeValuesVector(origF)
	m := n / 2

	alpha := make([]float64, n+1)
	beta := make([]float64, n+1)
	ksi := make([]float64, n+1)
	eta := make([]float64, n+1)
	y := make([]float64, n+1)

	alpha[1] = b[0] / c[0]
	beta[1] = f[0] / c[0]
	for i := 1; i <= m-1; i++ {
		alpha[i+1] = b[i] / (c[i] - a[i]*alpha[i])
		beta[i+1] = (f[i] + a[i]*beta[i]) / (c[i] - a[i]*alpha[i])
	}

	ksi[n] = a[n] / c[n]
	eta[n] = f[n] / c[n]
	for i := n - 1; i >= m+1; i-- {
		ksi[i] = a[i] / (c[i] - b[i]*ksi[i+1])
	}

	for i := n - 1; i >= m; i-- {
		eta[i] = (f[i] + b[i]*eta[i+1]) / (c[i] - b[i]*ksi[i+1])
	}

	//y[m] = (beta[m] + alpha[m]*eta[m]) / (1.0 - ksi[m]*alpha[m])
	y[m] = eta[m] + ksi[m]*beta[m]/(1.0-ksi[m]*alpha[m])

	for i := m - 1; i >= 0; i-- {
		y[i] = alpha[i+1]*y[i+1] + beta[i+1]
	}

	for i := m; i <= n-1; i++ {
		y[i+1] = ksi[i+1]*y[i] + eta[i+1]
	}

	return y
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
	filename := "../dataLab2.txt"
	A, a, c, b, f, err := ReadDiagonalsFromTxt(filename)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	/*fmt.Println("a:", a)
	fmt.Println("c:", c)
	fmt.Println("b:", b)
	fmt.Println("f:", f)*/

	y := MethodVstrProgonki(a, b, c, f)
	fmt.Printf("y: %12.5f\n", y)

	r := ResidualVector(A, y, f)
	fmt.Println("Residual Vector:")
	PrintExpVector(r)

	fmt.Printf("\nChebyshev Norm of Residual Vector: %12.5e\n", ChebyshevNormVector(r))
}
