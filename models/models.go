package models

import "github.com/astaxie/beego/orm"

//Import for ORM use

func init() {
	//register models
	orm.RegisterModel(new(Event))
	orm.RegisterModel(new(RepeatedEvent))
	orm.RegisterModel(new(OneTimeEvent))
	orm.RegisterModel(new(Schedule))

}
