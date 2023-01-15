package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
	Grade         float64 `json:"grade"`
	TotalPossible float64 `json:"totalPossible"`
	Percentage    float64 `json:"percentage"`
	Weight        float64 `json:"weight"`
}

// WeightedClass struct
type WeightedClass struct {
	ClassCode string  `json:"classCode"`
	GradeList []Grade `json:"gradeList"`
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

// EnsureBaseDir ensures the base directory exists
func EnsureBaseDir(path string) error {
	// Get the base directory
	baseDir := path
	// Ensure the base directory exists
	if _, Error := os.Stat(baseDir); os.IsNotExist(Error) {
		Error = os.MkdirAll(baseDir, 0755)
		if Error != nil {
			return Error
		}
	}
	return nil
}

func checkError(Error error) {
	if Error != nil {
		fmt.Println("Error:", Error)
	}
}

// Load a json file from the AppData Local directory
func (a *App) LoadJSONFile() string {
	// Get os specific AppData Local path
	appDataLocal, Error := os.UserCacheDir()
	checkError(Error)

	// Open the file
	path := filepath.Join(appDataLocal, "WeightedGradeCalculator/grades.json")

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		// path does *not* exist

		checkError(err)

		return ""
	}

	file, Error := os.ReadFile(path)
	checkError(Error)

	fmt.Println("File:", string(file))
	return string(file)

	// ONLY NEEDED FOR TESTING
	// data := WeightedClass{}
	// // Convert json to WeightedClass struct
	// Error = json.Unmarshal(file, &data)
	// checkError(Error)

	// // Print WeightedClass struct to the console
	// fmt.Printf("Loading WeightedClass: ClassCode: %s, \n",
	// 	data.ClassCode,
	// )
	// for _, grade := range data.GradeList {
	// 	fmt.Printf("Grade: %f, TotalPossible: %f, Percentage: %f, Weight: %f \n",
	// 		grade.Grade,
	// 		grade.TotalPossible,
	// 		grade.Percentage,
	// 		grade.Weight,
	// 	)
	// }

	// // Convert WeightedClass struct to json
	// jsonData, Error := json.Marshal(data)
	// checkError(Error)

	// return string(jsonData)
}

// Save a json file to the AppData Local directory
func (a *App) SaveJSONFile(Json string) {
	// Get os specific AppData Local path
	appDataLocal, Error := os.UserCacheDir()
	checkError(Error)

	// Create a new file
	path := filepath.Join(appDataLocal, "WeightedGradeCalculator")

	// Ensure the base directory exists
	Error = EnsureBaseDir(path)
	checkError(Error)

	file, Error := os.Create(path + "/grades.json")
	checkError(Error)

	// Close the file
	defer file.Close()
	// Write to the file
	file.WriteString(Json)
}

// Print WeightedClass struct to the console
func (a *App) SaveWeightedClass(w WeightedClass) {
	fmt.Printf("Saving WeightedClass: ClassCode: %s, \n",
		w.ClassCode,
	)
	for _, grade := range w.GradeList {
		fmt.Printf("Grade: %f, TotalPossible: %f, Percentage: %f, Weight: %f \n",
			grade.Grade,
			grade.TotalPossible,
			grade.Percentage,
			grade.Weight,
		)
	}

	// Convert WeightedClass struct to json
	jsonData, Error := json.Marshal(w)
	checkError(Error)

	a.SaveJSONFile(string(jsonData))
}
