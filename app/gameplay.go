package app

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strings"
)

type Receive struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseFiles("template/index.html"))
	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// parse form into json struct
	type FormData struct {
		TeamName string
		Password string
	}

	r.ParseForm()

	formData := FormData{
		strings.Join(r.Form["team_name"], ""),
		strings.Join(r.Form["password"], ""),
	}

	var user User

	// check for teamName
	if DB.db.Where("team_name = ?", formData.TeamName).First(&user).RecordNotFound() {

		//template for incorrect teamName or pw
		t := template.Must(template.ParseFiles("template/index.html"))
		t.Execute(w, "Incorrect team name or password")

	} else {

		DB.db.Where("team_name = ?", formData.TeamName).First(&user)
		if user.Password != formData.Password {
			//template for incorrect teamName or pw
			t := template.Must(template.ParseFiles("template/index.html"))
			t.Execute(w, "Incorrect team name or password")

		} else {

			stUserCookie := &STUserCookie{
				UID:      user.UserID,
				TeamName: user.TeamName,
			}

			SetCookieHandler(stUserCookie, w)

			numOfSets := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
			t := template.Must(template.ParseFiles("template/home.html"))
			t.Execute(w, numOfSets)

		}

	}
}

func Mcq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if IsLoggedIn(w, r) != true {

		t := template.Must(template.ParseFiles("template/error.html"))
		t.Execute(w, "User not logged in")

	} else {

		var mcqDetail McqDetail

		type Data struct {
			QuestionID string
			Question_1 string
			Option_1A  string
			Option_1B  string
			Option_1C  string
			Option_1D  string
			Question_2 string
			Option_2A  string
			Option_2B  string
			Option_2C  string
			Option_2D  string
			Question_3 string
			Option_3A  string
			Option_3B  string
			Option_3C  string
			Option_3D  string
			Question_4 string
			Option_4A  string
			Option_4B  string
			Option_4C  string
			Option_4D  string
			Result_1   int
			Result_2   int
			Result_3   int
			Result_4   int
		}

		if DB.db.Where("question_id = ?", "MCQ_"+ps.ByName("idx")).First(&mcqDetail).RecordNotFound() {

			t := template.Must(template.ParseFiles("template/error.html"))
			t.Execute(w, "QuestionId does not exist")

		} else {

			DB.db.Where("question_id = ?", "MCQ_"+ps.ByName("idx")).First(&mcqDetail)

			data := Data{
				QuestionID: mcqDetail.QuestionID,
				Question_1: mcqDetail.Question_1,
				Option_1A:  mcqDetail.Option_1A,
				Option_1B:  mcqDetail.Option_1B,
				Option_1C:  mcqDetail.Option_1C,
				Option_1D:  mcqDetail.Option_1D,
				Question_2: mcqDetail.Question_2,
				Option_2A:  mcqDetail.Option_2A,
				Option_2B:  mcqDetail.Option_2B,
				Option_2C:  mcqDetail.Option_2C,
				Option_2D:  mcqDetail.Option_2D,
				Question_3: mcqDetail.Question_3,
				Option_3A:  mcqDetail.Option_3A,
				Option_3B:  mcqDetail.Option_3B,
				Option_3C:  mcqDetail.Option_3C,
				Option_3D:  mcqDetail.Option_3D,
				Question_4: mcqDetail.Question_4,
				Option_4A:  mcqDetail.Option_4A,
				Option_4B:  mcqDetail.Option_4B,
				Option_4C:  mcqDetail.Option_4C,
				Option_4D:  mcqDetail.Option_4D,
				Result_1:   0,
				Result_2:   0,
				Result_3:   0,
				Result_4:   0,
			}

			t := template.Must(template.ParseFiles("template/preliminaryMcq.html"))
			t.Execute(w, data)

		}
	}
}

func CheckFlagMcq(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if IsLoggedIn(w, r) != true {

		t := template.Must(template.ParseFiles("template/error.html"))
		t.Execute(w, "User not logged in")

	} else {

		r.ParseForm()
		formData := strings.Join(r.Form["flag"], "")

		var flag Flag

		if DB.db.Where("question_id = ?", ps.ByName("question_id")).First(&flag).RecordNotFound() {

			t := template.Must(template.ParseFiles("template/error.html"))
			t.Execute(w, "QuestionId does not exist")

		} else {

			var mcqDetail McqDetail

			type Data struct {
				QuestionID string
				Question_1 string
				Option_1A  string
				Option_1B  string
				Option_1C  string
				Option_1D  string
				Question_2 string
				Option_2A  string
				Option_2B  string
				Option_2C  string
				Option_2D  string
				Question_3 string
				Option_3A  string
				Option_3B  string
				Option_3C  string
				Option_3D  string
				Question_4 string
				Option_4A  string
				Option_4B  string
				Option_4C  string
				Option_4D  string
				Result_1   int
				Result_2   int
				Result_3   int
				Result_4   int
			}

			var data Data

			DB.db.Where("question_id = ?", ps.ByName("question_id")).First(&flag)

			count := 0
			var (
				res1 int
				res2 int
				res3 int
				res4 int
			)

			if formData[0] == flag.Flag[0] {
				res1 = 1
				count++
			} else {
				res1 = -1
			}

			if formData[1] == flag.Flag[1] {
				res2 = 1
				count++
			} else {
				res2 = -1
			}

			if formData[2] == flag.Flag[2] {
				res3 = 1
				count++
			} else {
				res3 = -1
			}

			if formData[3] == flag.Flag[3] {
				res4 = 1
				count++
			} else {
				res4 = -1
			}

			if count < 4 {
				DB.db.Where("question_id = ?", ps.ByName("question_id")).First(&mcqDetail)
				data = Data{
					QuestionID: mcqDetail.QuestionID,
					Question_1: mcqDetail.Question_1,
					Option_1A:  mcqDetail.Option_1A,
					Option_1B:  mcqDetail.Option_1B,
					Option_1C:  mcqDetail.Option_1C,
					Option_1D:  mcqDetail.Option_1D,
					Question_2: mcqDetail.Question_2,
					Option_2A:  mcqDetail.Option_2A,
					Option_2B:  mcqDetail.Option_2B,
					Option_2C:  mcqDetail.Option_2C,
					Option_2D:  mcqDetail.Option_2D,
					Question_3: mcqDetail.Question_3,
					Option_3A:  mcqDetail.Option_3A,
					Option_3B:  mcqDetail.Option_3B,
					Option_3C:  mcqDetail.Option_3C,
					Option_3D:  mcqDetail.Option_3D,
					Question_4: mcqDetail.Question_4,
					Option_4A:  mcqDetail.Option_4A,
					Option_4B:  mcqDetail.Option_4B,
					Option_4C:  mcqDetail.Option_4C,
					Option_4D:  mcqDetail.Option_4D,
					Result_1:   res1,
					Result_2:   res2,
					Result_3:   res3,
					Result_4:   res4,
				}

				// read cookie to get team name to pass it in EditScoreFunction

				stUserCookie := ReadCookieHandler(w, r)

				EditScoreMcq(stUserCookie.TeamName, ps.ByName("question_id"))

				t := template.Must(template.ParseFiles("template/preliminaryMcq.html"))
				t.Execute(w, data)
			} else {

				t := template.Must(template.ParseFiles("template/finalCtf.html"))
				t.Execute(w, "hello")
			}

		}
	}
}
