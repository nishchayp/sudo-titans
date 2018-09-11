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
