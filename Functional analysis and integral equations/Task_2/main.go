package main

import (
	"fmt"
	"math"
)

func main() {
	epsilon := 1e-4

	A := [][]float64{
		{6, 2, 7, -3},
		{10, 5, -4, 8},
		{15, -7, 2, 6},
		{-8, 4, 5, 1},
	}
	B := []float64{20, 70, 645, -210}
	n := len(B)

	At := transpose(A)
	AtA := multiplyMatrices(At, A)
	AtB := matVecMul(At, B)
	lambdaMax := powerIteration(AtA, 1000)

	C := make([][]float64, n)
	for i := range C {
		C[i] = make([]float64, n)
		for j := range n {
			delta := 0.0
			if i == j {
				delta = 1.0
			}
			C[i][j] = delta - AtA[i][j]/lambdaMax
		}
	}

	D := make([]float64, n)
	for i := range n {
		D[i] = AtB[i] / lambdaMax
	}

	alpha := 0.0
	for i := 0; i < n; i++ {
		rowSum := 0.0
		for j := 0; j < n; j++ {
			rowSum += math.Abs(C[i][j])
		}
		if rowSum > alpha {
			alpha = rowSum
		}
	}
	fmt.Println("Коэффициент сжатия alpha (норма матрицы C):", alpha)

	X0 := make([]float64, n)
	X1 := make([]float64, n)
	copy(X1, D)

	diff01Vec := make([]float64, n)
	for i := 0; i < n; i++ {
		diff01Vec[i] = X1[i] - X0[i]
	}
	diff01 := normVector(diff01Vec)

	var nApr float64
	if alpha < 1 {
		nApr = math.Floor(math.Log(epsilon*(1-alpha)/diff01)/math.Log(alpha)) + 1
		fmt.Println("Априорная оценка количества итераций (n_apr):", nApr)
	} else {
		fmt.Println("alpha >= 1 по кубической норме, достаточные условия не выполнены")
	}

	xStart := make([]float64, n)
	copy(xStart, X0)
	xNext := make([]float64, n)
	nIter := 0

	for {
		nIter++
		for i := 0; i < n; i++ {
			sum := D[i]
			for j := 0; j < n; j++ {
				sum += C[i][j] * xStart[j]
			}
			xNext[i] = sum
		}

		diffVectors := make([]float64, n)
		for i := 0; i < n; i++ {
			diffVectors[i] = xNext[i] - xStart[i]
		}

		if normVector(diffVectors) <= epsilon || nIter > 10000 {
			break
		}
		copy(xStart, xNext)
	}

	fmt.Println("Результаты:")
	fmt.Println("Номер последней итерации (фактический):", nIter)
	fmt.Println("Вектор решения X*:")
	for i, val := range xNext {
		fmt.Printf("x[%d] = %.6f\n", i+1, val)
	}
}

func multiplyMatrices(A, B [][]float64) [][]float64 {
	rA, cA := len(A), len(A[0])
	cB := len(B[0])
	res := make([][]float64, rA)
	for i := range res {
		res[i] = make([]float64, cB)
		for j := 0; j < cB; j++ {
			sum := 0.0
			for k := 0; k < cA; k++ {
				sum += A[i][k] * B[k][j]
			}
			res[i][j] = sum
		}
	}
	return res
}

func matVecMul(A [][]float64, v []float64) []float64 {
	res := make([]float64, len(A))
	for i := range A {
		sum := 0.0
		for j := range v {
			sum += A[i][j] * v[j]
		}
		res[i] = sum
	}
	return res
}

func transpose(A [][]float64) [][]float64 {
	rows, cols := len(A), len(A[0])
	T := make([][]float64, cols)
	for i := range T {
		T[i] = make([]float64, rows)
		for j := range rows {
			T[i][j] = A[j][i]
		}
	}
	return T
}

func normVector(v []float64) float64 {
	maxVal := 0.0
	for _, val := range v {
		if math.Abs(val) > maxVal {
			maxVal = math.Abs(val)
		}
	}
	return maxVal
}

func powerIteration(A [][]float64, maxIter int) float64 {
	n := len(A)
	b := make([]float64, n)
	for i := range b {
		b[i] = 1.0
	}
	lambda := 0.0
	for iter := 0; iter < maxIter; iter++ {
		nextB := matVecMul(A, b)
		norm := 0.0
		for _, val := range nextB {
			norm += val * val
		}
		norm = math.Sqrt(norm)
		for i := range nextB {
			b[i] = nextB[i] / norm
		}

		num, den := 0.0, 0.0
		Ab := matVecMul(A, b)
		for i := range b {
			num += Ab[i] * b[i]
			den += b[i] * b[i]
		}
		lambda = num / den
	}
	return lambda
}
