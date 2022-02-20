package main

import (
	"main/boot"
)

// @title           Swagger JD-WEB API
// @version         1.01

// @description.markdown

// @host      stellaris.wang
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
// @schemes https

func main() {

	boot.ViperSetup()
	boot.ZapSetup()
	boot.MySQLSetup()
	boot.RedisSetup()
	boot.ServerSetup()

}
