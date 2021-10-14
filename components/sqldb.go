package components

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/daqnext/cli-config-template/cli"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
db_host
db_port
db_name
db_username
db_password
*/
func InitDB() (*gorm.DB, *sql.DB, error) {

	db_host, db_host_err := cli.AppToDO.ConfigJson.GetString("db_host")
	if db_host_err != nil {
		return nil, nil, errors.New("db_host [string] in config.json not defined," + db_host_err.Error())
	}

	db_port, db_port_err := cli.AppToDO.ConfigJson.GetInt("db_port")
	if db_port_err != nil {
		return nil, nil, errors.New("db_port [int] in config.json not defined," + db_port_err.Error())
	}

	db_name, db_name_err := cli.AppToDO.ConfigJson.GetString("db_name")
	if db_name_err != nil {
		return nil, nil, errors.New("db_name [string] in config.json not defined," + db_name_err.Error())
	}

	db_username, db_username_err := cli.AppToDO.ConfigJson.GetString("db_username")
	if db_username_err != nil {
		return nil, nil, errors.New("db_username [string] in config.json not defined," + db_username_err.Error())
	}

	db_password, db_password_err := cli.AppToDO.ConfigJson.GetString("db_password")
	if db_password_err != nil {
		return nil, nil, errors.New("db_password [string] in config.json not defined," + db_password_err.Error())
	}

	dsn := db_username + ":" + db_password + "@tcp(" + db_host + ":" + strconv.Itoa(db_port) + ")/" + db_name + "?charset=utf8mb4&loc=UTC"

	GormDB, erropen := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if erropen != nil {
		return nil, nil, erropen
	}

	sqlDB, errsql := GormDB.DB()
	if errsql != nil {
		return nil, nil, errsql
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)

	return GormDB, sqlDB, nil

}
