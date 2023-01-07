package main

import (
	"context"
	"fmt"
	"os"
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
	a.ctx = ctx
}

// Grade struct as a json object
// id: number
// grade: number
// totalPossible: number
// percentage: number
// weight: number

// Grade struct
type Grade struct {
	ID            int     `json:"id"`
	Grade         float64 `json:"grade"`
	TotalPossible float64 `json:"totalPossible"`
	Percentage    float64 `json:"percentage"`
	Weight        float64 `json:"weight"`
}

// Print AppData Local to the console
func (a *App) PrintAppDataLocal() {
	// Get os specific AppData Local path
	appDataLocal, Error := os.UserCacheDir()
	if Error != nil {
		fmt.Println("Error:", Error)
	}
	fmt.Println("AppData Local:", appDataLocal)
}
