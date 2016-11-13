package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/EvanLib/evan_monitor/controllers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

//Global Variable for templates
var t *template.Template

func main() {
	//Template parsing
	t = template.Must(template.ParseGlob("views/*.html"))
	//Set Up Database Handler
	orm.RegisterDataBase("default", "mysql", "root:lol626465@/me_schedule?charset=utf8")
	//Create controllers
	evscontroller := controllers.Events
	schscontroller := controllers.Schedules
	//Http Router
	router := httprouter.New()
	router.GET("/events", evscontroller.Perform(evscontroller.Index))
	router.POST("/events/createRepeat", evscontroller.Perform(evscontroller.CreateRepeatEvent))

	router.GET("/schedules", schscontroller.Perform(schscontroller.Index))
	router.GET("/schedules/addevents", schscontroller.Perform(schscontroller.ScheduleEvents))
	router.POST("/schedule/create", schscontroller.Perform(schscontroller.Create))

	//Serving Static Files...
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	//Log for Http server
	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
