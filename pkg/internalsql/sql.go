package internalsql

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(dataSourceName string) (*gorm.DB, error){
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil{
		log.Fatal("error connecting database %+v\n", err)
		return nil, err
	}
	return db, nil
}