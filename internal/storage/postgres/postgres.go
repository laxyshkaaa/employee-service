package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"time"
)

type Storage struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewPostgresStorage(log *slog.Logger, ctx context.Context, dsn string) *Storage {
	ctxWithTmt, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	log.Info("Connecting to PostgreSQL...")

	db, err := pgxpool.New(ctxWithTmt, dsn)
	if err != nil {
		log.Error("Failed to connect to PostgreSQL", "error", err)

	}
	log.Info("Successfully connected to PostgreSQL")

	return &Storage{log: log, db: db}

}
