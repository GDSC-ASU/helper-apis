package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type orientation string

const (
	landscape orientation = "landscape"
	portrait  orientation = "portrait"
)

var imagesPaths []string

func init() {
	dir := "./files/images/"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		basename := filepath.Base(dir)
		removeIndex := strings.Index(path, basename)
		imagesPaths = append(imagesPaths, "/files/"+path[removeIndex:])
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func getRandomImage(o orientation) string {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	found := ""
	for !strings.Contains(found, string(o)) {
		found = imagesPaths[r.Intn(len(imagesPaths))]
	}
	return found
}

func handleGetImage(w http.ResponseWriter, r *http.Request) {
	orientation := landscape
	switch r.URL.Query().Get("orientation") {
	case string(landscape):
		orientation = landscape
	case string(portrait):
		orientation = portrait
	}
	count := 1
	if parsedCount, err := strconv.Atoi(r.URL.Query().Get("count")); err == nil {
		count = parsedCount
	}

	images := []string{getRandomImage(orientation)}
	for i := 1; i < count; i++ {
		images = append(images, getRandomImage(orientation))
	}

	_ = json.NewEncoder(w).Encode(map[string]any{
		"images": images,
	})
}
