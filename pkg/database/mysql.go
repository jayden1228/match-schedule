package database

import (
	"fmt"
	"log"
	"time"

	"match-schedule/pkg/configs"

	"github.com/jinzhu/gorm"

	// 引用数据库驱动初始化
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var engine *gorm.DB

func GetDB() *gorm.DB {
	return engine
}

func init() {
	var err error
	mysqlConf := configs.EnvConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		mysqlConf.User,
		mysqlConf.Pwd,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DBName,
		mysqlConf.Charset,
		true,
		"Local")

	engine, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Println("connect to mysql fail, ", dsn, err)
		panic(err)
	}

	logSQL := true
	if configs.EnvConfig.ProjectEnv == "prod" {
		logSQL = false
	}
	engine.LogMode(logSQL)

	engine.DB().SetConnMaxLifetime(100 * time.Second)
	engine.DB().SetMaxOpenConns(150)
	engine.DB().SetMaxIdleConns(100)

	// 注册创建/更新回调函数，自动插入时间戳
	engine.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	engine.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// 禁止update/delete传空对象
	engine.BlockGlobalUpdate(true)
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now().UTC().Unix()

		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if createdAtField.IsBlank {
				createdAtField.Set(now)
			}
		}

		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(now)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		now := time.Now().UTC().Unix()
		scope.SetColumn("UpdatedAt", now)
	}
}