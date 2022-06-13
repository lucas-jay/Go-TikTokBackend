package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type Following struct {
	gorm.Model
	HostId  uint
	GuestId uint
}

type Followers struct {
	gorm.Model
	HostId  uint
	GuestId uint
}
