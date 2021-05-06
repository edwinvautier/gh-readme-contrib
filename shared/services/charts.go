package services

import (
	"fmt"
	"math"

	"github.com/edwinvautier/gh-readme-contrib/api/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GenerateChartFromContribs(config ChartConfig) (string, error) {
	var max uint
	for _, week := range config.WeeklyStats {
		if uint(week.Total) > max {
			max = uint(week.Total)
		}
	}
	config.MaxValue = max
	svg := `
	<svg width="440" height="270" xmlns="http://www.w3.org/2000/svg">
	path{
		fill : url(#gradient);
	}
	<defs>
    <linearGradient id="gradient" x1="50%" y1="0%" x2="50%" y2="100%">
      <stop offset="0%"   stop-color="#1645ff"/>
      <stop offset="50%"   stop-color="#1645ff"/>
      <stop offset="150%" stop-color="#7b99fc"/>
    </linearGradient>
  </defs>
	<g>
	<title>Commits chart</title>
	<rect id="svg_3" height="270" width="440" fill="#` + config.UI.BackgroundColor + `"/>
	<rect y="40" x="10" id="svg_7" height="200" width="420" fill="#fff"/>
	<text xml:space="preserve" text-anchor="start" font-family="sans-serif" font-size="12" stroke-width="0" id="svg_4" y="20" x="10" stroke="#000" fill="#` + config.UI.TextColor + `">` + config.Author + "/" + config.Name + `</text>
	<line id="svg_5" y2="240" x2="115" y1="40" x1="115" stroke="#AAA" fill="none"/>
	<line id="svg_8" y2="240" x2="220" y1="40" x1="220" stroke="#AAA" fill="none"/>
	<line id="svg_9" y2="240" x2="325" y1="40" x1="325" stroke="#AAA" fill="none"/>
	`
	points := []Point{
		{
			X: 10,
			Y: 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, uint(config.WeeklyStats[len(config.WeeklyStats) - 5].Total)),
		},
		{
			X: 115,
			Y: 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, uint(config.WeeklyStats[len(config.WeeklyStats) - 4].Total)),
		},
		{
			X: 220,
			Y: 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, uint(config.WeeklyStats[len(config.WeeklyStats) - 3].Total)),
		},
		{
			X: 325,
			Y: 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, uint(config.WeeklyStats[len(config.WeeklyStats) - 2].Total)),
		},
		{
			X: 430,
			Y: 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, uint(config.WeeklyStats[len(config.WeeklyStats) - 1].Total)),
		},
	}
	path := renderCurve(points)
	svg += path
	svg += `
	</g>
 </svg>`

	return svg, nil
}

func calcOffsetBottom(maxHeight, maxValue, value uint) uint {
	return (value * maxHeight) / maxValue
}

type ChartConfig struct {
	MaxHeight,
	MaxValue uint
	UI UIConfig
	Author,
	Name string
	WeeklyStats []models.Week
}

type UIConfig struct {
	BackgroundColor string
	LineColor       string
	TextColor       string
}

func InitChartConfig(c *gin.Context) ChartConfig {
	var config ChartConfig
	config.Name = c.Param("repository")
	config.Author = c.Param("author")

	// UI Config
	config.UI = UIConfig{
		BackgroundColor: "ddd",
		LineColor:       "f0f",
		TextColor:       "333",
	}
	config.MaxHeight = 190

	return config
}

func renderCurve(points []Point) string {
	path := "<path fill=\"url(#gradient)\" stroke=\"none\" d=\""

	for i := 0; i < len(points); i++ {
		point := points[i]
		if i == 0 {
			path += fmt.Sprintf("M 10 240 L %d %d", point.X, point.Y)
		} else {
			path+= bezier(point, i, points)
		}
	}
	path += "\"/>"

	return path
}

func bezier(point Point, i int, points []Point) string {
	previousIndex := i-2
	if i - 2 < 0 {
		previousIndex = i
	}
	nextIndex := i+1
	if i + 1 == len(points) {
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
	logrus.Info("points ", point1, point2)
	lengthX := float64(point1.X) - float64(point2.X) 
	lengthY := float64(point1.Y) - float64(point2.Y)
	return Line {
		Length: math.Sqrt(math.Pow(lengthX, 2) + math.Pow(lengthY, 2)),
		Angle: math.Atan2(lengthY, lengthX),
	}
}

func controlPoint(current, previous, next Point, reverse bool) Point {	
	smoothing := 0.2
	o := line(previous, next)
	logrus.Info("line ", o.Angle, o.Length)
	var angle float64
	if reverse {
		angle = math.Pi
	} else {
		angle = 0
	}
	length := o.Length * smoothing
	x := float64(current.X) + math.Cos(angle) * length
	y := float64(current.Y) + math.Sin(angle) * length
	logrus.Info(math.Round(x))
	return Point {
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
	Angle float64
}