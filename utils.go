package main

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type File struct {
	Path string `json:"path"`
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

// Converts given files to format
func (a *App) FileConverterDialog() string {
	files, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Only HEIC file types are allowed.",
				Pattern:     "*.heic;*.HEIC",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
		return "Error"
	}

	var outDir string = a.ctx.Value("outDir").(string)

	for _, f := range files {
		splitPath := strings.Split(f, "/")
		filename := strings.Split(splitPath[len(splitPath)-1], ".")[0]

		runtime.LogInfof(a.ctx, "Image info %s\n \n%s", filename, f)
		file, err := os.Open(f)
		if err != nil {
			runtime.LogError(a.ctx, "Error opening file: "+err.Error())
			return "Error"
		}

		img, _, err := image.Decode(file)
		if err != nil {
			file.Close()
			runtime.LogInfof(a.ctx, "Error decoding %s", err)
			return "Error"
		}

		// Encode the image to PNG
		var out bytes.Buffer
		if err := png.Encode(&out, img); err != nil {
			file.Close()
			runtime.LogErrorf(a.ctx, "could not encode image as PNG: %v", err)
			return "Error"
		}

		filename = outDir + filename + ".png"

		// Save the PNG data to a file
		if err := os.WriteFile(filename, out.Bytes(), 0644); err != nil {
			file.Close()
			runtime.LogErrorf(a.ctx, "could not save PNG image as %s: %v", filename, err)
			return "Error"
		}

		runtime.LogInfof(a.ctx, "Image successfully written to %s\n", filename)
		file.Close()
	}

	return "Success"
}
