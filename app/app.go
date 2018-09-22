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
	DB.db.AutoMigrate(&Flag{})
	DB.db.AutoMigrate(&PointsAndAccess{})
	DB.db.AutoMigrate(&PresentCTFValue{})
	DB.db.AutoMigrate(&McqDetail{})

	router := httprouter.New()

	router.GET("/api/", Index)
	router.POST("/api/login", Login)
	router.GET("/api/logout", Logout)
	router.GET("/api/mcq/:idx", Mcq)
	router.POST("/api/submitFlag", SubmitFlag)

	router.GET("/api/admin", AdminIndex)
	router.GET("/api/admin/setPresentCTFValues", SetPresentCTFValues)
	router.POST("/api/admin/addUser", AddUser)

	log.Println("Server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
