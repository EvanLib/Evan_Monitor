package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Schedule struct {
	Id             int `orm:"auto"`
	Name           string
	RepeatedEvents []*RepeatedEvent `orm:"rel(m2m)"`
	StartDate      time.Time
	EndDate        time.Time
}

func (sch *Schedule) Save() error {
	o := orm.NewOrm()

	_, err := o.Insert(sch)
	return err
}

func GetAllSchedules() ([]Schedule, error) {
	o := orm.NewOrm()
	schs := o.QueryTable("schedule")

	var allSchedules []Schedule
	schs.All(&allSchedules)

	return allSchedules, nil
}
