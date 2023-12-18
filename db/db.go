package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"projecttimer/config"
	"projecttimer/utils"
)

var DB *gorm.DB

var conf *config.Config

func Create(config *config.Config) {
	if DB == nil {
		conf = config
		db, err := gorm.Open(sqlite.Open(config.Database.Name), &gorm.Config{
			NowFunc: utils.Local,
			Logger:  logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		DB = db
		createUserTab(config.App.User, config.App.Pwd)
		createCustomerTab()
	}
}

func createUserTab(userName string, pwd string) {
	user := User{Number: userName, Pwd: pwd}
	if !DB.Migrator().HasTable(&user) {
		DB.AutoMigrate(&user)
		DB.Create(&user)
	}
}
func createCustomerTab() {
	dst := &Customer{}
	if !DB.Migrator().HasTable(dst) {
		err := DB.AutoMigrate(dst)
		utils.Log.Println(err)
	}
}

// Db returns the default *gorm.DB connection.
func Db() *gorm.DB {
	return DB
}

// UnscopedDb returns an unscoped *gorm.DB connection
// that returns all records including deleted records.
func UnscopedDb() *gorm.DB {
	return Db().Unscoped()
}
