package main

import (
	"fmt"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Interval struct {
	Start float64 
	End float64
	Step float64
}

type Point struct {
	X float64
	Y float64
}

func main() {
	points := generatePoints(
		func(x float64) float64 {
			return math.Pow(1 + x, 1/x)
		},
		Interval{Start: 0, End: 1, Step: 0.1},
	)

	fmt.Printf("Points: %v", points)
}
func plotsTester() {
	generatePlot(trig, Interval{Start: -10, End: 10, Step: 0.01}, "Plot of sin(x)/x", "sinx-over-x.png")
	generatePlot(
		func(x float64) float64 { return 1 + x},
		Interval{Start: 1, End: 5, Step: 0.5},
		"Plot of 1+x",
		"increment.png",
	)
	generatePlot(
		func(x float64) float64 { return math.Exp(-x*x/2) * math.Cos(5 * x)},
		Interval{Start: 1, End: 5, Step: 0.01},
		"Plot of some complicated function",
		"complicated.png",
	)
}

// Given a function and an interval on that function's domain
// generate a set of points on the function
func generatePoints(fn func(float64) float64, interval Interval) []Point {
	points := make([]Point, 0)
	for x := interval.Start; x < interval.End; x += interval.Step {
		points = append(points, Point{X: x, Y: fn(x)})
	}
	return points
}

// plot any arbitrary function
func generatePlot(fn func(float64) float64, interval Interval, description string, filename string) {
	// Generate points for plotting
	points := make(plotter.XYs, 0)

	for x := interval.Start; x < interval.End; x += interval.Step {
		y := fn(x)
		fmt.Printf("x: %.2f, y: %.2f\n", x, y)
		points = append(points, plotter.XY{X: x, Y: y})
	}

	// fmt.Printf("Generated points: %v", points)

	// Create a new plot 
	plot := plot.New()
	plot.Title.Text = description
	plot.X.Label.Text = "x"
	plot.Y.Label.Text = "y"

	// Add a line plotter
	line, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	plot.Add(line)

	// Save the plot to a PNG file
	if err := plot.Save(4 * vg.Inch, 4 * vg.Inch, filename); err != nil {
		panic(err)
	}
}

func trig(x float64) float64 {
	if x == 0 {
		return 1
	}
	return math.Sin(x) / x
}