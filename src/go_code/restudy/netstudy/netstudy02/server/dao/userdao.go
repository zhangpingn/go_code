package dao

import (
	"go_code/restudy/netstudy/netstudy02/common/model"
	"go_code/restudy/netstudy/netstudy02/server/db"
	"fmt"
)
type UserDao struct {

}

func (userDao *UserDao) GetUserByUserId(userId string) (user *model.User, err error){
	user = &model.User{}
	err = db.DB.QueryRow("select user_id, user_name, password, user_status from user where user_id = ? and user_status != 'A'", userId).Scan(&user.UserId, &user.UserName, &user.Password, &user.UserStatus)
	if err != nil {
		fmt.Println("查询数据失败", err)
		return
	}
	return
}

func (userDao *UserDao) GetAllOnlineUser(userId string) (users []model.User, err error){
	rows, err := db.DB.Query("select user_id, user_name, user_status from user where user_status = 'A' and user_id != ?", userId)
	if err != nil {
		fmt.Println("获取在线用户失败")
		return
	}
	defer rows.Close()
	for rows.Next() { //每次循环会取下一行数据
		var user model.User 
		rows.Scan(&user.UserId, &user.UserName, &user.UserStatus)
		users = append(users, user)
	}
	return
}
func (userDao *UserDao) UpdateStatus(userId string, userStatus string) (int64, error){
	rs, _ := db.DB.Exec("update user set user_status = ? where user_id = ?", userStatus, userId)
	return rs.RowsAffected();
}

func (userDao *UserDao) AddUser(user *model.User) (int64, error){
	rs, _ := db.DB.Exec("insert into user values(?, ?, ?, ?)", user.UserId, user.UserName, user.Password, user.UserStatus) 
	return rs.RowsAffected()
}