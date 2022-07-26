package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Alptahta/user-service/cmd/internal"
	"github.com/Alptahta/user-service/internal/http/rest"
	"github.com/Alptahta/user-service/internal/service"
	"github.com/Alptahta/user-service/internal/storage/postgresql"
	"github.com/Alptahta/user-service/pkg/utility"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	postgreSQLConfig, err := utility.LoadPostgreSqlConfig()
	if err != nil {
		log.Fatalf("Couldn't load postgresql configs: %s", err)
	}
	fmt.Println(postgreSQLConfig)

	db, err := internal.NewPostgreSQL(postgreSQLConfig)
	if err != nil {
		log.Fatalf("Couldn't load postgresql configs: %s", err)
	}

	srv := newServer(db)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Couldn't start server: %s", err)
	}
}

func newServer(db *gorm.DB) *http.Server {
	r := gin.Default()

	repo := postgresql.NewUser(db)
	svc := service.NewUserService(repo)
	rest.NewUserHandler(svc).Register(r)

	address := "0.0.0.0:9234"

	return &http.Server{
		Handler:           r,
		Addr:              address,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       1 * time.Second,
	}
}
