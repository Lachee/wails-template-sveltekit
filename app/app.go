package app

import (
	"context"
	"fmt"
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
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// domready is called when the frontend has finished loading and is ready for use.
func (a *App) DomReady(ctx context.Context) {
}

// beforeclosed is called just before the window closes. Return true to abort the close.
func (a *App) BeforeClose(ctx context.Context) bool {
	return false
}

// shutdown is called when the app closes.
func (a *App) Shutdown(ctx context.Context) {
}

// Greeting says hello to the user
func (a *App) Greeting(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
