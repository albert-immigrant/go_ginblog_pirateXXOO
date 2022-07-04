package model

import (
	"fmt"
	//"github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
/*
	 dsn := fmt.Sprintf("host= %s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		utils.DbHost,
		utils.DbUser,
		utils.DbPassWord,
		utils.DbName,
		utils.DbPort,
		utils.DbSslMode,
		utils.DbTimeZone,
	)
*/
 //var DSN string = "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	var	dsn string = "root@tcp(127.0.0.1:3306)/ginblog?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Database connect failed with error: ", err)
	}

	// disalbe gingulartable
	// db.SingularTable(true)

	db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqldb, err := db.DB()

	if err != nil {
		fmt.Println("Get sql.db from gorm.db error: ", err)
	}

	// SetMaxIdleConns set most of idle connections in connection pool
	sqldb.SetMaxIdleConns(10)

	// SetMaxOpenConns set most connections
	sqldb.SetMaxOpenConns(100)

	// SetConnMaxLifetime set max time which a connect can be used again
	sqldb.SetConnMaxLifetime(10)

}
