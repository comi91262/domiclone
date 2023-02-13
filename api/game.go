package domilike

import (
	"context"
	"log"

	game "github.com/comi91262/domilike/gen/game"
)

// game service example implementation.
// The example methods log the requests and return zero values.
type gamesrvc struct {
	logger *log.Logger
}

// NewGame returns the game service implementation.
func NewGame(logger *log.Logger) game.Service {
	return &gamesrvc{logger}
}

// Get implements get.
func (s *gamesrvc) Get(ctx context.Context, p *game.GetPayload) (res string, err error) {
	s.logger.Print("game.get")
	return
}

// Create implements create.
func (s *gamesrvc) Create(ctx context.Context) (res int, err error) {
	s.logger.Print("game.create")
	return
}

// Delete implements delete.
func (s *gamesrvc) Delete(ctx context.Context, p *game.DeletePayload) (err error) {
	s.logger.Print("game.delete")
	return
}

// GetSupplies implements get_supplies.
func (s *gamesrvc) GetSupplies(ctx context.Context, p *game.GetSuppliesPayload) (res string, err error) {
	s.logger.Print("game.get_supplies")
	return
}

// GetTrashes implements get_trashes.
func (s *gamesrvc) GetTrashes(ctx context.Context, p *game.GetTrashesPayload) (res string, err error) {
	s.logger.Print("game.get_trashes")
	return
}
