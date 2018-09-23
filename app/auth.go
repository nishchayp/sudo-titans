package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ApiLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method != "POST" {
		return
	}

	type Receive struct {
		TeamName string `json:"team_name"`
		Password string `json:"password"`
	}

	// var response Response
	type ResponseLogin struct {
		Success  bool   `json:"success"`
		Message  string `json:"message"`
		UID      uint   `json: "uid"`
		TeamName string `json:"team_name"`
	}

	var responseLogin ResponseLogin

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
		responseLogin = ResponseLogin{
			false,
			"Incorrect team name or password",
			999999,
			"",
		}

	} else {

		DB.db.Where("team_name = ?", receive.TeamName).First(&user)
		if user.Password != receive.Password {
			//json for incorrect teamName or pw
			responseLogin = ResponseLogin{
				false,
				"Incorrect team name or password",
				999999,
				"",
			}
		} else {

			stUserCookie := &STUserCookie{
				UID:      user.UserID,
				TeamName: user.TeamName,
			}

			SetCookieHandler(stUserCookie, w)

			responseLogin = ResponseLogin{
				true,
				"User logged in",
				user.UserID,
				user.TeamName,
			}

		}

	}

	json, err := json.Marshal(responseLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func ApiLogout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// clears cookie logs out user

	ClearCookieHandler(w)
	http.Redirect(w, r, "/", 302)

}

func IsLoggedIn(w http.ResponseWriter, r *http.Request) (isAdminFlag bool) {

	// read cookie and check db if email exists in admins table

	stUserCookie := ReadCookieHandler(w, r)

	var user User

	if DB.db.Where("team_name = ?", stUserCookie.TeamName).First(&user).RecordNotFound() {
		return false
	} else {
		return true
	}
}

func IsAdmin(w http.ResponseWriter, r *http.Request) (isAdminFlag bool) {

	// read cookie and check db if email exists in admins table

	stUserCookie := ReadCookieHandler(w, r)

	var admin Admin

	if DB.db.Where("team_name = ?", stUserCookie.TeamName).First(&admin).RecordNotFound() {
		return false
	} else {
		return true
	}
}
