package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type PostgresDB struct {
	Pool *pgxpool.Pool
}

func NewPostgresDB(ctx context.Context, connString string) (*PostgresDB, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db := &PostgresDB{Pool: pool}

	err = db.migrate("migrations")
	if err != nil {
		err = db.migrate("../migrations")
		if err != nil {
			return nil, fmt.Errorf("failed to run migrations: %w", err)
		}
	}

	log.Println("Connected to PostgreSQL and migrations completed")
	return db, nil
}

func (db *PostgresDB) migrate(directory string) error {
	sqlDB := stdlib.OpenDBFromPool(db.Pool)
	defer sqlDB.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}

	if err := goose.Up(sqlDB, directory); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func (db *PostgresDB) Close() {
	db.Pool.Close()
}

// GetStdDB returns a standard database/sql.DB for compatibility
func (db *PostgresDB) GetStdDB() *sql.DB {
	return stdlib.OpenDBFromPool(db.Pool)
}
