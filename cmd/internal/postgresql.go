package internal

import (
	"fmt"
	"log"

	"github.com/Alptahta/user-service/internal/storage/model"
	"github.com/Alptahta/user-service/pkg/utility"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQL(postgreSQLConfig *utility.PostgreSQLConfig) (*gorm.DB, error) {

	dbURL := fmt.Sprintf("%s://%s:%s@localhost:%s/%s", postgreSQLConfig.DBDriver, postgreSQLConfig.DBUsername,
		postgreSQLConfig.DBPassword, postgreSQLConfig.DBPort, postgreSQLConfig.DBName)
	fmt.Println(dbURL)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalln(err)
	}

	return db, err
}
