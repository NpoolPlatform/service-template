package db

import (
	"context"

	"github.com/NpoolPlatform/service-template/pkg/db/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/mysql"
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
