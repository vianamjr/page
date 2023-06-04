package main

import (
	"flag"
	"log"
	"os"

	"github.com/vianamjr/page/internal/www"
)

var (
	port = 8080
)

func main() {
	flag.IntVar(&port, "port", port, "port to be used as addr")

	flag.Parse()

	c := www.Config{
		Port: port,
	}

	s, err := www.New(c)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if err := s.ListenAndServe(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
