package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"code.google.com/p/gcfg"
	"flag"
)

type HostPortPair struct {
	Host string
	Port int
}

type Config struct {
	Db     HostPortPair
	Listen HostPortPair
}

func fmtHostPort(hpp HostPortPair) string {
	return fmt.Sprintf("%s:%d", hpp.Host, hpp.Port)
}

func main() {
	// setup default config
	cfg := Config{
		Db: HostPortPair{
			Host: "localhost",
			Port: 6397,
		},
		Listen: HostPortPair{
			Host: "",
			Port: 8080,
		},
	}

	// read config file
	err := gcfg.ReadFileInto(&cfg, "reststorage.config")
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("Failed to parse gcfg data: %s", err)
			return
		}
	}

	flag.StringVar(&cfg.Db.Host, "dbhost", cfg.Db.Host, "database host")
	flag.IntVar(&cfg.Db.Port, "dbport", cfg.Db.Port, "database port")

	flag.StringVar(&cfg.Listen.Host, "listenhost", cfg.Listen.Host, "listen host")
	flag.IntVar(&cfg.Listen.Port, "listenport", cfg.Listen.Port, "listen port")

	flag.Parse()

	InitDB(fmtHostPort(cfg.Db))

	listenAddress := fmtHostPort(cfg.Listen)
	log.Printf("Listening on '%s'...", listenAddress)

	log.Fatal(http.ListenAndServe(listenAddress, NewRouter()))
}
