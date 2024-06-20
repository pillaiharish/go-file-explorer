package handlers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type FileInfo struct {
	Name  string
	Path  string
	IsDir bool
}

func FileExplorerHandler(basePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files, err := ioutil.ReadDir(basePath)
		if err != nil {
			http.Error(w, "Unable to read files", http.StatusInternalServerError)
			return
		}

		var fileInfos []FileInfo
		for _, file := range files {
			fileInfos = append(fileInfos, FileInfo{
				Name:  file.Name(),
				Path:  filepath.Join("/files", file.Name()),
				IsDir: file.IsDir(),
			})
		}

		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, fileInfos)
	}
}

func FileHandler(basePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join(basePath, r.URL.Path[len("/files/"):])
		http.ServeFile(w, r, filePath)
	}
}
