package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	port := "8080"
	url := fmt.Sprintf("http://localhost:%s/viewer.html", port)

	// Get absolute path of the directory where the source file is located
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	fmt.Printf("🚀 AI Command Center starting at %s\n", url)
	fmt.Printf("Serving files from: %s\n", baseDir)
	fmt.Println("Press Ctrl+C to stop.")
	
	go openBrowser(url)

	// Serve files relative to the script directory
	// Also allowing access to parent for kanban.json fallback
	fs := http.FileServer(http.Dir(filepath.Join(baseDir, "..")))
	
	// Handle /ai_command_center_skill/ and root
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(baseDir))))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("❌ Error starting server: %v\n", err)
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
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Printf("⚠️ Could not open browser automatically: %v\n", err)
	}
}