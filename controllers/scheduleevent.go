package controllers

import (
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
		return
	}

	tpl, err := template.ParseFiles("views/header_footer.html", "views/scheduleevent.html")

	if err != nil {
		return err
	}

	tpl.Execute(w, schevents)

	return nil
}
