package config

import (
	"database/sql"

	"github.com/ovitorhilario/gopher/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializatePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	// dbPath := "./db/main.db"

	// _, err := os.Stat(dbPath)
	// if os.IsNotExist(err) {
	// 	logger.Info("database file not found, creating...")
	// 	// Create the directory
	// 	err = os.MkdirAll("./db", os.ModePerm)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	file, err := os.Create(dbPath)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	file.Close()
	// }

	// Get local database
	dsn := "host=localhost user=postgres password=postgres dbname=gopher port=5432 sslmode=disable"
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Errorf("postgres opening error: %v", err)
		return nil, err 
	}

	// Opening the postgres with connection
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		logger.Errorf("gorm opening postgres-sql database error: %v", err)
		return nil, err 
	}
	
	// Migrate the Schema
	err = gormDB.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("postgres auto-migration gorm error %v", err)
		return nil, err
	}

	return gormDB, nil
}