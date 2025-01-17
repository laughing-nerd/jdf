package src

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/laughing-nerd/jdf/utils"
)

func StartServer(resultch chan string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, utils.GetHTML())
	})

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		for {
			select {
			case data := <-resultch:
				data = utils.RemoveANSIColor(data)
				data = strings.ReplaceAll(data, "\n", "<br>")
				data = strings.ReplaceAll(data, " ", "&nbsp;")
				data = strings.ReplaceAll(data, "&nbspcustom;", " ")
				fmt.Fprintf(w, "data: %s\n\n", data)

				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
			}
		}
	})

	port := getFreePort()
	fmt.Println("Web mode enabled! jdf is running on http://localhost" + port + " 🚀")
	if err := http.ListenAndServe(port, nil); err != nil {
		panic("Unable to serve on port " + port + "\n" + err.Error())
	}
}

func getFreePort() string {
	port := utils.FlagPort

	for port < 65535 {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
      listener.Close()
			break
		}
		fmt.Println("Port", port, "is already in use. Trying next port...")
		port++
	}

	if port == 65535 {
		panic("No free ports available")
	}
	return fmt.Sprintf(":%d", port)
}
