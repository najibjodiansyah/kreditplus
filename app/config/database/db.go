package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	userName := fmt.Sprintf("%v", viper.GetString("USERNAME"))
	password := fmt.Sprintf("%v", viper.GetString("PASSWORD"))
	host := fmt.Sprintf("%v", viper.GetString("HOST"))
	port := fmt.Sprintf("%v", viper.GetString("PORT"))
	database := fmt.Sprintf("%v", viper.GetString("DATABASE"))
	db, err := gorm.Open(mysql.Open(generateUri(userName, password, host, port, database)), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
		os.Exit(0)
	}
	return db
}

func generateUri(userName, password, host, port, database string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", userName, password, host, port, database)
}
