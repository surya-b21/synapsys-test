package service

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB global variable
var DB *gorm.DB

// InitDB to initialize connection to db
func InitDB() {
	// read all env variable
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	if DB == nil {
		conn := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", viper.GetString("DB_USERNAME"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_HOST"), viper.GetString("DB_PORT"), viper.GetString("DB_DATABASE")))
		config := &gorm.Config{
			Logger:         logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
		}

		db, err := gorm.Open(conn, config)

		if err == nil {
			if viper.GetBool("ENABLE_MIGRATION") {
				AutoMigrate(db)

				SeedAll(db)
			}
			DB = db.Debug()
		} else {
			log.Panic(err)
		}
	}
}

// logger.New(
// 	log.New(os.Stdout, "[GORM] ", log.LstdFlags), // io writer
// 	logger.Config{
// 		SlowThreshold:             time.Second,   // Slow SQL threshold
// 		LogLevel:                  logger.Silent, // Log level
// 		IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
// 		Colorful:                  true,          // Disable color
// 	},
// ),
