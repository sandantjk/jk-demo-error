package conf

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

func InitConfig(configFile string) (*Config, error) {
	config := new(Config)
	err := unMarshalYamlFile(configFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func InitDB(dbConf DBConfig) (*sql.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", dbConf.UserName, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DbName)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func unMarshalYamlFile(yamlPath string, out interface{}) error {
	file, err := os.Open(yamlPath)
	if err != nil {
		return err
	}
	defer func(fio *os.File) {
		_ = fio.Close()
	}(file)

	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(all, out)
	if err != nil {
		return err
	}
	return nil
}
