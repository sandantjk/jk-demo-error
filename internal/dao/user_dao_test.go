package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"jk-demo-error/internal/conf"
	"testing"
)

func TestSelectUserById(t *testing.T) {
	db, err := getDB()
	if err != nil {
		t.Error(err)
		return
	}
	user, err := GetUserById(db, 1)
	if err != nil {
		t.Error(err)
		return
	}
	if user == nil {
		t.Log("user not found")
	} else {
		t.Log(user)
	}
}

func getDB() (*sql.DB, error) {
	config, err := conf.InitConfig("../../configs/application.yml")
	if err != nil {
		return nil, err
	}
	return conf.InitDB(config.DB)
}
