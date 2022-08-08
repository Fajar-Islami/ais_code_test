package entity

import "time"

type Article struct {
	ID      uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Author  string    `gorm:"type:TEXT;not null" json:"author"`
	Title   string    `gorm:"type:TEXT;not null" json:"title"`
	Body    string    `gorm:"type:TEXT;not null" json:"body"`
	Created time.Time `gorm:"autoCreateTime:true" json:"created"`
	Search  string    `gorm:"-" json:"-"`
	Limit   uint      `gorm:"-" json:"-"`
	Page    uint      `gorm:"-" json:"-"`
}
