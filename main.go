package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/evanlib/evan_monitor/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

//Global Variable for templates
var t *template.Template

func main() {
	//Template parsing
	t = template.Must(template.ParseGlob("views/*.html"))
	//Set Up Database Handler
	orm.RegisterDataBase("default", "mysql", "root:lol626465@/me_events?charset=utf8")

	//Create Controllers and Models
	eventsCon := controllers.Events

	//Http Router
	router := httprouter.New()
	router.GET("/events", eventsCon.Perform(eventsCon.Index))
	router.GET("/events/create", eventsCon.Perform(eventsCon.Index))

	//Serving Static Files...
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	//Log for Http server
	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
