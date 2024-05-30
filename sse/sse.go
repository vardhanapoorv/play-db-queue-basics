package sse

import (
	"fmt"
	"net/http"
	"time"
)

func StartServer() {
	http.HandleFunc("/events", handleSSE)
	// Serve the HTML file
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}

func handleSSE(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to text/event-stream
	w.Header().Set("Content-Type", "text/event-stream")
	// Set additional headers to allow cross-origin requests if needed
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Create a new ticker that sends a message every second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Infinite loop to continuously send messages
	for {
		select {
		case <-ticker.C:
			// Construct an SSE message with a unique event ID and data
			event := fmt.Sprintf("id: %d\ndata: %s\n\n", time.Now().Unix(), "Hello from server!")
			// Write the event to the ResponseWriter
			_, err := fmt.Fprint(w, event)
			if err != nil {
				// If writing fails, the client may have disconnected, so we break the loop
				return
			}
			// Flush the response to ensure the event is sent immediately
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		case <-r.Context().Done():
			// If the client closes the connection, we break the loop
			return
		}
	}
}
