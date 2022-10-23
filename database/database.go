package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	models "setup/models"
)

var db *gorm.DB
var err error
/* var dbAddress = env.GoDotEnvVariable("DB_USERNAME") + ":" +
	env.GoDotEnvVariable("DB_PASSWORD") +
	"@tcp(" +
	env.GoDotEnvVariable("DB_HOST") +
	":" +
	env.GoDotEnvVariable("DB_PORT") +
	")/"
 */
func Init() gorm.DB {
	dns := "admin:root12345@tcp(aws-simplified.c8i70zggrjmm.us-east-1.rds.amazonaws.com:3306)/cryptology_homework?parseTime=true"
	db, err = gorm.Open(mysql.Open(dns))
	db.AutoMigrate(&models.User{}, &models.Record{})
	if err != nil {
		panic(err.Error())
	}
	return *db
}
func DBManager() gorm.DB {
	return *db
}
