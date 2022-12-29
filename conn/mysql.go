package conn

import (
	"log"
	"time"

	"img-svc/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type gormMysql struct {
	Db *gorm.DB
}

var Client gormMysql

func ConnectDB() error {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	dsn := domain.DbUserName + ":" + domain.DbPass + "@tcp(" + domain.DbHost + ":" + domain.DbPort + ")/" + domain.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Println(err.Error())
		return err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(domain.DbMaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(domain.DbMaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(domain.DbConnMaxLifetime) * time.Hour)

	Client.Db = db

	return nil
}
