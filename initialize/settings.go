package initialize

import (
	"github.com/AleRosmo/engine/models/events"
	"github.com/YalkChat/database"
)

func saveInitialSettings(dbCon database.DatabaseOperations) error {
	serverSettings := &events.ServerSettings{IsInitialized: true}
	if err := dbCon.SaveServerSettings(serverSettings); err != nil {
		return err
	}
	return nil
}
