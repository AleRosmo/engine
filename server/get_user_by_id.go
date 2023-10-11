package server

import (
	"github.com/AleRosmo/engine/models/db"
)

func (s *serverImpl) GetUserByID(userID uint) (*db.User, error) {
	return s.db.GetUserByID(userID)
}
