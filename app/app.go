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

	router.GET("/", Index)
	router.POST("/login", Login)
	router.GET("/mcq/:idx", Mcq)
	router.POST("/submitFlag/mcq/:question_id", CheckFlagMcq)

	// router.GET("/api/", ApiIndex)
	// router.POST("/api/login", ApiLogin)
	// router.GET("/api/logout", ApiLogout)
	// router.GET("/api/mcq/:idx", ApiMcq)
	// router.POST("/api/submitFlag", ApiSubmitFlag)

	// router.GET("/api/admin", AdminIndex)
	// router.GET("/api/admin/setPresentCTFValues", SetPresentCTFValues)
	// router.POST("/api/admin/addUser", AddUser)

	// router.GET("/example", Example)

	// http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	log.Println("Server listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
