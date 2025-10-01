package connection

import (
	"fmt"

	"github.com/sale-tickets/manager-api/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func (a *DatabaseConfig) connection() {
	var err error

	dns := fmt.Sprintf("host=%s user=%s password=%s  dbname=%s port=%s sslmode=disable ",
		a.Host,
		a.Username,
		a.Password,
		a.Name,
		a.Port,
	)
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("error connection database: ", err.Error())
		return
	}

	err = db.AutoMigrate(
		model.CinemaRoom{},
		model.MovieTheater{},
	)
	if err != nil {
		fmt.Println("error migrate database: ", err.Error())
		return
	}

	fmt.Println("conection database successfully!")
}

func (a *DatabaseConfig) GetConection() *gorm.DB {
	return db
}
