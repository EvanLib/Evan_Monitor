package models

import "time"

//Holy Fuck I love ORM
type OneTimeEvent struct {
	Id            int    `orm:"auto"`
	Event         *Event `orm:"rel(one)"`
	DateStart     time.Time
	TimeDuraction int //Int * minutes builds duration

}
