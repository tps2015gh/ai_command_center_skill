package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type ServerInfo struct {
	Language string `json:"language"`
	Command  string `json:"command"`
	BaseDir  string `json:"baseDir"`
}

func main() {
	port := "8080"
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	projectRootDir := filepath.Join(baseDir, "..")

	http.HandleFunc("/server-info", func(w http.ResponseWriter, r *http.Request) {
		info := ServerInfo{
			Language: "Go",
			Command:  "go run serve.go",
			BaseDir:  baseDir,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(info)
	})

	http.HandleFunc("/kanban.json", func(w http.ResponseWriter, r *http.Request) {
		localPath := filepath.Join(baseDir, "kanban.json")
		rootPath := filepath.Join(projectRootDir, "kanban.json")

		target := ""
		if _, err := os.Stat(localPath); err == nil {
			target = localPath
		} else if _, err := os.Stat(rootPath); err == nil {
			target = rootPath
		}

		if target != "" {
			w.Header().Set("Content-Type", "application/json")
			http.ServeFile(w, r, target)
		} else {
			http.Error(w, "kanban.json not found", http.StatusNotFound)
		}
	})

	// Serve UI files
	http.Handle("/", http.FileServer(http.Dir(baseDir)))

	url := fmt.Sprintf("http://localhost:%s/viewer.html", port)
	fmt.Printf("🚀 AI Command Center starting at %s\n", url)
	fmt.Printf("Serving from: %s\n", baseDir)
	
	go openBrowser(url)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
	}
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		fmt.Printf("⚠️ Browser error: %v\n", err)
	}
}