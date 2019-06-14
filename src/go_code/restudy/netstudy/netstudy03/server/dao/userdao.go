package dao

import (
	"go_code/restudy/netstudy/netstudy03/server/db"
	"go_code/restudy/netstudy/netstudy03/common/model"
	"fmt"
)

type UserDao struct {

}

func (this *UserDao) Login(userId string) (user *model.User) {
	user = &model.User{}
	err := db.DB.QueryRow("select user_id, user_name, password from user where user_id = ?", userId).Scan(&user.UserId, &user.UserName, &user.Password)
	if err != nil {
		fmt.Println("查询用户数据失败")
		return
	}
	return 
}