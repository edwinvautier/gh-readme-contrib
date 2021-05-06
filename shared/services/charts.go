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
	<g>
	<title>Commits chart</title>
	<rect id="svg_3" height="270" width="440" fill="#` + config.UI.BackgroundColor + `"/>
	<rect y="40" x="10" id="svg_7" height="200" width="420" fill="#fff"/>
	<text xml:space="preserve" text-anchor="start" font-family="sans-serif" font-size="12" stroke-width="0" id="svg_4" y="20" x="10" stroke="#000" fill="#` + config.UI.TextColor + `">` + config.Author + "/" + config.Name + `</text>
	<line id="svg_5" y2="240" x2="115" y1="40" x1="115" stroke="#AAA" fill="none"/>
	<line id="svg_8" y2="240" x2="220" y1="40" x1="220" stroke="#AAA" fill="none"/>
	<line id="svg_9" y2="240" x2="325" y1="40" x1="325" stroke="#AAA" fill="none"/>
	`
	for i := len(config.WeeklyStats) - 4; i < len(config.WeeklyStats); i++ {
		currWeekValue := uint(config.WeeklyStats[i].Total)
		prevWeekValue := uint(config.WeeklyStats[i-1].Total)
		leftOffset := 10 + (i-len(config.WeeklyStats)+4)*105
		prevHeight := 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, prevWeekValue)
		currHeight := 240 - calcOffsetBottom(config.MaxHeight, config.MaxValue, currWeekValue)

		line := fmt.Sprintf("<line y2=\"%d\" x2=\"%d\" y1=\"%d\" x1=\"%d\" stroke=\"#%s\"/>", currHeight, leftOffset+105, prevHeight, leftOffset, config.UI.LineColor)
		svg += line
	}
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
