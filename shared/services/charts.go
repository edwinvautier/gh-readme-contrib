package services

import (
	"fmt"

	"github.com/edwinvautier/gh-readme-contrib/api/models"
	"github.com/gin-gonic/gin"
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
      <stop offset="0%"   stop-color="#` + config.UI.MainColor + `"/>
      <stop offset="50%"   stop-color="#` + config.UI.MainColor + `"/>
      <stop offset="150%" stop-color="#` + config.UI.BackgroundColor + `"/>
    </linearGradient>
  </defs>
	<g>
	<title>Commits chart</title>
	<rect rx="15" id="svg_3" height="270" width="440" fill="#` + config.UI.BackgroundColor + `"/>
	<rect y="40" x="10" id="svg_7" height="200" width="420" fill="#` + config.UI.BackgroundColor + `"/>
	<text font-weight="bold" xml:space="preserve" text-anchor="start" font-family="sans-serif" font-size="15" stroke-width="0" id="svg_4" y="27" x="10" stroke="#000" fill="#` + config.UI.TextColor + `">Weekly activity</text>
	<text font-weight="bold" xml:space="preserve" text-anchor="end" font-family="sans-serif" font-size="15" stroke-width="0" id="svg_10" y="27" x="430" stroke="#000" fill="#` + config.UI.TextColor + `">` + fmt.Sprint(config.WeeklyStats[len(config.WeeklyStats)-1].Total) + ` commits this week</text>
	<path class="gridPath" d="M 115, 239 L 115 41" stroke="#` + config.UI.MainColor + `" fill="none" />
	<text xml:space="preserve" text-anchor="middle" font-family="sans-serif" font-size="10" stroke-width="0" id="svg_4" y="260" x="115" stroke="#000" fill="#` + config.UI.TextColor + `">` + config.WeeklyStats[len(config.WeeklyStats)-4].Date.Format("January 2") + `</text>
	
	<path class="gridPath" d="M 220, 239 L 220 41" stroke="#` + config.UI.MainColor + `" fill="none" />
	<text xml:space="preserve" text-anchor="middle" font-family="sans-serif" font-size="10" stroke-width="0" id="svg_4" y="260" x="220" stroke="#000" fill="#` + config.UI.TextColor + `">` + config.WeeklyStats[len(config.WeeklyStats)-3].Date.Format("January 2") + `</text>
	
	<path class="gridPath" d="M 325, 239 L 325 41" stroke="#` + config.UI.MainColor + `" fill="none" />
	<text xml:space="preserve" text-anchor="middle" font-family="sans-serif" font-size="10" stroke-width="0" id="svg_4" y="260" x="325" stroke="#000" fill="#` + config.UI.TextColor + `">` + config.WeeklyStats[len(config.WeeklyStats)-2].Date.Format("January 2") + `</text>
	`
	points := []Point{
		{
			X: 10,
			Y: 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, uint(config.WeeklyStats[len(config.WeeklyStats)-5].Total)),
		},
		{
			X: 115,
			Y: 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, uint(config.WeeklyStats[len(config.WeeklyStats)-4].Total)),
		},
		{
			X: 220,
			Y: 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, uint(config.WeeklyStats[len(config.WeeklyStats)-3].Total)),
		},
		{
			X: 325,
			Y: 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, uint(config.WeeklyStats[len(config.WeeklyStats)-2].Total)),
		},
		{
			X: 430,
			Y: 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, uint(config.WeeklyStats[len(config.WeeklyStats)-1].Total)),
		},
	}
	path := renderCurve(points, config.UI)
	svg += path
	svg += `
	<path d="M 10 242 L 10 40 L 430 40 L 430 242" stroke-width="4px" stroke="#` + config.UI.BackgroundColor + `" fill="none"/>
	</g>
	<style>
	.path {
		stroke-dasharray: 1000;
		stroke-dashoffset: 1000;
		fill-opacity: 0;
		animation: dash 3s linear 0s forwards, opacity 1s linear 2s forwards;
	}
	.gridPath {
		stroke-width: .5;
		stroke-dasharray: 200;
		stroke-dashoffset: 200;
		animation: dash 1s linear .5s forwards, enlarge .5s linear 2s forwards;
	}
	@keyframes dash {
		to {
			stroke-dashoffset: 0;
		}
	}
	@keyframes enlarge {
		to {
			stroke-width: 1px;
		}
	}
	@keyframes opacity {
		to {
			fill-opacity: 1;
		}
	}
	</style>

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
	MainColor       string
	TextColor       string
}

func InitChartConfig(c *gin.Context) ChartConfig {
	var config ChartConfig
	config.Name = c.Param("repository")
	config.Author = c.Param("author")

	// UI Config
	config.UI = UIConfig{
		BackgroundColor: "141321",
		MainColor:       "ff7f00",
		TextColor:       "D83A7C",
	}
	if c.Query("bg") != "" {
		config.UI.BackgroundColor = c.Query("bg")
	}
	if c.Query("main") != "" {
		config.UI.MainColor = c.Query("main")
	}
	if c.Query("text") != "" {
		config.UI.TextColor = c.Query("text")
	}
	config.MaxHeight = 160

	return config
}
