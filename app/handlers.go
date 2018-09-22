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

func SubmitFlag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Method != "POST" {
		return
	}

	var response Response

	if IsLoggedIn(w, r) != true {

		response = Response{
			false,
			"User not logged in",
		}

	} else {

		type Receive struct {
			QuestionID    string `json:"question_id"`
			CandidateFlag string `json:"candidate_flag"`
		}

		var flag Flag

		decoder := json.NewDecoder(r.Body)
		var receive Receive
		err := decoder.Decode(&receive)
		if err != nil {
			panic(err)
		}

		if DB.db.Where("question_id = ?", receive.QuestionID).First(&flag).RecordNotFound() {

			response = Response{
				false,
				"QuestionId does not exist",
			}

		} else {

			DB.db.Where("question_id = ?", receive.QuestionID).First(&flag)
			if receive.CandidateFlag == flag.Flag {

				response = Response{
					true,
					"Correct answer, flag found",
				}

			} else {

				response = Response{
					false,
					"Wrong answer, flag not found",
				}

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

func Mcq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var response Response

	if IsLoggedIn(w, r) != true {

		response = Response{
			false,
			"User not logged in",
		}

	} else {

		var mcqDetail McqDetail

		if DB.db.Where("question_id = ?", "MCQ_"+ps.ByName("idx")).First(&mcqDetail).RecordNotFound() {

			response = Response{
				false,
				"QuestionId does not exist",
			}

		} else {

			DB.db.Where("question_id = ?", "MCQ_"+ps.ByName("idx")).First(&mcqDetail)

			json, err := json.Marshal(mcqDetail)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(json)

			return

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
