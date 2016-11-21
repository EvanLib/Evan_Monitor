package models

import "github.com/astaxie/beego/orm"

type Event struct {
	Id          int  `orm:"auto"`
	Type        int8 //Positive: 1 Negative: -1 Neutral: 0
	Name        string
	Description string
}

func (ev *Event) Save() error {
	//Create ORM
	o := orm.NewOrm()
	//Insert event
	_, err := o.Insert(ev)
	return err
}

func GetAllEvents() ([]Event, error) {
	o := orm.NewOrm()
	evs := o.QueryTable("event") // Returns a number of elements and an error

	var allEvents []Event
	evs.All(&allEvents)
	return allEvents, nil
}

func GetEventByID(newid int) (ev Event, err error) {
	//Create orm
	o := orm.NewOrm()

	//Create Event
	event := Event{Id: newid}
	error := o.Read(&event)

	return event, error
}
