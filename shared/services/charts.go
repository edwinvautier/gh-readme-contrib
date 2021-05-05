package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v35/github"
)

func GenerateChartFromContribs(config ChartConfig) (string, error) {
	svg := `
	<svg width="440" height="270" xmlns="http://www.w3.org/2000/svg">
	<g>
	 <title>Commits chart</title>
	 <rect id="svg_3" height="270" width="440" fill="#`+ config.UI.BackgroundColor +`"/>
	 <rect y="40" x="10" id="svg_7" height="200" width="420" fill="#fff"/>
	 <text xml:space="preserve" text-anchor="start" font-family="sans-serif" font-size="12" stroke-width="0" id="svg_4" y="20" x="10" stroke="#000" fill="#`+ config.UI.TextColor+`">`+ config.Author + "/" +  config.Name +`</text>
	 <line id="svg_5" y2="240" x2="115" y1="40" x1="115" stroke="#AAA" fill="none"/>
	 <line id="svg_8" y2="240" x2="220" y1="40" x1="220" stroke="#AAA" fill="none"/>
	 <line id="svg_9" y2="240" x2="325" y1="40" x1="325" stroke="#AAA" fill="none"/>
	 `
	
	svg += `
	 <line id="svg_6" y2="290" x2="200" y1="120" x1="100" stroke="#000" fill="none"/>
	</g>
 </svg>`
	return svg, nil
}

type ChartConfig struct {
	MaxHeight,
	MaxValue 		uint
	UI 					UIConfig
	Author,
	Name				string
	WeeklyStats []*github.WeeklyCommitActivity
}

type UIConfig struct {
	BackgroundColor	string
	LineColor				string
	TextColor 			string
}

func InitChartConfig(c *gin.Context) ChartConfig {
	var config ChartConfig
  config.Name = c.Param("repository")
  config.Author = c.Param("author")

	// UI Config
	config.UI = UIConfig{
		BackgroundColor: "ddd",
		LineColor: "000",
		TextColor: "333",
	}
	config.MaxHeight = 190

	return config
}