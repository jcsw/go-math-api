package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/jcsw/go-math-api/pkg"
)

var env string

func main() {
	flag.StringVar(&env, "env", "prod", "app environment")
	flag.Parse()

	server := pkg.Server{}
	server.Initialize(env)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		server.Stop()
	}()

	server.Start()
}
