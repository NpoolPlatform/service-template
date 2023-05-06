package migrator

import (
	"context"

	"github.com/NpoolPlatform/service-template/pkg/db"
	"github.com/NpoolPlatform/service-template/pkg/db/ent"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

func migrateAutoID(ctx context.Context, table string, tx *ent.Tx) error {
	logger.Sugar().Infow(
		"migrateAutoID",
		"Merge table", table,
	)

	rows, err := tx.
		QueryContext(
			ctx,
			`select null from information_schema.tables where table_name=? and table_schema=?`,
			table,
			"service_template",
		)
	if err != nil {
		logger.Sugar().Infow(
			"migrateAutoID",
			"Query error", err,
		)
		return err
	}
	if !rows.Next() {
		logger.Sugar().Infow(
			"migrateAutoID",
			"Table exist", false,
		)
		rows.Close()
		return nil
	}
	rows.Close()

	rows, err = tx.
		QueryContext(
			ctx,
			`select null from information_schema.columns where table_name=? and table_schema=? and column_name=?`,
			table,
			"service_template",
			"auto_id",
		)
	if err != nil {
		logger.Sugar().Infow(
			"migrateAutoID",
			"Query error", err,
		)
		return err
	}

	if rows.Next() {
		logger.Sugar().Infow(
			"migrateAutoID",
			"Columns exist", true,
		)
		rows.Close()
		return nil
	}
	rows.Close()

	_, err = tx.
		ExecContext(
			ctx,
			`alter table `+table+` add column auto_id int unsigned not null auto_increment unique`,
		)
	if err != nil {
		logger.Sugar().Infow(
			"migrateAutoID",
			"Add auto_id error", err,
		)
		return err
	}
	return nil
}

func Migrate(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := migrateAutoID(ctx, "details", tx); err != nil {
			return err
		}
		if err := migrateAutoID(ctx, "pubsub_messages", tx); err != nil {
			return err
		}
		return nil
	})
}
