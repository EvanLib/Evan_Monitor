package models

type Schdule struct {
	Id     int
	Name   string
	Events []ScheduleEvent
	Active bool
}
