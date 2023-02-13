package domilike

import (
	"context"
	"log"

	user "github.com/comi91262/domilike/gen/user"
)

// user service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	logger *log.Logger
}

// NewUser returns the user service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersrvc{logger}
}

// Get implements get.
func (s *usersrvc) Get(ctx context.Context, p *user.GetPayload) (res string, err error) {
	s.logger.Print("user.get")
	return
}
