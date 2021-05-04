package services

import ( 
	"github.com/google/go-github/v35/github"
	log "github.com/sirupsen/logrus"
)

func GenerateChartFromContribs(stats []*github.ContributorStats) (string, error) {
	for _, stat := range stats {
		log.Info("user: ", *stat.Author.Login)
		for _, week := range stat.Weeks {
			log.Info("week: ", week.String())
		}
	}

	return "info", nil
}