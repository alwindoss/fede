// package server contains all the server related files
package server

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

type Config struct {
	Port string
	Host string
}

func Run(cfg *Config) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Time now is: %s", time.Now().String())
	})
	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	http.ListenAndServe(addr, nil)
	return nil
}
