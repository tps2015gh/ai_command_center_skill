package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	port := "8080"
	url := fmt.Sprintf("http://localhost:%s/viewer.html", port)

	fmt.Printf("🚀 AI Command Center starting at %s\n", url)
	fmt.Println("Press Ctrl+C to stop.")
	
	// Automatically open the viewer in the default browser
	go openBrowser(url)

	// Serve files from the current directory
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

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