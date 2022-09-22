package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"bili-monitor-system/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() error {
	var err error
	DB, err = connect()
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&Video{})
	if err != nil {
		return err
	}
	return nil
}

func connect() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level [Silent Error Warn Info]
			Colorful:      true,          // 禁用彩色打印
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlConf.Username,
		config.MysqlConf.Password,
		config.MysqlConf.Address,
		config.MysqlConf.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
