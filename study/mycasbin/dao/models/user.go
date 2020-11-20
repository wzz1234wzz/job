package models

import "time"

// User
type User struct {
	Ysid        int64     `json:"ysid" xorm:"not null pk autoincr BIGINT(20)"`
	Name        string    `json:"name" xorm:"not null VARCHAR(128) UNIQUE"`
	PwdHash     string    `json:"pwd_hash" xorm:"VARCHAR(128)"`
	Group       string    `json "group" xorm:"varchar(20) index not null default ''"`
	Rank        int       `json "rank" xorm:"default 1 TINYINT(1)"`
	Email       string    `json:"email" xorm:"VARCHAR(128)"`
	Phone       string    `json:"phone" xorm:"VARCHAR(16)"`
	CreatedAt   time.Time `json:"created_at" xorm:"default CURRENT_TIMESTAMP created TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" xorm:"default CURRENT_TIMESTAMP updated TIMESTAMP"`
	IsActivated int       `json:"is_activated" xorm:"default 1 TINYINT(1)"`
}


