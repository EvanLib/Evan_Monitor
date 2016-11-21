package models

import "github.com/astaxie/beego/orm"

//Holy Fuck I love ORM
type RepeatedEvent struct {
	Id            int    `orm:"auto"`
	Event         *Event `orm:"rel(one)"`
	Days          string `json:"days"`
	TimeStart     string
	TimeDuraction int //Int * minutes builds duration

}

func (rev *RepeatedEvent) Save() error {
	o := orm.NewOrm()

	_, err := o.Insert(rev)
	return err
}

func GetAllRepeatEvents() ([]*RepeatedEvent, error) {
	o := orm.NewOrm()
	var allRepeatEvents []*RepeatedEvent
	_, err := o.QueryTable("repeated_event").RelatedSel().All(&allRepeatEvents)
	if err != nil {
		return nil, err
	}
	return allRepeatEvents, nil
}

func GetRepeatEventById(newid int) (rev RepeatedEvent, err error) {
	o := orm.NewOrm()

	revent := RepeatedEvent{Id: newid}
	error := o.Read(&revent)

	return revent, error
}
