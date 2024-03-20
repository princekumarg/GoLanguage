package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	user := "root"
	password := "Ecnirp@1Kumartea@"
	database := "movieapi"
	mysqlURL := "%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", fmt.Sprintf(mysqlURL, user, password, database))
	if err != nil {
		panic(err)
	}
	db = d
}
func GetDB() *gorm.DB {
	return db
}
