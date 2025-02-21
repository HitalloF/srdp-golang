package config

import (
	"fmt"
	"go-app/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=admin password=admin dbname=goapp port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Migrar tabelas automaticamente
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Erro ao migrar banco de dados:", err)
	}

	DB = database
	fmt.Println("Banco de dados conectado e migrado com sucesso!")
}
