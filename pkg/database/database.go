package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustaphalimar/prepilotapp-backend/pkg/config"
	"gorm.io/gorm"
)

type Database struct {
	PgxPool *pgxpool.Pool
	GormDB  *gorm.DB
}

const DatabasePingTimeout = 10

func New(cfg *config.Config) (*Database, error) {
	pgxPool, err := NewPgxDatabase(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize pgx database: %w", err)
	}

	gormDB, err := NewGormDB(cfg)
	if err != nil {
		pgxPool.Close()
		return nil, fmt.Errorf("failed to initialize GORM database: %w", err)
	}

	database := &Database{
		PgxPool: pgxPool,
		GormDB:  gormDB,
	}

	ctx, cancel := context.WithTimeout(context.Background(), DatabasePingTimeout*time.Second)
	defer cancel()
	if err = pgxPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return database, nil
}

// GetGormDB returns the GORM database instance
func (db *Database) GetGormDB() *gorm.DB {
	return db.GormDB
}

// GetPgxPool returns the pgx pool instance
func (db *Database) GetPgxPool() *pgxpool.Pool {
	return db.PgxPool
}

// AutoMigrate runs GORM auto-migration for the provided models
func (db *Database) AutoMigrate(models ...interface{}) error {
	if db.GormDB == nil {
		return fmt.Errorf("GORM database not initialized")
	}

	err := db.GormDB.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("failed to run auto-migration: %w", err)
	}

	return nil
}

func (db *Database) HealthCheck() error {
	if err := db.PgxPool.Ping(context.Background()); err != nil {
		return fmt.Errorf("pgx health check failed: %w", err)
	}

	if db.GormDB != nil {
		sqlDB, err := db.GormDB.DB()
		if err != nil {
			return fmt.Errorf("failed to get GORM underlying DB: %w", err)
		}

		if err := sqlDB.Ping(); err != nil {
			return fmt.Errorf("GORM health check failed: %w", err)
		}
	}

	return nil
}

func (db *Database) Close() error {
	log.Println("closing database connection pool")
	db.PgxPool.Close()

	if db.GormDB != nil {
		if sqlDB, err := db.GormDB.DB(); err == nil {
			sqlDB.Close()
		}
	}

	return nil
}
