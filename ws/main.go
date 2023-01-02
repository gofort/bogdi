package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/ws", wsHandler)

	panic(http.ListenAndServe(":8080", nil))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	timeoutSecStr := r.URL.Query().Get("timeout_sec")
	if timeoutSecStr != "" {
		if dur, err := time.ParseDuration(timeoutSecStr + "s"); err == nil {
			time.Sleep(dur)
		}
	}

	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

	i := 0

	for {

		if err = conn.WriteJSON(&map[string]string{
			"message": "hello world " + fmt.Sprint(i),
		}); err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(time.Second)
		i++
	}
}
