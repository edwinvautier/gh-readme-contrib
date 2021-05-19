package services

import (
	"fmt"
	"math"
)

func renderCurve(points []Point, config UIConfig) string {
	path := `<path class="path" fill-opacity="0.7" fill="url(#gradient)" stroke-width="3px" stroke="#` + config.MainColor + `" d="`

	for i := 0; i < len(points); i++ {
		point := points[i]
		if i == 0 {
			path += fmt.Sprintf("M 10 240 L %d %d", point.X, point.Y)
		} else {
			path += bezier(point, i, points)
		}
	}
	path += " L 430 240\"/>"

	return path
}

func bezier(point Point, i int, points []Point) string {
	previousIndex := i - 2
	if i-2 < 0 {
		previousIndex = i
	}
	nextIndex := i + 1
	if i+1 == len(points) {
		nextIndex = i
	}

	next := points[nextIndex]
	previous := points[previousIndex]
	startCtlPt := controlPoint(points[i-1], previous, point, false)
	endCtlPt := controlPoint(point, points[i-1], next, true)

	str := fmt.Sprintf("C %d,%d %d,%d %d,%d", startCtlPt.X, startCtlPt.Y, endCtlPt.X, endCtlPt.Y, point.X, point.Y)
	return str
}

func line(point1, point2 Point) Line {
	lengthX := float64(point1.X) - float64(point2.X)
	lengthY := float64(point1.Y) - float64(point2.Y)
	return Line{
		Length: math.Sqrt(math.Pow(lengthX, 2) + math.Pow(lengthY, 2)),
		Angle:  math.Atan2(lengthY, lengthX),
	}
}

func controlPoint(current, previous, next Point, reverse bool) Point {
	smoothing := 0.2
	o := line(previous, next)
	var angle float64
	if reverse {
		angle = math.Pi
	} else {
		angle = 0
	}
	length := o.Length * smoothing
	x := float64(current.X) + math.Cos(angle)*length
	y := float64(current.Y) + math.Sin(angle)*length
	return Point{
		X: uint(math.Round(x)),
		Y: uint(math.Round(y)),
	}
}

type Point struct {
	X uint
	Y uint
}

type Line struct {
	Length float64
	Angle  float64
}

func generateRect(val uint, position uint, maxVal uint, color string) string {
	height := calcOffsetBottom(150, maxVal, val) + 50
	yPos := fmt.Sprint(270 - height)
	xPos := fmt.Sprint(70 + 100*position)
	return `<rect y="` + yPos + `" x="` + xPos + `" height="` + fmt.Sprint(height) + `" width="100" fill="#` + color + `" />`
}
