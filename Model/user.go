package Model

import "time"

const (
	USER_WOMAN = "W"
	USER_MAN = "M"
	USER_UNKNOWN = "U"
	)

type User struct {
	Id       int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`
	Mobile   string    `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd   string    `xorm:"varchar(40)" form:"passwd" json:"-"` // 什么角色
	Avatar   string    `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex      string    `xorm:"varchar(2)" form:"sex" json:"sex"`            // 什么角色
	Nickname string    `xorm:"varchar(20)" form:"nickname" json:"nickname"` // 什么角色
	Salt     string    `xorm:"varchar(10)" form:"salt" json:"-"`            // 什么角色
	Online   int       `xorm:"int(10)" form:"online" json:"online"`         //是否在线
	Token    string    `xorm:"varchar(40)" form:"token" json:"token"`       // 什么角色
	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`        // 什么角色
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`    // 什么角色
}
