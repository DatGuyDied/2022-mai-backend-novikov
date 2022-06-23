package postgres

import (
	"context"
	"fmt"
	"os"

	"entgo.io/ent/dialect"
	"github.com/DatGuyDied/2022-mai-backend-novikov/ent"

	_ "github.com/lib/pq"
)

type postgres struct {
	client *ent.Client
}

func New(ctx context.Context) (*postgres, error) {
	client, err := ent.Open(dialect.Postgres, os.Getenv("POSTGRES"))
	if err != nil {
		return nil, fmt.Errorf("can not open db: %w", err)
	}

	err = client.Schema.Create(ctx)
	if err != nil {
		return nil, fmt.Errorf("can not migrate schema: %w", err)
	}

	return &postgres{client: client}, nil
}
