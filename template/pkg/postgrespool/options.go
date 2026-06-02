package postgrespool

import "time"

type Option func(*Postgres)

func MaxOpenCons(cons int) Option {
	return func(p *Postgres) {
		p.maxOpenCons = cons
	}
}

func MinIdleCons(cons int) Option {
	return func(p *Postgres) {
		p.minIdleCons = cons
	}
}

func MaxConnIdleTime(t time.Duration) Option {
	return func(p *Postgres) {
		p.maxConnIdleTime = t
	}
}

func MaxConnLifetime(t time.Duration) Option {
	return func(p *Postgres) {
		p.maxConnLifetime = t
	}
}
