package db

import (
        "log"
        "fmt"

        "github.com/jinzhu/gorm"
        _ "github.com/jinzhu/gorm/dialects/mysql"
        "srv.nazrol.tech/nuzaihan/erating-go/config"
)

var DBCon *gorm.DB

func init(){
    // We set up the database
    var err error

    dbConfig := config.Config.DB

    DBCon, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(localhost:3306)/%v?charset=utf8&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Name))
    DBCon.DB().Ping()
    DBCon.DB().SetMaxIdleConns(10)
    DBCon.DB().SetMaxOpenConns(100)
    DBCon.LogMode(true) // SQL Logging

    // Stop running if there are problems connecting to database
    if err != nil {
      log.Fatal("%s - cannot connect to database", err)
    }
}