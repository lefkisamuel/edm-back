package repository

import (
	"edm-back/entity"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "samuellefki"
	password = ""
	dbname   = "protodb"
)

type TypeDefinitionRepository interface {
	Save(typeDefinition entity.TypeDefinition)
	FindAll() []entity.TypeDefinition
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func New() TypeDefinitionRepository {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("Failed to connect to database")
	}
	conn.AutoMigrate(&entity.TypeDefinition{}, &entity.TypeDefinitionTable{}, &entity.TypeDefinitionField{})
	return &database{
		connection: conn,
	}
}

func (db *database) Save(typeDefinition entity.TypeDefinition) {
	db.connection.Save(&typeDefinition)
}
func (db *database) FindAll() []entity.TypeDefinition{
	var typeDefinitions []entity.TypeDefinition
	db.connection.Set("gorm:auto_preload", true).Find(&typeDefinitions)
	return typeDefinitions
}

func (db *database) CloseDB(){
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

