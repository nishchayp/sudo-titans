package view

import (
	"bytes"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Receive struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func Example(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("here")
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	t := template.Must(template.ParseFiles("template/ex.html"))
	t.Execute(w, data)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseFiles("template/index.html"))
	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// parse form into json struct
	type Request struct {
		TeamName string `json:"team_name"`
		Password string `json:"password"`
	}

	r.ParseForm()

	request := Request{
		strings.Join(r.Form["team_name"], ""),
		strings.Join(r.Form["password"], ""),
	}

	// marshall struct to json
	jsonReq, err := json.Marshal(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// make post request and get back response
	u, err := url.Parse("http://127.0.0.1:8080/api/login")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatal(err)
	}

	// read response in byte form
	defer resp.Body.Close()

	resp_bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshall json to struct
	var receive Receive

	err = json.Unmarshal(resp_bytes, &receive)
	if err != nil {
		log.Fatal(err)
	}

	// parse and execute apropriate template
	if receive.Success == false {
		t := template.Must(template.ParseFiles("template/index.html"))
		t.Execute(w, "Incorrect team name or password")
	} else {
		numOfSets := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		t := template.Must(template.ParseFiles("template/home.html"))
		t.Execute(w, numOfSets)
	}
}

func Mcq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	log.Println("pb1") //
	// make get request and get back response
	u, err := url.Parse("http://127.0.0.1:8080/api/mcq/" + ps.ByName("idx"))
	log.Println(u.String()) //
	if err != nil {
		log.Println(err)
	}

	log.Println("pb2") //
	resp, err := http.Get(u.String())
	if err != nil {
		log.Println("here1") //
		log.Println(err)
	}

	// read response in byte form
	log.Println("pb3") //
	defer resp.Body.Close()

	resp_bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("here2") //
		log.Println(err)
	}

	log.Println(string(resp_bytes))

	// unmarshall json to struct
	var receive Receive
	log.Println("pb4") //

	err = json.Unmarshal(resp_bytes, &receive)
	if err != nil {
		log.Println(err)
	}
	log.Println("pb5") //

	// parse and execute apropriate template
	// if receive.Success == false {
	// 	t := template.Must(template.ParseFiles("template/index.html"))
	// 	t.Execute(w, "Incorrect team name or password")
	// } else {
	// 	numOfSets := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 	t := template.Must(template.ParseFiles("template/home.html"))
	// 	t.Execute(w, numOfSets)
	// }
}
