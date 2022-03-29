package database

import (
	"log"

	"github.com/hrqcds/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DB_conn() {
	dns := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dns))

	DB.AutoMigrate(&models.Aluno{})

	if err != nil {
		log.Panic("Erro na conexão com o banco de dados", err.Error())
	}

}
