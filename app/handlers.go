package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	response := &Response{
		true,
		"Welcome to Sudo Titans",
	}

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method != "POST" {
		return
	}

	type Receive struct {
		TeamName string `json:"team_name"`
		Password string `json:"password"`
	}

	var response Response

	var user User

	decoder := json.NewDecoder(r.Body)
	var receive Receive
	err := decoder.Decode(&receive)
	if err != nil {
		panic(err)
	}

	// check for teamName
	if DB.db.Where("team_name = ?", receive.TeamName).First(&user).RecordNotFound() {

		//json for incorrect teamName or pw
		response = Response{
			false,
			"Incorrect team name or password",
		}

	} else {

		DB.db.Where("team_name = ?", receive.TeamName).First(&user)
		if user.Password != receive.Password {
			//json for incorrect teamName or pw
			response = Response{
				false,
				"Incorrect team name or password",
			}
		} else {

			stUserCookie := &STUserCookie{
				UID:      user.UserID,
				TeamName: user.TeamName,
			}

			SetCookieHandler(stUserCookie, w)

			response = Response{
				true,
				"User logged in",
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

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// clears cookie logs out user

	ClearCookieHandler(w)
	http.Redirect(w, r, "/", 302)

}
