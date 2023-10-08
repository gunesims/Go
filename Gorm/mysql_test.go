package Gorm

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGormMysql(t *testing.T) {

	db, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}
