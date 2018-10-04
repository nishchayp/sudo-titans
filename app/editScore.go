package app

import (
	"log"
	"reflect"
	"strings"
)

func getFieldPointsAndAccess(v *PointsAndAccess, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func getFieldPresentCTFValue(v *PresentCTFValue, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func EditScoreMcq(isCorrect bool, teamName string, questionID string, level int, decPoints int) {
	// check if already answered
	var pointsAndAccess PointsAndAccess
	DB.db.Where("team_name = ?", teamName).First(&pointsAndAccess)
	val := getFieldPointsAndAccess(&pointsAndAccess, questionID)
	if val > 0 {
		log.Println("Already answered MCQ Correctly")
		return
	} else {
		var points int
		if isCorrect == true {
			switch level {
			case 1:
				points = MCQEasy
				break
			case 2:
				points = MCQMedium
				break
			case 3:
				points = MCQHard
			}
			DB.db.Model(&pointsAndAccess).Update(strings.ToLower(questionID), points)
			DB.db.Model(&pointsAndAccess).Update(strings.Replace(strings.ToLower(questionID), "mcq", "ctf", -1), 0)
			DB.db.Model(&pointsAndAccess).Update("total_points", pointsAndAccess.TotalPoints+points)
		} else {
			DB.db.Model(&pointsAndAccess).Update("total_points", pointsAndAccess.TotalPoints-decPoints)
		}
	}
}

func EditScoreCtf(isCorrect bool, teamName string, questionID string, level int) (result int) {
	// check if already answered
	// check if unlocked

	// return values
	//  1-->Correct
	//  0-->Already answered
	// -1-->Wrong
	// -2-->Not accesible

	var pointsAndAccess PointsAndAccess
	DB.db.Where("team_name = ?", teamName).First(&pointsAndAccess)
	val := getFieldPointsAndAccess(&pointsAndAccess, questionID)
	if val < 0 {
		return -2
	} else if val > 0 {
		return 0
	} else {
		var presentCTFValue PresentCTFValue
		DB.db.First(&presentCTFValue)
		var decValue int
		if isCorrect == true {
			switch level {
			case 1:
				decValue = DecCTFEasy
				break
			case 2:
				decValue = DecCTFMedium
				break
			case 3:
				decValue = DecCTFHard
			}
			points := getFieldPresentCTFValue(&presentCTFValue, questionID)
			DB.db.Model(&pointsAndAccess).Update(strings.ToLower(questionID), points)
			DB.db.Model(&pointsAndAccess).Update("total_points", pointsAndAccess.TotalPoints+points)
			DB.db.Model(&presentCTFValue).Update(questionID, getFieldPresentCTFValue(&presentCTFValue, questionID)-decValue)
			log.Println("supporting log", teamName, questionID, points)
			return 1
		} else {
			return -1
		}
	}
}
