package models

import "github.com/astaxie/beego/orm"

//Holy Fuck I love ORM
type RepeatedEvent struct {
	Id            int    `orm:"auto"`
	Event         *Event `orm:"rel(one)"`
	Days          string
	TimeStart     string
	TimeDuraction int //Int * minutes builds duration

}

func (rev *RepeatedEvent) Save() error {
	o := orm.NewOrm()

	_, err := o.Insert(rev)
	return err
}

func GetAllRepeatEvents() ([]RepeatedEvent, error) {
	o := orm.NewOrm()
	evs := o.QueryTable("repeated_event")

	var allRepeatEvents []RepeatedEvent
	evs.All(&allRepeatEvents)

	return allRepeatEvents, nil
}

func GetRepeatEventById(newid int) (rev RepeatedEvent, err error) {
	o := orm.NewOrm()

	revent := RepeatedEvent{Id: newid}
	error := o.Read(&revent)

	return revent, error
}
