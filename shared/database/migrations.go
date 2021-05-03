package database

import (
	"github.com/edwinvautier/gh-readme-contrib/api/models"
	log "github.com/sirupsen/logrus"
)

// Migrate executes migrations once the db is connected
func Migrate() {
	log.Info("Executing migrations...")
	Db.AutoMigrate(&models.Repository{})
}
