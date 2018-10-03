package app

import (
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func AdminIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if IsAdmin(w, r) != true {
		http.Redirect(w, r, "/", 302)
		return
	}

	response := &Response{
		true,
		"Welcome Admin to Sudo Titans",
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}

func SetPresentCTFValues(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if IsAdmin(w, r) != true {
		http.Redirect(w, r, "/", 302)
		return
	}

	// set up initial value of CTF questions
	var initialPresentCTFValue = PresentCTFValue{
		CTF_1:  MaxCTFEasy,
		CTF_2:  MaxCTFEasy,
		CTF_3:  MaxCTFEasy,
		CTF_4:  MaxCTFMedium,
		CTF_5:  MaxCTFMedium,
		CTF_6:  MaxCTFMedium,
		CTF_7:  MaxCTFMedium,
		CTF_8:  MaxCTFHard,
		CTF_9:  MaxCTFHard,
		CTF_10: MaxCTFHard,
	}

	var response Response

	err = DB.db.Create(&initialPresentCTFValue).Error
	if err != nil {
		log.Printf("Error initializing new user", err.Error())
		response = Response{
			false,
			"Error initializing present ctf value",
		}
	} else {
		response = Response{
			true,
			"Initialized present ctf value",
		}
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}

func AddUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method != "POST" {
		return
	}

	if IsAdmin(w, r) != true {
		http.Redirect(w, r, "/", 302)
		return
	}

	type Receive struct {
		TeamName string `json:"team_name"`
		Password string `json:"password"`
		Name1    string `json:"name1"`
		Name2    string `json:"name2"`
		RegNo1   string `json:"reg_no_1"`
		RegNo2   string `json:"reg_no_2"`
	}

	decoder := json.NewDecoder(r.Body)
	var receive Receive
	err := decoder.Decode(&receive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user = User{
		TeamName: receive.TeamName,
		Password: receive.Password,
		Name1:    receive.Name1,
		Name2:    receive.Name2,
		RegNo1:   receive.RegNo1,
		RegNo2:   receive.RegNo2,
	}

	var response Response

	var initialPointsAndAccess = PointsAndAccess{
		TeamName: receive.TeamName,
	}

	err = DB.db.Create(&user).Error
	if err != nil {
		log.Printf("Error creating new user", err.Error())
		response = Response{
			false,
			"Error creating new user",
		}
	} else {
		err = DB.db.Create(&initialPointsAndAccess).Error
		if err != nil {
			log.Printf("Error initializing new user", err.Error())
			response = Response{
				false,
				"Error initializing new user",
			}
		} else {
			response = Response{
				true,
				"New user created and initialized",
			}
		}
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
