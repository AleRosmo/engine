package server

import (
	"fmt"

	"github.com/AleRosmo/engine/client"
)

func (s *serverImpl) UnregisterClient(client client.Client) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if client exists
	if _, exists := s.clients[client.ID()]; !exists {
		return fmt.Errorf("client with ID %d not found", client.ID())
	}

	// Remove client from internal tracking
	delete(s.clients, client.ID())

	// Notify other components or clients as needed

	return nil
}
