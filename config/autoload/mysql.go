package autoload

type MysqlConfig struct {
	Host     string `default:"localhost" mapstructure:"host"`
	Port     string `default:"3306" mapstructure:"port"`
	DBName   string `default:"local" mapstructure:"db_name"`
	Username string `default:"root" mapstructure:"username"`
	Password string `default:"123456" mapstructure:"password"`
	Charset  string `default:"utf8mb4" mapstructure:"charset"`
}
