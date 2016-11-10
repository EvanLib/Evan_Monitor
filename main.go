package main

import (
	"fmt"
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
	//Http Router
	router := httprouter.New()
	router.GET("/events", evscontroller.Perform(evscontroller.Index))

	//Serving Static Files...
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	//Log for Http server
	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
