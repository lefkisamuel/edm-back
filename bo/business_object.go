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
	return bo.ID
}

type Product struct {
	BusinessObject
	Code  string
	Price uint
	ProductClientId uint
	ProductClient ProductClient
}

type ProductClient struct{
	gorm.Model
	FirstName string
	LastName string
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

	var oldVal, newVal interface{}

	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		oldVal = field.Uint()
		newVal = uint64(value.(int))
		field.SetUint(newVal.(uint64))
	case reflect.Float64:
		oldVal = field.String()
		newVal = value.(float64)
		field.SetFloat(value.(float64))
	case reflect.String:
		oldVal = field.String()
		newVal = value.(string)
		field.SetString(newVal.(string))
	case reflect.Bool:
		field.SetBool(value.(bool))
	default :
		return errors.New("type not supported")
	}

	if oldVal == newVal {
		return nil //nothing to do here...
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
		db.SingularTable(true)

		//Migrate the schema
		bo := bop.getBusinessObject()


		//db.AutoMigrate(*bo)
		db.AutoMigrate(&ProductClient{},&Product{})
		//db.Create(&ProductClient{})
		db.Create(*bo)

	}
	return nil
}

func (bop *BusinessObjectProxy) getBusinessObject() *iBusinessObject{
	return &bop.bo
}

func (bop *BusinessObjectProxy) hasChanged() bool{
	return len(bop.updates) != 0
}

func (bop *BusinessObjectProxy) isNew() bool{
	return bop.bo.getId() == 0
}
