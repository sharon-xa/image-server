package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func (app *Application) saveImage(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "couldn't parse form", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "couldn't parse image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	splitedFileName := strings.Split(fileHeader.Filename, ".")
	fileExt := splitedFileName[len(splitedFileName)-1]
	imagePath := filepath.Join(app.imgDirPath, strings.Join([]string{uuid.NewString(), fileExt}, "."))

	dst, err := os.Create(imagePath)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(imagePath))
}

func (app *Application) getImage(w http.ResponseWriter, r *http.Request) {
	imagePath := r.URL.Query().Get("image")
	if imagePath == "" {
		app.logger.Error("no path provided")
		http.Error(w, "no path provided", http.StatusBadRequest)
		return
	}
	http.ServeFile(w, r, imagePath)
}

func (app *Application) deleteImage(w http.ResponseWriter, r *http.Request) {
	imagePath := r.URL.Query().Get("image")
	if imagePath == "" {
		app.logger.Error("no path provided")
		http.Error(w, "no path provided", http.StatusBadRequest)
		return
	}

	err := os.Remove(imagePath)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "couldn't remove image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}
