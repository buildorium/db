package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB(appName string, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("root:root@tcp(mysql-%s-%s:3306)/default?charset=utf8&parseTime=True&loc=Local", appName, dbName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))

	return db, err
}
