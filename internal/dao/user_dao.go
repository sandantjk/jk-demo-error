package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"jk-demo-error/internal/model"
)

// GetUserById 根据用户ID查询用户信息
// sql.ErrNoRows 此查询方法无论查询到用户信息或者查询不到，都是查询结果，不应该是异常，因此此方法需要处理掉此error
// 对于其他错误，一般是写法有bug或者db异常引起的，因此需要wrap之后抛给调用方处理
func GetUserById(db *sql.DB, id int64) (*model.User, error) {
	row := db.QueryRow("SELECT id, name, age FROM func_base.t_user where id=?", id)
	user := &model.User{}
	err := row.Scan(user.Id, user.Name, user.Age)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return user, errors.Wrap(err, "get user by id fail")
}
