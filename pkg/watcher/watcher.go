//nolint:dupl
package watcher

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
)

func Watch(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			logger.Sugar().Infow("Watch", "State", "Tick")
		case <-ctx.Done():
			if ctx.Err() == nil {
				logger.Sugar().Infow("Watch", "State", "Done")
				return
			}
			logger.Sugar().Errorw("Watch", "State", "Error", "Error", ctx.Err())
			return
		}
	}
}
