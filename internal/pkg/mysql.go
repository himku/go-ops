package pkg

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqlOnce sync.Once
	mysqlDB   *gorm.DB
	mysqlErr  error
)

// GetDB 返回全局 GORM MySQL 连接
func GetDB() (*gorm.DB, error) {
	mysqlOnce.Do(func() {
		host := viper.GetString("mysql.host")
		port := viper.GetInt("mysql.port")
		user := viper.GetString("mysql.user")
		password := viper.GetString("mysql.password")
		dbname := viper.GetString("mysql.dbname")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
			user, password, host, port, dbname)
		mysqlDB, mysqlErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	})
	return mysqlDB, mysqlErr
}
