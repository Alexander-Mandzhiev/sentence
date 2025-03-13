package dbManager

import (
	"backend/pkg/config"
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenPostgreSQLConnection(cfg config.PostgreSQLConfig, logger *slog.Logger) (*pgxpool.Pool, error) {
	const op = "db_manager.OpenPostgreSQLConnection"

	if cfg.ConnectionString == "" {
		return nil, fmt.Errorf("%s: database configuration is not initialized", op)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	poolConfig, err := pgxpool.ParseConfig(cfg.ConnectionString)
	if err != nil {
		logger.Error(op, slog.String("message", "Error parsing PostgreSQL connection string"), slog.String("storagePath", maskPassword(cfg.ConnectionString)), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	poolConfig.MaxConnLifetime = cfg.ConnMaxLifetime
	poolConfig.MaxConns = cfg.MaxOpenConnections
	poolConfig.MinConns = cfg.MaxIdleConnections

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		logger.Error(op, slog.String("message", "Error opening connection to PostgreSQL database"), slog.String("storagePath", maskPassword(cfg.ConnectionString)), slog.Any("error", err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err = pool.Ping(ctx); err != nil {
		logger.Error(op, slog.String("message", "Error testing PostgreSQL database connection"), slog.Any("error", err))
		pool.Close()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	logger.Info(op, slog.String("message", "Opened connection to PostgreSQL database"), slog.Int("max_conns", int(poolConfig.MaxConns)), slog.Int("min_conns", int(poolConfig.MinConns)), slog.Duration("conn_max_lifetime", poolConfig.MaxConnLifetime))

	return pool, nil
}

// ClosePostgreSQLConnection закрывает пул подключений к PostgreSQL.
func ClosePostgreSQLConnection(pool *pgxpool.Pool, logger *slog.Logger) error {
	const op = "db_manager.ClosePostgreSQLConnection"

	if pool == nil {
		logger.Warn(op, slog.String("message", "Database connection is already closed or not initialized"))
		return nil
	}

	stats := pool.Stat()
	if stats.TotalConns() == 0 {
		logger.Warn(op, slog.String("message", "Database connection is already closed or not initialized"))
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool.Close()

	select {
	case <-ctx.Done():
		logger.Warn(op,
			slog.String("message", "Timeout while closing PostgreSQL database connection"), slog.Any("error", ctx.Err()))
		return fmt.Errorf("%s: timeout while closing database connection", op)
	default:
		logger.Info(op, slog.String("message", "PostgreSQL database connection closed"))
		return nil
	}
}

// maskPassword маскирует пароль в строке подключения.
func maskPassword(dsn string) string {
	parts := strings.Split(dsn, "?")
	if len(parts) > 0 {
		return parts[0] + "?password=***"
	}
	return dsn
}
