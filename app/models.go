package app

/*
** contains all models of the data base
 */

type User struct {
	UserID   uint   `json:"user_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	TeamName string `json:"team_name" sql:"not null;unique"`
	Password string `json:"password" sql:"not null"`
	Name1    string `json:"name1" sql:"not null"`
	Name2    string `json:"name2" sql:"not null"`
	RegNo1   string `json:"reg_no_1" sql:"not null;unique"`
	RegNo2   string `json:"reg_no_2" sql:"not null;unique"`
}

type Admin struct {
	AdminID  uint   `json:"admin_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	TeamName string `json:"team_name" sql:"not null;unique"`
}

type Flag struct {
	FlagID     uint   `json:"flag_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	QuestionID string `json:"question_id" sql:"not null;unique"`
	Flag       string `json:"flag" sql:"not null"`
}

type PointsAndAccess struct {
	PointsAndAccessID uint   `json:"points_and_access_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	TeamName          string `json:"team_name" sql:"not null;unique"`
	MCQ_1             int    `json:"MCQ_1" gorm:"default:'0'" sql:"not null"`
	MCQ_2             int    `json:"MCQ_2" gorm:"default:'0'" sql:"not null"`
	MCQ_3             int    `json:"MCQ_3" gorm:"default:'0'" sql:"not null"`
	MCQ_4             int    `json:"MCQ_4" gorm:"default:'0'" sql:"not null"`
	MCQ_5             int    `json:"MCQ_5" gorm:"default:'0'" sql:"not null"`
	MCQ_6             int    `json:"MCQ_6" gorm:"default:'0'" sql:"not null"`
	MCQ_7             int    `json:"MCQ_7" gorm:"default:'0'" sql:"not null"`
	MCQ_8             int    `json:"MCQ_8" gorm:"default:'0'" sql:"not null"`
	MCQ_9             int    `json:"MCQ_9" gorm:"default:'0'" sql:"not null"`
	MCQ_10            int    `json:"MCQ_10" gorm:"default:'0'" sql:"not null"`
	CTF_1             int    `json:"CTF_1" gorm:"default:'-1'" sql:"not null"`
	CTF_2             int    `json:"CTF_2" gorm:"default:'-1'" sql:"not null"`
	CTF_3             int    `json:"CTF_3" gorm:"default:'-1'" sql:"not null"`
	CTF_4             int    `json:"CTF_4" gorm:"default:'-1'" sql:"not null"`
	CTF_5             int    `json:"CTF_5" gorm:"default:'-1'" sql:"not null"`
	CTF_6             int    `json:"CTF_6" gorm:"default:'-1'" sql:"not null"`
	CTF_7             int    `json:"CTF_7" gorm:"default:'-1'" sql:"not null"`
	CTF_8             int    `json:"CTF_8" gorm:"default:'-1'" sql:"not null"`
	CTF_9             int    `json:"CTF_9" gorm:"default:'-1'" sql:"not null"`
	CTF_10            int    `json:"CTF_10" gorm:"default:'-1'" sql:"not null"`
	TotalPoints       int    `json:"total_points" gorm:"default:'0'" sql:"not null"`
}

type PresentCTFValue struct {
	PresentCTFValueID uint `json:"present_ctf_value_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	CTF_1             int  `json:"CTF_1" sql:"not null"`
	CTF_2             int  `json:"CTF_2" sql:"not null"`
	CTF_3             int  `json:"CTF_3" sql:"not null"`
	CTF_4             int  `json:"CTF_4" sql:"not null"`
	CTF_5             int  `json:"CTF_5" sql:"not null"`
	CTF_6             int  `json:"CTF_6" sql:"not null"`
	CTF_7             int  `json:"CTF_7" sql:"not null"`
	CTF_8             int  `json:"CTF_8" sql:"not null"`
	CTF_9             int  `json:"CTF_9" sql:"not null"`
	CTF_10            int  `json:"CTF_10" sql:"not null"`
}

type McqDetail struct {
	McqDetailID uint   `json:"mcq_detail" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	QuestionID  string `json:"question_id" sql:"not null;unique"`
	Level       int    `json:"level" sql:"not null"` // 1-->Easy, 2-->Medium, 2-->Hard
	Question_1  string `json:"question_1" sql:"type:text;not null"`
	Option_1A   string `json:"option_1_A" sql:"not null"`
	Option_1B   string `json:"option_1_B" sql:"not null"`
	Option_1C   string `json:"option_1_C" sql:"not null"`
	Option_1D   string `json:"option_1_D" sql:"not null"`
	Question_2  string `json:"question_2" sql:"type:text;not null"`
	Option_2A   string `json:"option_2_A" sql:"not null"`
	Option_2B   string `json:"option_2_B" sql:"not null"`
	Option_2C   string `json:"option_2_C" sql:"not null"`
	Option_2D   string `json:"option_2_D" sql:"not null"`
	Question_3  string `json:"question_3" sql:"type:text;not null"`
	Option_3A   string `json:"option_3_A" sql:"not null"`
	Option_3B   string `json:"option_3_B" sql:"not null"`
	Option_3C   string `json:"option_3_C" sql:"not null"`
	Option_3D   string `json:"option_3_D" sql:"not null"`
	Question_4  string `json:"question_4" sql:"type:text;not null"`
	Option_4A   string `json:"option_4_A" sql:"not null"`
	Option_4B   string `json:"option_4_B" sql:"not null"`
	Option_4C   string `json:"option_4_C" sql:"not null"`
	Option_4D   string `json:"option_4_D" sql:"not null"`
}

type CtfDetail struct {
	CtfDetailID uint   `json:"ctf_detail" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	QuestionID  string `json:"question_id" sql:"not null;unique"`
	Level       int    `json:"level" sql:"not null"` // 1-->Easy, 2-->Medium, 2-->Hard
	Question    string `json:"question" sql:"type:text;not null"`
	Hint        string `json:"hint" sql:"type:text"`
}
