package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/evanlib/evan_monitor/models"

	"github.com/julienschmidt/httprouter"
)

type EventsController struct {
	Controller
}

var Events EventsController

func (self EventsController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Fprint(w, "Error retrieving events")
		return err
	}

	tpl, err := template.ParseFiles("views/header_footer.html", "views/navbar.html", "views/events.html")
	if err != nil {
		return err
	}
	tpl.Execute(w, events)

	//fmt.Print(events)
	return nil
}

func (self EventsController) CreateRepeatEvent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {

	var eventHourDur = r.FormValue("event-hourdur")
	var eventMinuteDur = r.FormValue("event-minutedur")
	var eventHour = r.FormValue("event-hour")
	var eventMinute = r.FormValue("event-minute")

	//Find a better way
	duration, err := strconv.Atoi(eventHourDur)
	duration = duration * 60
	durationMin, err := strconv.Atoi(eventMinuteDur)
	duration = duration + durationMin

	//Find a better way
	var buffer bytes.Buffer
	if r.FormValue("event-sun") == "on" {
		buffer.WriteString("1")
	} else {
		buffer.WriteString("0")
	}
	if r.FormValue("event-mon") == "on" {
		buffer.WriteString("1")
	} else {
		buffer.WriteString("0")
	}
	if r.FormValue("event-tues") == "on" {
		buffer.WriteString("1")
	} else {
		buffer.WriteString("0")
	}
	if r.FormValue("event-wed") == "on" {
		buffer.WriteString("1")
	} else {
		buffer.WriteString("0")
	}
	if r.FormValue("event-thurs") == "on" {
		buffer.WriteString("1")
	} else {
		buffer.WriteString("0")
	}
	if r.FormValue("event-fri") == "on" {
		buffer.WriteString("1")
	} else {
		buffer.WriteString("0")
	}
	if r.FormValue("event-sat") == "on" {
		buffer.WriteString("1")
	} else {
		buffer.WriteString("0")
	}
	//Create event
	ev := models.Event{
		Type:        2,
		Name:        r.FormValue("event-name"),
		Description: r.FormValue("event-description"),
	}
	//Create RepeatedEvent
	rev := models.RepeatedEvent{
		Event:         &ev,
		Days:          buffer.String(),
		TimeStart:     eventHour + ":" + eventMinute,
		TimeDuraction: duration,
	}

	fmt.Printf("Repeated Event: %s created. \n", ev.Name)

	//Insert into DB
	err = ev.Save()
	err = rev.Save()

	if err != nil {
		return err
	}
	return nil
}
