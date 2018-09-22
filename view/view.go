package view

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

/*
** serves templates
** registers functions with routes
** opens up port and serves the app
 */

var err error

func Run() {

	// fs := http.FileServer(http.Dir("template"))
	// http.Handle("/sudo-titans/", http.StripPrefix("/sudo-titans/", fs))

	// http.HandleFunc("/example", Example)
	// http.HandleFunc("/", Index)
	// http.HandleFunc("/login", Login)

	// log.Println("Server listening at port 8000")
	// http.ListenAndServe(":8000", nil)

	router := httprouter.New()

	router.GET("/example", Example)
	router.GET("/", Index)
	router.POST("/login", Login)
	router.GET("/mcq/:idx", Mcq)

	log.Println("Server listening at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
