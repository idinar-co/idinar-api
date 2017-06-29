package config

import  "github.com/jinzhu/configor"

var Config = struct {
        DB   struct {
                Name     string `default:"idinar"`
                Adapter  string `default:"postgres"`
                User     string
                Password string
        }
}{}

func init() {
        if err := configor.Load(&Config, "config/database.yml"); err != nil {
                panic(err)
        }
}
