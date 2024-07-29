package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/headers", headersHandler)
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/slow-buffer", slowBufferResponseHandler)

	panic(http.ListenAndServe(":8080", nil))
}

func slowBufferResponseHandler(w http.ResponseWriter, r *http.Request) {
	// Create a large buffer (e.g., 10 MB of 'a')
	const bufferSize = 10 * 1024 * 1024 // 10 MB
	buffer := make([]byte, bufferSize)
	for i := range buffer {
		buffer[i] = 'a'
	}

	// Set the appropriate headers
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	// Define the chunk size
	const chunkSize = 1024 // 1 KB

	// Send the buffer in chunks with a delay between each chunk
	for i := 0; i < len(buffer); i += chunkSize {
		end := i + chunkSize
		if end > len(buffer) {
			end = len(buffer)
		}

		// Write the chunk to the response
		w.Write(buffer[i:end])
		w.(http.Flusher).Flush() // Flush the buffer to the client

		// Simulate network latency or slow response
		time.Sleep(500 * time.Millisecond) // Sleep for 500 milliseconds between chunks
	}
}

func headersHandler(w http.ResponseWriter, r *http.Request) {
	// Iterate over the headers and print them
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	timeoutSecStr := r.URL.Query().Get("timeout_sec")
	if timeoutSecStr != "" {
		if dur, err := time.ParseDuration(timeoutSecStr + "s"); err == nil {
			time.Sleep(dur)
		}
	}

	clientTestStr := r.URL.Query().Get("client_test")
	var clientTest bool = false
	if clientTestStr == "true" {
		clientTest = true
	}

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	if !clientTest {
		i := 0

		for {

			if conn == nil {
				fmt.Println("connection is nil")
				break
			}

			if err = conn.WriteJSON(&map[string]string{
				"message": "hello world " + fmt.Sprint(i),
			}); err != nil {
				fmt.Println(err)
				return
			}

			time.Sleep(time.Second)
			i++
		}
	} else {

		if conn == nil {
			fmt.Println("connection is nil 0")
			return
		}

		if err = conn.WriteJSON(&map[string]string{
			"message": "Send at time 0",
		}); err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(time.Second * 5)

		if conn == nil {
			fmt.Println("connection is nil 5")
			return
		}

		if err = conn.WriteJSON(&map[string]string{
			"message": "Send at time 5",
		}); err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(time.Second * 15)

		if conn == nil {
			fmt.Println("connection is nil 20")
			return
		}

		if err = conn.WriteJSON(&map[string]string{
			"message": "Send at time 20",
		}); err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(time.Second * 15)

		if conn == nil {
			fmt.Println("connection is nil 15")
			return
		}

		if err = conn.WriteJSON(&map[string]string{
			"message": "Send at time 35",
		}); err != nil {
			fmt.Println(err)
			return
		}

	}

}
