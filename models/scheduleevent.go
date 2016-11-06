package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type ScheduleEvent struct {
	Id           int  `orm:"auto"` //Auto Increase Id
	Type         int8 // Repeated or Not
	Name         string
	Description  string
	TimeStart    time.Time
	TimeDuration int    //TimeDuration * minute
	Days         string //0-6 0 Sunday, 1, Monday ...
	Date         time.Time
}

func (se *ScheduleEvent) Save() error {
	//ORM
	o := orm.NewOrm()

	//Insert
	_, err := o.Insert(se)
	return err
}

func GetAllScheduleEvents() ([]ScheduleEvent, error) {
	//ORM
	o := orm.NewOrm()
	sev := o.QueryTable("scheduleevent")

	var allScheduleEvents []ScheduleEvent
	sev.All(&allScheduleEvents)
	return allScheduleEvents, nil
}

func GetScheduleEventByID(newid int) (ScheduleEvent, error) {
	//ORM
	o := orm.NewOrm()

	//Create ScheduleEvent Scaffold
	schev := ScheduleEvent{Id: newid}
	err := o.Read(&schev)

	return schev, err
}
