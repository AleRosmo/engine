package server

import (
	"net/http"
	"yalk/config"

	"github.com/AleRosmo/engine/client"
	"github.com/AleRosmo/engine/models/db"
	"github.com/AleRosmo/engine/models/events"

	"nhooyr.io/websocket"
)

type Server interface {
	RegisterClient(client.Client) error
	UnregisterClient(client.Client) error
	SendChat(*events.BaseEvent, uint) error
	SendAll(*events.BaseEvent) error
	GetClientByID(uint) (client.Client, error)
	HandleEvent(*events.BaseEventWithMetadata) error
	GetUserByID(uint) (*db.User, error)
	GetUserByUsername(username string) (user *db.User, err error)
	UpgradeHttpRequest(http.ResponseWriter, *http.Request, *config.Config) (*websocket.Conn, error)
	AuthenticateUser(loginUser *db.User) (userID uint, err error)
	CreateUser(user *db.User) (userID uint, err error)
}

// Additional type definitions for Client, Message, etc.
// TODO: IMPORTANT. Server is the only things to be used, it's db
// .. methods for now are private and server provides an abstraction
