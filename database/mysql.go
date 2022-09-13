package database

import (
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var engine *xorm.Engine

/**
 * 连接数据库
 */
func InitMysql() *xorm.Engine {
	if engine == nil {
		engine, _ = xorm.NewEngine("sqlite3", "./database/test_focus.db")
	}
	return engine
}
