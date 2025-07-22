package pkg

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	mysqlOnce sync.Once
	mysqlDB   *sql.DB
	mysqlErr  error
)

// GetMySQL 返回全局 MySQL 连接
func GetMySQL() (*sql.DB, error) {
	mysqlOnce.Do(func() {
		host := viper.GetString("mysql.host")
		port := viper.GetInt("mysql.port")
		user := viper.GetString("mysql.user")
		password := viper.GetString("mysql.password")
		dbname := viper.GetString("mysql.dbname")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local",
			user, password, host, port, dbname)
		mysqlDB, mysqlErr = sql.Open("mysql", dsn)
		if mysqlErr != nil {
			return
		}
		// 连接测试
		mysqlErr = mysqlDB.Ping()
	})
	return mysqlDB, mysqlErr
}
