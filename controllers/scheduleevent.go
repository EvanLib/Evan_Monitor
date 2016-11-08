package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/evanlib/evan_monitor/models"

	"github.com/julienschmidt/httprouter"
)

type ScheduleEventsController struct {
	Controller
}

var ScheculeEvents ScheduleEventsController

func (self ScheduleEventsController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	schevents, err := models.GetAllScheduleEvents()
	if err != nil {
		fmt.Fprint(w, "Error retrieving scheduled events.")
		return err
	}

	tpl, err := template.ParseFiles("views/header_footer.html", "views/schedulesevents.html", "views/navbar.html")

	if err != nil {
		fmt.Println(err)
		return err
	}
	tpl.Execute(w, schevents)

	return nil
}

func (self ScheduleEventsController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {

	var buffer bytes.Buffer //I need to find a better way of doing this...
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

	fmt.Println(buffer.String())
	fmt.Println(r.FormValue("event-name"))
	fmt.Println(r.FormValue("event-description"))
	fmt.Println(r.FormValue("event-sun"))
	fmt.Println(r.FormValue("event-mon"))
	fmt.Println(r.FormValue("event-tues"))
	fmt.Println(r.FormValue("event-wed"))
	fmt.Println(r.FormValue("event-thurs"))
	fmt.Println(r.FormValue("event-fri"))
	fmt.Println(r.FormValue("event-sat"))
	fmt.Println(r.FormValue("event-sat"))
	fmt.Println(r.FormValue("event-startdate"))
	fmt.Println(r.FormValue("event-hour"))
	fmt.Println(r.FormValue("event-minute"))
	fmt.Println(r.FormValue("event-hourdur"))
	fmt.Println(r.FormValue("event-minutedur"))

	return nil
}
