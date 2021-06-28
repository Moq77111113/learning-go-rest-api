package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"moq.com/test/cmd/models"
)

type RandomDb interface {
	Open() error
	Close() error
	Create (p *models.Post) error
	Get () ([]*models.Post, error)
	
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error {
	pg, err := sqlx.Open("postgres", pgCS)
	if err != nil {
		return err
	}
	log.Println("Connected to pg")
	d.db = pg
	pg.MustExec(createSchema)

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

