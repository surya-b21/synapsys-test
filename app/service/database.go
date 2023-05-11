package service

import (
	"fmt"
	"log"
	"os"
	"time"

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
			Logger: logger.New(
				log.New(os.Stderr, "[GORM] ", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             time.Second,   // Slow SQL threshold
					LogLevel:                  logger.Silent, // Log level
					IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
					Colorful:                  true,          // Disable color
				},
			),
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
		}

		if db, err := gorm.Open(conn, config); err == nil {
			if viper.GetBool("ENABLE_MIGRATION") {
				AutoMigrate(db)

				SeedAll(db)
			}

			DB = db.Debug()
		}
	}
}
