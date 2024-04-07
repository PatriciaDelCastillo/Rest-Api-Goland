package database

import (
	"fmt"
	"strconv"
	"github.com/PatriciaDelCastillo/Rest-Api-Goland/config"
	"github.com/PatriciaDelCastillo/Rest-Api-Goland/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")

	port, err:=strconv.ParseUint(p,10,32)

	if err !=nil {
		panic("Falla del puerot de la base dato port")
	}

	dsn:= fmt.Sprintf("host=%s  port=%d  user=%s password=%s dbname=%s sslmode=disable ",
	config.Config("DB_HOST"),port,config.Config("DB_USER"),config.Config("DB_PASSWORD"),config.Config("DB_NAME"))

	DB,err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!= nil{
			panic("Falla al conectar a la Base de datos")
	}
	fmt.Println("Conexion abierta de la base de datos")
	DB.AutoMigrate(&models.Task{})
	fmt.Println("Base de dato Migrada")
}