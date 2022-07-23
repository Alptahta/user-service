package main

import (
	"fmt"
	"log"

	"github.com/Alptahta/user-service/pkg/utility"
)

func main() {
	err := initDB()
	if err != nil {
		log.Fatalf("Couldn't initDB: %s", err)
	}
}

func initDB() error {
	postgreSQLConfig, err := utility.LoadPostgreSqlConfig()
	if err != nil {
		return err
	}
	fmt.Println(postgreSQLConfig)

	return nil
}
