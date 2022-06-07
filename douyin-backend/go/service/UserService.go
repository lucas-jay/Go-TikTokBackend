package service

import (
	"douyin-backend/go/dao"
	"douyin-backend/go/entity"
)

/**
新建User信息
*/
func CreateUser(user *entity.User) (err error) {
	if err = dao.SqlSession.Create(user).Error; err != nil {
		return err
	}
	return
}

/**
获取user集合
*/
func GetAllUser() (userList []*entity.User, err error) {
	if err := dao.SqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

/**
根据id删除user
*/
func DeleteUserById(id string) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.User{}).Error
	return
}

/**
根据id查询用户User
*/
func GetUserById(id string) (user *entity.User, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

/**
根据id查询用户User
*/
func GetUserByName(name string) (user entity.User, err error) {
	if err = dao.SqlSession.Where("name=?", name).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return
}

/**
更新用户信息
*/
func UpdateUser(user *entity.User) (err error) {
	err = dao.SqlSession.Save(user).Error
	return
}
