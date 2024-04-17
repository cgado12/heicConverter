package main

import (
	"bytes"
	"image"
	"image/png"
	"net/http"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (a *App) initRouter() {
	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		a.FileConverterRoute(w, r)
	})
	http.ListenAndServe(":3269", nil)
}

func (a *App) FileConverterRoute(w http.ResponseWriter, r *http.Request) {
	var outDir string = a.ctx.Value("outDir").(string)
	// Set max memory at 50MB
	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		http.Error(w, "Error parsing multipart form: "+err.Error(), http.StatusInternalServerError)
		return
	}

	files := r.MultipartForm.File["upload"]

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error opening file: "+err.Error(), http.StatusInternalServerError)
			return
		}

		img, _, err := image.Decode(file)
		if err != nil {
			file.Close()
			runtime.LogInfof(a.ctx, "Error decoding %s", err)
			return
		}

		// Encode the image to PNG
		var out bytes.Buffer
		if err := png.Encode(&out, img); err != nil {
			file.Close()
			runtime.LogInfof(a.ctx, "could not encode image as PNG: %v", err)
			return
		}
		filename := outDir + fileHeader.Filename + ".png"
		// Save the PNG data to a file
		if err := os.WriteFile(filename, out.Bytes(), 0644); err != nil {
			file.Close()
			runtime.LogInfof(a.ctx, "could not save PNG image as %s: %v", filename, err)
			return
		}

		runtime.LogInfof(a.ctx, "Image successfully written to %s\n", filename)
		file.Close()
	}
}
