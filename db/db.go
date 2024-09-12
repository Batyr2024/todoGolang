package db

import (
	"context"
	"github.com/Batyr2024/todoGolang/config"
	"github.com/Batyr2024/todoGolang/db/sqlc"
	"github.com/jackc/pgx/v5"
	"log"
)

func Connect(cfg config.Config) *pgx.Conn {
	conn, errCfg := pgx.Connect(context.Background(), cfg.DBUrl)
	if errCfg != nil {
		log.Fatal(errCfg)
	}
	sqlc.New(conn)
	return conn
}
