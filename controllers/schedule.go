package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/evanlib/evan_monitor/models"
	"github.com/julienschmidt/httprouter"
)

type SchedulesController struct {
	Controller
}

var Schedules SchedulesController

func (self SchedulesController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	schedules, err := models.GetAllSchedules()
	if err != nil {
		fmt.Fprint(w, "Error retrieving events")
		return err
	}

	tpl, err := template.ParseFiles("views/header_footer.html", "views/navbar.html", "views/schedulesevents.html")
	if err != nil {
		return err
	}
	tpl.Execute(w, schedules)

	//fmt.Print(events)
	return nil
}

func (self SchedulesController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	var scheduleName = r.FormValue("schedule-name")
	var scheduleStartDate = r.FormValue("schedule-startDate")
	var scheduleEndDate = r.FormValue("schedule-endDate")

	const shortForm = "01/02/2006"
	startt, _ := time.Parse(shortForm, scheduleStartDate)
	endt, _ := time.Parse(shortForm, scheduleEndDate)

	sch := models.Schedule{
		Name:      scheduleName,
		StartDate: startt,
		EndDate:   endt,
	}

	//Save New Schedule
	err := sch.Save()
	if err != nil {
		fmt.Printf("Error in creating schedule: %s \n", err)
		return err
	}
	return nil
}

func (self SchedulesController) ScheduleEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	revents, err := models.GetAllRepeatEvents()
	for _, el := range revents {
		fmt.Println(el.Days)
		fmt.Println(el.Event.Name)
	}
	if err != nil {
		fmt.Printf("Error in retrieving repated events: %s \n", err)
		return err
	}

	tpl, err := template.ParseFiles("views/header_footer.html", "views/navbar.html", "views/scheduleaddevents.html")
	if err != nil {
		fmt.Printf("Error in Parsing html files: %s \n", err)
		return err
	}
	tpl.Execute(w, revents)

	return nil
}

func (self SchedulesController) ScheduleListEvents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error {
	revents, err := models.GetAllRepeatEvents()

	if err != nil {
		fmt.Printf("Error in Doing Stuff %s", err)
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(revents)

	return nil
}
