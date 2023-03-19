package db

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
)

var Db *gorm.DB

// Init 初始化数据库
func Init() error {
	source := "%s:%s@tcp(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	var user, pwd, addr, dataBase string
	initConfig()
	user = viper.GetString("MYSQL_USERNAME")
	pwd = viper.GetString("MYSQL_PASSWORD")
	addr = viper.GetString("MYSQL_ADDRESS")
	dataBase = viper.GetString("MYSQL_DATABASE")
	if dataBase == "" {
		dataBase = "golang_demo"
	}
	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	fmt.Println("start init mysql with ", source)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		}})
	if err != nil {
		fmt.Println("DB Open error,err=", err.Error())
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("DB Init error,err=", err.Error())
		return err
	}

	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(200)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	// db.AutoMigrate(&model.Question{})
	// db.AutoMigrate(&model.QuestionOption{})
	// db.AutoMigrate(&model.User{})
	Db = db

	fmt.Println("finish init mysql")
	return nil
}

func initConfig() {
	// viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigName("config_local") // name of config file (without extension)
	viper.SetConfigType("yaml")         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/")    // path to look for the config file in
	err := viper.ReadInConfig()         // Find and read the config file
	if err != nil {                     // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
