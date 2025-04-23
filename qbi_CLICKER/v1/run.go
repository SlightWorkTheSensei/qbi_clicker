package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/go-vgo/robotgo" // For mouse automation
)

var (
	clicking = false
	interval = 1000 // Default interval in milliseconds
)

// Test click route
func testClick(w http.ResponseWriter, r *http.Request) {
	robotgo.Click() // Perform a single click
	w.Write([]byte("Test click performed"))
}

// Start clicking
func startClicking(w http.ResponseWriter, r *http.Request) {
	if clicking {
		w.Write([]byte("Already clicking!"))
		return
	}
	clicking = true
	go func() {
		for clicking {
			robotgo.Click()
			time.Sleep(time.Duration(interval) * time.Millisecond)
		}
	}()
	w.Write([]byte("Clicking started"))
}

// Stop clicking
func stopClicking(w http.ResponseWriter, r *http.Request) {
	clicking = false
	w.Write([]byte("Clicking stopped"))
}

// Set interval
func setInterval(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	newInterval := r.FormValue("interval")
	if newInterval == "" {
		w.Write([]byte("Interval not provided"))
		return
	}
	fmt.Sscanf(newInterval, "%d", &interval)
	w.Write([]byte(fmt.Sprintf("Interval set to %d ms", interval)))
}

// Serve the HTML/JS GUI
func serveFrontend(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", serveFrontend)
	http.HandleFunc("/test", testClick)    // Test click route
	http.HandleFunc("/start", startClicking)
	http.HandleFunc("/stop", stopClicking)
	http.HandleFunc("/set-interval", setInterval)

	port := ":8090" // Custom port
	fmt.Printf("Server started on %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
