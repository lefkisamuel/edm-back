package entity

import (
	"github.com/jinzhu/gorm"
)

type TypeDefinition struct {
	gorm.Model `json:"-"`
	Name string `json:"name"`
	Tables []TypeDefinitionTable `json:"tables"`
}
type TypeDefinitionTable struct {
	gorm.Model `json:"-"`
	TypeDefinitionID uint `json:"-"`
	TableName string `json:"table_name"`
	Fields []TypeDefinitionField `json:"fields"`
}
type TypeDefinitionField struct {
	gorm.Model `json:"-"`
	TypeDefinitionTableID uint `json:"-"`
	FieldName string `json:"field_name"`
	FieldType string `json:"field_type"`
	FieldSize uint 	 `json:"field_size"`
}