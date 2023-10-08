package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var listenPort = 9000

type Protocol string

func NewProtocol(name string) Protocol {
	return Protocol(strings.ToUpper(name))
}

type Config map[Protocol]Upstream

type Upstream struct {
	Host string
	Port int
}

var config = Config{
	"IMAP": Upstream{
		Host: "127.0.0.1",
		Port: 1143,
	},
	"SMTP": Upstream{
		Host: "127.0.0.1",
		Port: 1025,
	},
}

func main() {
	http.HandleFunc("/", AuthHandler(config))

	fmt.Printf("Listening on port %d\n", listenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func AuthHandler(config Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		protocol := NewProtocol(req.Header.Get("Auth-Protocol"))
		upstream, ok := config[protocol]
		if !ok {
			w.Header().Set("Auth-Status", "Unknown protocol")
			w.WriteHeader(http.StatusOK)
			return
		}
		w.Header().Set("Auth-Status", http.StatusText(http.StatusOK))
		w.Header().Set("Auth-Server", upstream.Host)
		w.Header().Set("Auth-Port", fmt.Sprintf("%d", upstream.Port))
		w.WriteHeader(http.StatusOK)
	}
}
