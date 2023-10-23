package initialize

import (
	"github.com/YalkChat/database"

	"github.com/AleRosmo/engine/models/db"
	"github.com/AleRosmo/engine/models/events"
)

func createChannelType(dbConn database.DatabaseOperations) (*db.ChatType, error) {
	chatType := &events.ChatType{Name: "channel"}
	DbChatType, err := dbConn.NewChatType(chatType)
	if err != nil {
		return nil, err
	}
	return DbChatType, nil
}
