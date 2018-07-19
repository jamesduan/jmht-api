package rpc

import (
	"database/sql"
	"jmht-api/common/model"
	"jmht-api/db/user"
	"log"
)

//CreateUser params: req 请求数据结构 类型为model.User 包含用户的信息
//res 和req拥有相同的数据结构 error: 返回值为异常类型
func (rpcuser *User) CreateUser(req *model.User, res *model.User) error {
	err := user.Create(req)
	if err != nil {
		log.Println(err)
	}
	res = nil
	return nil
}

//GetUser 使用id获取单个用户 params: userid it's belong to model.User id field
// res is a response equals user object
func (rpcuser *User) GetUser(userID int, res *model.User) error {
	err := user.SelectUserById(userID, res)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
	return nil
}

//UpdateUser 更新用户 params: req, res
func (rpcuser *User) UpdateUser(req *model.User, res *sql.Result) error {

	err := user.Update(req)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (rpcuser *User) DeleteUser(id int, res *sql.Result) error {
	err := user.Delete(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
