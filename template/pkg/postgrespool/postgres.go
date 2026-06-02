package postgrespool

import (
	"context"
	"fmt"
	"math"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultMaxOpenCons     = 30               // max cons
	defaultMinIdleCons     = 5                // min cons
	defaultMaxConnIdleTime = 5 * time.Minute  // lifetime for open cons
	defaultMaxConnLifetime = 30 * time.Minute // lifetime for all cons
)

type Postgres struct {
	maxOpenCons     int
	minIdleCons     int
	maxConnIdleTime time.Duration
	maxConnLifetime time.Duration

	Pool    *pgxpool.Pool
	Builder sq.StatementBuilderType
}

func New(url string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxOpenCons:     defaultMaxOpenCons,
		minIdleCons:     defaultMinIdleCons,
		maxConnIdleTime: defaultMaxConnIdleTime,
		maxConnLifetime: defaultMaxConnLifetime,
	}

	for _, opt := range opts {
		opt(pg)
	}

	// Add SQL Builder - squirrel
	pg.Builder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// Creating config for pool
	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse postgres url: %w", err)
	}
	poolConfig.MaxConns = safeIntToInt32(pg.maxOpenCons)
	poolConfig.MinConns = safeIntToInt32(pg.minIdleCons)
	poolConfig.MaxConnIdleTime = pg.maxConnIdleTime
	poolConfig.MaxConnLifetime = pg.maxConnLifetime

	// Creating pool by poolConfig
	pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres pool: %w", err)
	}

	// Context for ping timeout
	healthCheckCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping to postgres
	if err = pg.Pool.Ping(healthCheckCtx); err != nil {
		return nil, fmt.Errorf("postgres ping timeout or error: %w", err)
	}

	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

func safeIntToInt32(v int) int32 {
	clamped := v & math.MaxInt32

	return int32(clamped)
}