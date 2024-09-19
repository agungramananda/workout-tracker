package configs

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine{
	r := gin.Default()
	
	return r
}

func GetServerPort() string {
	port := GetDotEnvVariable("PORT")
	if port != "" {
		return ":" + port
	}
	return ":8080"
}