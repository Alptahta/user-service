package main

import (
	"fmt"
	"log"

	"github.com/Alptahta/user-service/pkg/utility"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}
}

func run() error {
	postgreSQLConfig, err := utility.LoadPostgreSqlConfig(".")
	if err != nil {
		return err
	}
	fmt.Println(postgreSQLConfig.DBDriver)
	return nil
}
