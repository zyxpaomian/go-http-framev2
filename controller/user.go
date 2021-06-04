package controller

import (
	"go-http-frame/dao"
	"go-http-frame/structs"
	se "go-http-frame/common/error"
	"github.com/satori/go.uuid"
	log "go-http-frame/common/formatlog"
)

var UserController *UserCtrl

type UserCtrl struct {
	userDao *dao.UserDao
}

func init() {
	UserController = &UserCtrl{
		userDao: &dao.UserDao{},
	}
}

// 获取所有用户信息
func (u *UserCtrl) GetAllUsers() ([]*structs.User, error) {
	return u.userDao.ListAllUsers()
}

// 用户认证并获取token
func (u *UserCtrl) GetToken(username, password string) (string, error) {
	_, err := u.userDao.UserAuth(username, password)
	if err != nil {
		log.Errorln("用户认证失败，无法获取token")
		return "", se.AuthError()
	}
	token := uuid.NewV4()
	tokenStr := token.String()
	err = u.userDao.TokenSave(tokenStr, username)
	if err != nil {
		log.Errorf("生成token失败, 失败原因: %v", err.Error())
		return "",se.DBError()
	}
	return tokenStr, nil
}

// token认证
func (u *UserCtrl) GetTokenUser(token string) (string, error) {
	username, err := u.userDao.TokenAuth(token)
	if err != nil {
		log.Errorln("用户token认证失败")
		return "", se.AuthError()
	} else {
		return username, nil
	}
}




