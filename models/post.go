package models

import "time"

type Post struct {
	Id      int64
	Title   string
	Body    string `xorm:"text"`
	Image   string
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}
