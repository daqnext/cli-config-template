package components

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	localLog "github.com/daqnext/LocalLog/log"

	fj "github.com/daqnext/fastjson"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

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
func InitDB(localLogger *localLog.LocalLog, ConfigJson *fj.FastJson) (*gorm.DB, *sql.DB, error) {

	if localLogger == nil {
		return nil, nil, errors.New("localLogger is required")
	}

	db_host, db_host_err := ConfigJson.GetString("db_host")
	if db_host_err != nil {
		return nil, nil, errors.New("db_host [string] in config.json not defined," + db_host_err.Error())
	}

	db_port, db_port_err := ConfigJson.GetInt("db_port")
	if db_port_err != nil {
		return nil, nil, errors.New("db_port [int] in config.json not defined," + db_port_err.Error())
	}

	db_name, db_name_err := ConfigJson.GetString("db_name")
	if db_name_err != nil {
		return nil, nil, errors.New("db_name [string] in config.json not defined," + db_name_err.Error())
	}

	db_username, db_username_err := ConfigJson.GetString("db_username")
	if db_username_err != nil {
		return nil, nil, errors.New("db_username [string] in config.json not defined," + db_username_err.Error())
	}

	db_password, db_password_err := ConfigJson.GetString("db_password")
	if db_password_err != nil {
		return nil, nil, errors.New("db_password [string] in config.json not defined," + db_password_err.Error())
	}

	dsn := db_username + ":" + db_password + "@tcp(" + db_host + ":" + strconv.Itoa(db_port) + ")/" + db_name + "?charset=utf8mb4&loc=UTC"

	GormDB, errOpen := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: New_gormLocalLogger(localLogger),
	})

	if errOpen != nil {
		return nil, nil, errOpen
	}

	sqlDB, errsql := GormDB.DB()
	if errsql != nil {
		return nil, nil, errsql
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)

	return GormDB, sqlDB, nil

}

///////////////////////////

type gormLocalLogger struct {
	LocalLogger           *localLog.LocalLog
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

func New_gormLocalLogger(localLogger *localLog.LocalLog) *gormLocalLogger {
	return &gormLocalLogger{
		LocalLogger:           localLogger,
		SkipErrRecordNotFound: true,
	}
}

func (l *gormLocalLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *gormLocalLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {

	//when no err
	if err == nil {
		//return when only interested in Error,Fatal,Panic
		if l.LocalLogger.Level < localLog.LLEVEL_WARN {
			return
		}
	}

	elapsed := time.Since(begin)
	if err == nil && l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		//slow log
		sql, _ := fc()
		fields := localLog.Fields{}
		if l.SourceField != "" {
			fields[l.SourceField] = utils.FileWithLineNum()
		}
		l.LocalLogger.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	///errors , when error happens logs it at any loglevel
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		sql, _ := fc()
		fields := localLog.Fields{}
		if l.SourceField != "" {
			fields[l.SourceField] = utils.FileWithLineNum()
		}
		fields[localLog.ErrorKey] = err
		l.LocalLogger.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	//info
	if l.LocalLogger.Level >= localLog.LLEVEL_DEBUG {
		sql, _ := fc()
		fields := localLog.Fields{}
		if l.SourceField != "" {
			fields[l.SourceField] = utils.FileWithLineNum()
		}
		if l.LocalLogger.Level == localLog.LLEVEL_DEBUG {
			l.LocalLogger.WithContext(ctx).WithFields(fields).Debugln("%s [%s]", sql, elapsed)
		} else {
			l.LocalLogger.WithContext(ctx).WithFields(fields).Traceln("%s [%s]", sql, elapsed)
		}

	}

}

func (l *gormLocalLogger) Info(ctx context.Context, s string, args ...interface{}) {
	//not used
}

func (l *gormLocalLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	//not used
}

func (l *gormLocalLogger) Error(ctx context.Context, s string, args ...interface{}) {
	//not used
}
