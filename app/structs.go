package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
** contains all structs used in the app
 */

type Database struct {
	db *gorm.DB
}

type CredentialMysql struct {
	DBUsername string `json:"db_username"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type STUserCookie struct {
	UID      uint   `json: "uid"`
	TeamName string `json:"team_name"`
}
