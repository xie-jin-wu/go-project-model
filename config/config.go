package config

import "go-project-model/pkg/logs"

type Config struct {
	HttpAddress       string        `yaml:"http_address"`         //http地址
	PProfAddress      string        `yaml:"pprof_address"`        //pprof地址
	ProgramLogLevel   logs.LogLevel `yaml:"program_log_level"`    //程序日志等级
	ProgramLogToFile  bool          `yaml:"program_log_to_file"`  //程序日志输出到文件(true:文件,false:终端)
	ProgramLogFileDir string        `yaml:"program_log_file_dir"` //程序日志输出到文件目录名
	HttpLogToFile     bool          `yaml:"http_log_to_file"`     //http日志输出到文件(true:文件,false:终端)
	HttpLogFileDir    string        `yaml:"http_log_file_dir"`    //http日志输出到文件目录名
	Mysql             struct {
		Dns   string `yaml:"dns"`
		Debug bool   `yaml:"debug"`
	} `yaml:"mysql"`
	Redis struct {
		IP       string `yaml:"ip"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
	Mongo struct {
		Address  string `yaml:"address"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"mongo"`
}
