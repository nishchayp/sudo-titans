package app

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"log"
)

/*
** contains init function to cofigure OAuth2 and Mysql
 */

func init() {

	// doesen't needs to be called explicitly
	// configures Mysql

	var credMysql CredentialMysql

	fileMysql, err := ioutil.ReadFile("./credMysql.json")
	if err != nil {
		log.Fatalf("Error while reading mysql config file.\nError: %v\n", err.Error())
	}
	json.Unmarshal(fileMysql, &credMysql)

	connectionString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		credMysql.DBUsername,
		credMysql.DBPassword,
		credMysql.DBName)

	DB.db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Could not open database : ", err.Error())
	}

}
