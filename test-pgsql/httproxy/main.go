package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "httproxy/routers"
)

func main() {
	beego.AppConfig.String("pgsql_user")
	beego.AppConfig.String("pgsql_password")
	beego.AppConfig.String("pgsql_host")
	beego.AppConfig.String("pgsql_dbname")
	beego.Run()
}
