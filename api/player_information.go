package domilike

import (
	"context"
	"fmt"
	"log"

	playerinformation "github.com/comi91262/domilike/gen/player_information"
	"goa.design/goa/v3/security"
)

// player-information service example implementation.
// The example methods log the requests and return zero values.
type playerInformationsrvc struct {
	logger *log.Logger
}

// NewPlayerInformation returns the player-information service implementation.
func NewPlayerInformation(logger *log.Logger) playerinformation.Service {
	return &playerInformationsrvc{logger}
}

// JWTAuth implements the authorization logic for service "player-information"
// for the "jwt" security scheme.
func (s *playerInformationsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// Create implements create.
func (s *playerInformationsrvc) Create(ctx context.Context, p *playerinformation.CreatePayload) (res string, err error) {
	s.logger.Print("playerInformation.create")
	return
}

// Delete implements delete.
func (s *playerInformationsrvc) Delete(ctx context.Context, p *playerinformation.DeletePayload) (err error) {
	s.logger.Print("playerInformation.delete")
	return
}

// GetCoins implements get_coins.
func (s *playerInformationsrvc) GetCoins(ctx context.Context, p *playerinformation.GetCoinsPayload) (res int, err error) {
	s.logger.Print("playerInformation.get_coins")
	return
}

// GetVictoryPoints implements get_victory_points.
func (s *playerInformationsrvc) GetVictoryPoints(ctx context.Context, p *playerinformation.GetVictoryPointsPayload) (res int, err error) {
	s.logger.Print("playerInformation.get_victory_points")
	return
}

// GetDecks implements get_decks.
func (s *playerInformationsrvc) GetDecks(ctx context.Context, p *playerinformation.GetDecksPayload) (res int, err error) {
	s.logger.Print("playerInformation.get_decks")
	return
}

// GetDiscards implements get_discards.
func (s *playerInformationsrvc) GetDiscards(ctx context.Context, p *playerinformation.GetDiscardsPayload) (res int, err error) {
	s.logger.Print("playerInformation.get_discards")
	return
}

// GetHands implements get_hands.
func (s *playerInformationsrvc) GetHands(ctx context.Context, p *playerinformation.GetHandsPayload) (res int, err error) {
	s.logger.Print("playerInformation.get_hands")
	return
}

// GetPlayArea implements get_play_area.
func (s *playerInformationsrvc) GetPlayArea(ctx context.Context, p *playerinformation.GetPlayAreaPayload) (res int, err error) {
	s.logger.Print("playerInformation.get_play_area")
	return
}
