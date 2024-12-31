package autoload

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

type MysqlConfig struct {
	Host     string `default:"localhost" mapstructure:"host"`
	Port     string `default:"3306" mapstructure:"port"`
	DBName   string `default:"local" mapstructure:"db_name"`
	Username string `default:"root" mapstructure:"username"`
	Password string `default:"123456" mapstructure:"password"`
	Charset  string `default:"utf8mb4" mapstructure:"charset"`
}

var Mysql = MysqlConfig{
	Host:     "localhost",
	Port:     "3306",
	DBName:   "local",
	Username: "root",
	Password: "123456",
	Charset:  "utf8mb4",
}

var GormDb *gorm.DB

func InitMysql() {
	//连接mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		Mysql.Username,
		Mysql.Password,
		Mysql.Host,
		Mysql.Port,
		Mysql.DBName,
		Mysql.Charset,
	)
	configs := &gorm.Config{}
	var err error
	GormDb, err = gorm.Open(mysql.Open(dsn), configs)
	if err != nil {
		panic("mysql connect error:" + err.Error())
	}
	sqlDB, err := GormDb.DB()
	if err != nil {
		fmt.Println("获取数据库实例失败:", err)
		return
	}

	if err := sqlDB.Ping(); err != nil {
		fmt.Println("连接数据库失败:", err)
		return
	}
	fmt.Println("成功连接到数据库！")

}
