package domilike

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	card "github.com/comi91262/domilike/gen/card"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

// card service example implementation.
// The example methods log the requests and return zero values.
type cardsrvc struct {
	logger *log.Logger
}

// NewCard returns the card service implementation.
func NewCard(logger *log.Logger) card.Service {
	return &cardsrvc{logger}
}

// Get implements get.
func (s *cardsrvc) Get(ctx context.Context, p *card.GetPayload) (res string, err error) {
	dialect := goqu.Dialect("sqlite3")

	d, err := sql.Open("sqlite3", "../database/database.sqlite")
	if err != nil {
		return "", err
	}
	defer d.Close()

	db := dialect.DB(d)

	count, err := db.From("cards").Count()
	if err != nil {
		return "", err
	}
	fmt.Printf("card count = %d", count)

	return fmt.Sprint(count), nil
}
