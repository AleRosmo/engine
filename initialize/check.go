package initialize

import (
	"github.com/YalkChat/database"
)

func checkIsInitialized(dbConn database.DatabaseOperations) bool {
	isInitialized, err := dbConn.IsServerInitialized()
	if err != nil {
		return false
	}
	return isInitialized
}
