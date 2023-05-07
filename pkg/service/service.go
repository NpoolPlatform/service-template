package service

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/watcher"
)

var w *watcher.Watcher

func Watch(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	w = watcher.NewWatcher()

	for {
		select {
		case <-ticker.C:
			logger.Sugar().Infow(
				"Watch",
				"State", "Tick",
			)
		case <-ctx.Done():
			logger.Sugar().Infow(
				"Watch",
				"State", "Done",
				"Error", ctx.Err(),
			)
			close(w.ClosedChan())
			return
		case <-w.CloseChan():
			close(w.ClosedChan())
		}
	}
}

func Shutdown() {
	if w != nil {
		w.Shutdown()
	}
}
