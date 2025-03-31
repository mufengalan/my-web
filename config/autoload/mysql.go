package autoload

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MysqlDB *gorm.DB

type MysqlConfig struct {
	Host         string `default:"localhost" mapstructure:"host"`
	Port         string `default:"3306" mapstructure:"port"`
	DBName       string `default:"local" mapstructure:"db_name"`
	Username     string `default:"root" mapstructure:"username"`
	Password     string `default:"123456" mapstructure:"password"`
	Charset      string `default:"utf8mb4" mapstructure:"charset"`
	MaxIdleConns int    `default:"10" mapstructure:"max_idle_conns"`
	MaxOpenConns int    `default:"100" mapstructure:"max_open_conns"`
	MaxLifetime  int    `default:"3600" mapstructure:"max_lifetime"`
}

var mysqlConfig = MysqlConfig{
	DBName:  "local",
	Charset: "utf8mb4",
}

var GormDb *gorm.DB

func InitMysql(cnf MysqlConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cnf.Username,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		mysqlConfig.DBName,
		mysqlConfig.Charset,
	)

	var err error
	for i := 0; i < 3; i++ {
		GormDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
	if err != nil {
		return fmt.Errorf("failed to connect to database after 3 retries: %v", err)
	}

	sqlDB, err := GormDb.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(cnf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cnf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cnf.MaxLifetime) * time.Second)

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to database!")
	return nil
}
