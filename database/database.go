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
	dns := "bf455b2579977c:43489f3a@tcp(us-cdbr-east-06.cleardb.net:3306)/heroku_b5def9f32733d0c"
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
