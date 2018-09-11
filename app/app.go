package app

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

/*
** migrates and creates tables in data base
** registers functions with routes
** opens up port and serves the app
 */

var err error
var DB Database

func Run() {

	defer DB.db.Close()

	DB.db.AutoMigrate(&User{})
	DB.db.AutoMigrate(&Admin{})

	router := httprouter.New()

	router.GET("/", Index)
	router.POST("/login", Login)
	router.GET("/logout", Logout)

	router.GET("/admin", AdminIndex)

	log.Println("Server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
