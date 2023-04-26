package main

import (
	"gotodo/api"
	"gotodo/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()                                      //นิยามตัว gin ขึ้นมา (นิยาม api)
	config := cors.DefaultConfig()                          //นิยามตัว cors ขึ้นมา
	config.AllowOrigins = []string("http://localhost:3000") //ให้สิทธิ์ในการเข้าถึง cors ตาม port ที่ระบุ
	r.Use(cors.New(config))

	database.ConnectDatabase()

	r.GET("/", api.GetAllLists)
	r.Run()
}
