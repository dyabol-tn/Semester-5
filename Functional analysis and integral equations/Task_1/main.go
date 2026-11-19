package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func g(x float64) float64 {
	return x*math.Sqrt(5) + math.Sqrt(1+x*x) - math.Pow(math.Sin(x), 2) + 4
}

func f(x float64) float64 {
	return (math.Pow(math.Sin(x), 2) - math.Sqrt(1+x*x) - 4) / math.Sqrt(5)
}

func main() {
	epsilon := 1e-4
	alpha := 2.0 / math.Sqrt(5)
	x0 := 0.0
	x1 := f(x0)
	diff01 := math.Abs(x0 - x1)

	nApr := math.Floor(math.Log(epsilon*(1-alpha)/diff01)/math.Log(alpha)) + 1

	fmt.Println("Коэффициент сжатия:", alpha)
	fmt.Println("Априорная оценка количества итераций:", nApr)

	xStart := x0
	xNext := f(xStart)
	n := 0

	for math.Abs(xNext-xStart) > epsilon*(1-alpha)/alpha {
		n++
		xStart = xNext
		xNext = f(xStart)
	}
	fmt.Println("Номер последней итерации:", n)
	fmt.Println("Найденный корень x*:", xNext)
	fmt.Println("Проверка g(x*):", g(xNext))

	p := plot.New()
	p.Title.Text = "График функции g(x) и найденный корень"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "g(x)"

	numPoints := 200
	pts := make(plotter.XYs, numPoints)
	xmin, xmax := -3.5, 1.0

	for i := 0; i < numPoints; i++ {
		x := xmin + float64(i)*(xmax-xmin)/float64(numPoints-1)
		pts[i].X = x
		pts[i].Y = g(x)
	}

	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	line.Color = color.RGBA{R: 0, G: 102, B: 204, A: 255}
	line.Width = vg.Points(2)

	rootPt := make(plotter.XYs, 1)
	rootPt[0].X = xNext
	rootPt[0].Y = 0.0

	scatter, err := plotter.NewScatter(rootPt)
	if err != nil {
		panic(err)
	}
	scatter.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	scatter.Radius = vg.Points(3)

	grid := plotter.NewGrid()
	p.Add(line, scatter, grid)

	p.Legend.Add("g(x)", line)
	p.Legend.Add("Найденный корень", scatter)

	p.Y.Min = -1.0
	p.Y.Max = 7.5
	if err := p.Save(7*vg.Inch, 4.5*vg.Inch, "plot.png"); err != nil {
		panic(err)
	}
}
