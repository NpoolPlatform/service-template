package db

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/service-template/pkg/db/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/mysql"

	_ "github.com/NpoolPlatform/service-template/pkg/db/ent/runtime" //nolint
)

func client() (*ent.Client, error) {
	conn, err := mysql.GetConn()
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.MySQL, conn)
	return ent.NewClient(ent.Driver(drv)), nil
}

func Init() error {
	cli, err := client()
	if err != nil {
		return err
	}
	return cli.Schema.Create(context.Background())
}

func Client() (*ent.Client, error) {
	return client()
}

func WithTx(ctx context.Context, tx *ent.Tx, fn func(ctx context.Context) error) error {
	defer func() {
		if v := recover(); v != nil {
			err := tx.Rollback()
			if err != nil {
				logger.Sugar().Errorf("fail to rollback: %v", err)
			}
			panic(v)
		}
	}()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := fn(ctx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %v (%v)", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %v", err)
	}
	return nil
}

func Do(ctx context.Context, fn func(ctx context.Context, cli *ent.Client) error) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cli, err := Client()
	if err != nil {
		return fmt.Errorf("fail get db client: %v", err)
	}

	if err := fn(ctx, cli); err != nil {
		return err
	}
	return nil
}

type Entity struct {
	Tx *ent.Tx
}

func NewEntity(ctx context.Context, _tx *ent.Tx) (*Entity, error) {
	if _tx != nil {
		return &Entity{
			Tx: _tx,
		}, nil
	}

	cli, err := Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}
	_tx, err = cli.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail get client transaction: %v", err)
	}

	return &Entity{
		Tx: _tx,
	}, nil
}
