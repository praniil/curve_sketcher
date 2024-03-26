package plotter

import (
	"fmt"
	"math"

	"github.com/Arafatk/glot"
)

func Plot() {
	dimension := 2
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimension, persist, debug)
	plot.SetTitle("Pranil's Plotter")
	fct := func(x float64) float64 { return (math.Sin(x)) }
	groupName := "Sine Curve"
	style := "lines"
	pointsX := make([]float64, 500)
	for i := 0.0; i <= 11.17011; i = i + 0.01 {
		pointsX = append(pointsX, i)
	}
	// pointsX := []float64{0, 0.5235988, 1.047198, 1.570796, 2.094395, 3.141593}
	plot.AddFunc2d(groupName, style, pointsX, fct)

	fileName := "plot.png"
	err := plot.SavePlot(fileName)
	if err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}
	fmt.Println("file saved as: ", fileName)
}
