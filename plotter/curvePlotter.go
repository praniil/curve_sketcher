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
	// fct := func(x float64) float64 { return (math.Sin(x)) }
	groupName := "Sine Curve"
	style := "lines"
	pointsX := make([]float64, 0, 100) //type, length, capacit:= length 0 aaile ko so no elements there, capacity can have this much initial capacity
	for i := 0.0; i <= 11.17011; i = i+0.01 {
		pointsX = append(pointsX, i)
	}
	pointsY := make([]float64, 0, 100)
	for _, pointX := range pointsX {
		pointsY = append(pointsY, math.Sin(pointX))
	}
	points := [][]float64{pointsX, pointsY}
	plot.AddPointGroup(groupName, style, points)
	fmt.Println(pointsX)
	fmt.Println(pointsY)
	// pointsX := []float64{0, 0.5235988, 1.047198, 1.570796, 2.094395, 3.141593}
	// plot.AddFunc2d(groupName, style, pointsX, fct)

	fileName := "plot.png"
	err := plot.SavePlot(fileName)
	if err != nil {
		fmt.Println("Error saving plot:", err)
		return
	}
	fmt.Println("file saved as: ", fileName)
}
