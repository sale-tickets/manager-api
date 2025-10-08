package connection

import (
	"fmt"

	"github.com/sale-tickets/manager-api/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionPsql(config *Config) *gorm.DB {
	var err error

	dns := fmt.Sprintf("host=%s user=%s password=%s  dbname=%s port=%s sslmode=disable ",
		config.Database.Host,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("error connection database: ", err.Error())
		return nil
	}

	err = db.AutoMigrate(
		model.CinemaRoom{},
		model.MovieTheater{},
		model.TheaterSeating{},
	)
	if err != nil {
		fmt.Println("error migrate database: ", err.Error())
		return nil
	}

	fmt.Println("conection database successfully!")
	return db
}
