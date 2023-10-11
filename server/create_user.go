package server

import "github.com/AleRosmo/engine/models/db"

func (s *serverImpl) CreateUser(newUser *db.User) (uint, error) {
	newDbUser, err := s.db.NewUser(newUser)
	if err != nil {
		return 0, err
	}
	return newDbUser.ID, nil
}
