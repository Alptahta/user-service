package main

import (
	"log"

	"github.com/Alptahta/user-service/cmd/internal"
	"go.uber.org/zap"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}
}

func run() error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	pool, err := internal.NewPostgreSQL()
	if err != nil {
		return err
	}
	return nil
}
