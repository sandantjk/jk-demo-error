package conf

import "database/sql"

type Config struct {
	ServerPort int      `yaml:"serverPort" json:"serverPort"`
	DB         DBConfig `yaml:"db" json:"db"`
}

type DBConfig struct {
	UserName string `yaml:"userName" json:"userName"`
	Password string `yaml:"password" json:"password"`
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	DbName   string `yaml:"dbName" json:"dbName"`
}

var (
	Global *Config
	DB *sql.DB
)


