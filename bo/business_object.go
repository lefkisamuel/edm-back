package bo

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"reflect"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "samuellefki"
	password = ""
	dbname   = "protodb"
)

type iBusinessObject interface{
	getId() uint
}

type BusinessObjectProxy struct {
	bo      iBusinessObject
	updates map[string]interface{}
}

type BusinessObject struct {
	gorm.Model
	iBusinessObject
}

func (bo BusinessObject) getId() uint{
	return 0//bo.ID
}

type Product struct {
	BusinessObject
	Code  string
	Price uint
}

func NewBusinessObjectProxy(iBusinessObject iBusinessObject) *BusinessObjectProxy {
	return &BusinessObjectProxy{bo: iBusinessObject, updates: make(map[string]interface{})}
}

func (bop *BusinessObjectProxy) SetFieldValue(fieldName string, value interface{}) error {
	bo := bop.getBusinessObject()

	if bo == nil {
		return errors.New("no business object associated")
	}

	boType := reflect.ValueOf(bo)
	if boType.Kind() != reflect.Ptr {
		return errors.New("a business object is expected to be of type pointer")
	}

	field := reflect.ValueOf(bop.bo).Elem().FieldByName(fieldName)
	if !field.IsValid() {
		return errors.New("field not found")
	}

	switch field.Kind() {
	case reflect.Uint:
		field.SetUint(uint64(value.(int)))
	case reflect.Float64:
		field.SetFloat(value.(float64))
	case reflect.String:
		field.SetString(value.(string))
	case reflect.Bool:
		field.SetBool(value.(bool))
	default :
		return errors.New("type not supported")
	}

	bop.updates[fieldName] = value

	fmt.Println("field changed")
	return nil
}

func (bop *BusinessObjectProxy) Save() error {
	if bop.isNew(){

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err := gorm.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}

		//Migrate the schema
		bo := bop.getBusinessObject()
		db.AutoMigrate(*bo)
		db.Create(*bo)
	}
	return nil
}

func (bop *BusinessObjectProxy) getBusinessObject() *iBusinessObject{
	return &bop.bo
}

func (bop *BusinessObjectProxy) hasChange() bool{
	return len(bop.updates) != 0
}

func (bop *BusinessObjectProxy) isNew() bool{
	return true
}
