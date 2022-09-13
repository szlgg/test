package config

import "gopkg.in/ini.v1"

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	Port       int `ini:"port"`
	*Server    `ini:"server"`
	*Database  `ini:"database"`
	*Upload    `ini:"upload"`
	*Index    `ini:"index"`
}

// Index 前台配置
type Index struct {
	Num int `ini:"num"`
}

// Server 服务器配置
type Server struct {
	Address          string `ini:"address"`
	ServerRoot       string `ini:"serverRoot"`
	DumpRouterMap    string `ini:"dumpRouterMap"`
	RouteOverWrite   string `ini:"routeOverWrite"`
	AccessLogEnabled int    `ini:"accessLogEnabled"`
	AccessLogPattern int    `ini:"accessLogPattern"`
	SessionPath      int    `ini:"sessionPath"`
	LogPath          int    `ini:"logPath"`
}

// Database 数据库配置
type Database struct {
	*Logger  `ini:"database.logger"`
	*Default `ini:"database.default"`
}

// Logger 数据库日志配置
type Logger struct {
	Path   string `ini:"path"`
	Level  string `ini:"level"`
	Stdout bool   `ini:"stdout"`
}

// Default 数据库连接配置
type Default struct {
	Link  string `ini:"link"`
	Debug bool   `ini:"debug"`
}

type Upload struct {
	Path string `ini:"path"`
}

// Init 把先关初始的参数加载到全局变量，然后方便调用
func Init(file string) error {
	return ini.MapTo(Conf, file)
}
