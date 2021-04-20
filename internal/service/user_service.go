package service

import (
	"fmt"
	"jk-demo-error/internal/conf"
	"jk-demo-error/internal/dao"
)
// PrintUserInfo 打印用户信息
// 处理业务逻辑，将非业务错误返回给框架处理
func PrintUserInfo(id int64) error {
	user, err := dao.GetUserById(conf.DB, id)
	if err != nil {
		return err
	}
	if user == nil {
		fmt.Println(fmt.Sprintf("ID为{%d}的用户不存在", id))
	} else {
		fmt.Println(fmt.Sprintf("ID为{%d}的用户为: %s %d", user.Id, user.Name, user.Age))
	}
	return nil
}
