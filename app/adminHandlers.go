package app

import (
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

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

func AdminIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if IsAdmin(w, r) != true {
		http.Redirect(w, r, "/", 302)
	} else {
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

}
