package dbrepo

import (
	"database/sql"

	"github.com/Lasang3012/go-Fort-Smith-Bed-and-Breakfast/internal/config"
	"github.com/Lasang3012/go-Fort-Smith-Bed-and-Breakfast/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
