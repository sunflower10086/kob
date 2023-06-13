// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameBot = "bot"

// Bot mapped from table <bot>
type Bot struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true;uniqueIndex:bot_id_uindex,priority:1" json:"id"`
	UserID      int32     `gorm:"column:user_id;not null" json:"user_id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Code        string    `gorm:"column:code" json:"code"`
	Createtime  time.Time `gorm:"column:createtime" json:"createtime"`
	Modifytime  time.Time `gorm:"column:modifytime" json:"modifytime"`
}

// TableName Bot's table name
func (*Bot) TableName() string {
	return TableNameBot
}