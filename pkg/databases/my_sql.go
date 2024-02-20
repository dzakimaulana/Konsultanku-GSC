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

func DatabaseConn() (*Database, error) {
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
		return nil, err
	}

	if err := database.AutoMigrate(
		&models.Tags{},
		&models.MSME{},
		&models.Student{},
		&models.Problem{},
		&models.Comment{},
		&models.Team{},
		&models.Collaboration{},
		&models.ProblemsTags{},
		&models.UsersTags{},
	); err != nil {
		return nil, err
	}
	return &Database{DB: database}, nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.DB
}
