package main

import (
	"context"

	"github.com/gsvd/goeland/internal/store"
)

type App struct {
	ctx   context.Context
	store *store.Store
}

func NewApp(store *store.Store) *App {
	return &App{store: store}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
