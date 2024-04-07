package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/MaestroError/go-libheif"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	os.Setenv("WEBVIEW2_ADDITIONAL_BROWSER_ARGUMENTS", "--disable-web-security --allow-file-access-from-files --allow-file-access")
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	a.ctx = ctx
	a.ctx = context.WithValue(a.ctx, "outDir", dirname+"/Pictures/")

}

func (a *App) SetOutDir() string {
	outDir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		log.Fatal(err)
		return "error"
	}
	a.ctx = context.WithValue(a.ctx, "outDir", outDir+"/")
	return outDir
}

func (a *App) GetOutDir() string {
	return a.ctx.Value("outDir").(string)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type File struct {
	Path string `json:"path"`
}

// Converts given files to format
func (a *App) FileConverter() string {
	files, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		log.Fatal(err)
		return "Error"
	}

	var outDir string = a.ctx.Value("outDir").(string)

	fmt.Printf("\nfile: %s - %s", files, outDir)
	for _, f := range files {
		splitPath := strings.Split(f, "/")
		filename := strings.Split(splitPath[len(splitPath)-1], ".")

		err := libheif.HeifToPng(f, outDir+filename[0]+".png")
		if err != nil {
			log.Fatal(err)
			return "Error"
		}
	}

	return "Success"

}
