package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username      string `json:"username"`
	Password      string `json:"password"`
	FollowCount   uint   `json:"follow_count"`
	FollowerCount uint   `json:"follower_count"`
	//Following []User `gorm:"many2many:relationships;association_joint able_foreigner:follow_to"`
}

//func (u *User) Follow(followTo User) (err error) {
//	err = DB.Model(u).Association("Following").Append(&followTo).Error
//	if err != nil {
//		err = errors.New("Follow failed")
//	}
//	return
//}
