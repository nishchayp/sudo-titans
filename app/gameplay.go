package app

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
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
			Level      int // 1-->Easy, 2-->Medium, 2-->Hard
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
				Level:      mcqDetail.Level,
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

			// read cookie to get team name to pass it in EditScoreFunction
			stUserCookie := ReadCookieHandler(w, r)

			var mcqDetail McqDetail
			DB.db.Where("question_id = ?", ps.ByName("question_id")).First(&mcqDetail)

			type Data struct {
				QuestionID string
				Level      int // 1-->Easy, 2-->Medium, 2-->Hard
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

			data := Data{
				QuestionID: mcqDetail.QuestionID,
				Level:      mcqDetail.Level,
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

			DB.db.Where("question_id = ?", ps.ByName("question_id")).First(&flag)

			count := 0

			if formData[0] == flag.Flag[0] {
				data.Result_1 = 1
				count++
			} else {
				data.Result_1 = -1
			}

			if formData[1] == flag.Flag[1] {
				data.Result_2 = 1
				count++
			} else {
				data.Result_2 = -1
			}

			if formData[2] == flag.Flag[2] {
				data.Result_3 = 1
				count++
			} else {
				data.Result_3 = -1
			}

			if formData[3] == flag.Flag[3] {
				data.Result_4 = 1
				count++
			} else {
				data.Result_4 = -1
			}

			if count < 4 {

				EditScoreMcq(false, stUserCookie.TeamName, data.QuestionID, data.Level)

				t := template.Must(template.ParseFiles("template/preliminaryMcq.html"))
				t.Execute(w, data)
			} else {

				EditScoreMcq(true, stUserCookie.TeamName, data.QuestionID, data.Level)

				var ctfDetail CtfDetail
				DB.db.Where("question_id = ?", strings.Replace(ps.ByName("question_id"), "MCQ", "CTF", -1)).First(&ctfDetail)

				t := template.Must(template.ParseFiles("template/finalCtf.html"))
				t.Execute(w, ctfDetail)
			}

		}
	}
}

func CheckFlagCtf(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

			// read cookie to get team name to pass it in EditScoreFunction
			stUserCookie := ReadCookieHandler(w, r)

			var ctfDetail CtfDetail
			DB.db.Where("question_id = ?", ps.ByName("question_id")).First(&ctfDetail)
			type Data struct {
				CtfDetailID uint
				QuestionID  string
				Level       int
				Question    string
				Hint        string
				Result      int
				Message     string
			}
			data := Data{
				CtfDetailID: ctfDetail.CtfDetailID,
				QuestionID:  ctfDetail.QuestionID,
				Level:       ctfDetail.Level,
				Question:    ctfDetail.Question,
				Hint:        ctfDetail.Hint,
				Result:      0,
			}

			DB.db.Where("question_id = ?", ps.ByName("question_id")).First(&flag)

			// func EditScoreCtf
			// return values
			//  1-->Correct
			//  0-->Already answered
			// -1-->Wrong
			// -2-->Not accesible

			if formData == flag.Flag {
				data.Result = EditScoreCtf(true, stUserCookie.TeamName, ctfDetail.QuestionID, ctfDetail.Level)
			} else {
				data.Result = EditScoreCtf(false, stUserCookie.TeamName, ctfDetail.QuestionID, ctfDetail.Level)
			}

			if data.Result == 1 {
				data.Message = "Correct answer!"
				t := template.Must(template.ParseFiles("template/scoreBoard.html"))
				t.Execute(w, data.Message)
			} else if data.Result == 0 {
				data.Message = "You have already answered this question correctly"
				t := template.Must(template.ParseFiles("template/scoreBoard.html"))
				t.Execute(w, data.Message)
			} else if data.Result == -1 {
				data.Message = "Wrong answer!"
				t := template.Must(template.ParseFiles("template/finalCtf.html"))
				t.Execute(w, data)
			} else if data.Result == -2 {
				data.Message = "Question not unlocked, answer the preliminary MCQs to unlock this question"
				t := template.Must(template.ParseFiles("template/finalCtf.html"))
				t.Execute(w, data)
			} else {
				log.Println(stUserCookie.TeamName + ": Invalid data.result value returned from EditScoreCtf")
			}

		}

	}
}

func Scoreboard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// var pointsAndAccess PointsAndAccess
	// DB.db.Limit(8).Order("total_points").Find(&pointsAndAccess)
	// for _, row := range pointsAndAccess {
	// 	log.Println(row.TeamName)
	// }
}
