package main

import (
	"context"
	"log"
	"os"

	_ "github.com/strukturag/libheif/go/heif"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	const outDirKey string = "outDir"

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	a.ctx = ctx
	a.ctx = context.WithValue(a.ctx, outDirKey, dirname+"/Downloads/")

	go func() {
		a.initRouter()
	}()
}
