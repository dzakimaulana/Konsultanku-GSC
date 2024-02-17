package databases

import (
	"fmt"
	"konsultanku-v2/pkg/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func DatabaseConn() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(".env not detected")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))
	database, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err.Error())
	}

	if err := database.AutoMigrate(
		&models.Problem{},
	); err != nil {
		panic(err.Error())
	}
}

func (d *Database) GetDB() *gorm.DB {
	return d.DB
}
