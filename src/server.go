package src

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/laughing-nerd/jdf/utils"
)

func StartServer(resultch chan string) {
	fmt.Println("Web mode is enabled. Visit http://localhost:6969")

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

	if err := http.ListenAndServe(":6969", nil); err != nil {
		panic("Unable to serve on port 6969\n" + err.Error())
	}
}
