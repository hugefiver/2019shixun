package main

import (
	"log"
	"sx/middleware"
	"sx/model"
	. "sx/route"
	. "sx/route/api"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	db, err := model.ConnectDB("./lite.db")
	if err != nil {
		log.Fatal("DB Init Error:", err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Person{})
	db.AutoMigrate(&model.Image{})
	db.AutoMigrate(&model.Mochine{})
	db.AutoMigrate(&model.McUse{})
	db.AutoMigrate(&model.McBorrow{})
	//model.AddPerson("admin", "Admin", model.AdminLevel, "admin")

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("template/*")

	router.Use(static.Serve("/", static.LocalFile("html", false)))

	router.GET("/login", Login)

	api := router.Group("/api")
	api.Use(middleware.CorsMiddle())
	{
		// Login and Check
		api.GET("/login", LoginApi)
		api.GET("/check", CheckLoginApi)

		// Get Image
		api.GET("/people/:id/image", GetImage)

		user := api.Group("/")
		user.Use(middleware.LoginOnlyMiddle())
		{
			user.GET("/people", GetPeople)
			user.GET("/people/:id", GetPersonInfo)

			user.GET("/mochines", GetMochines)
			user.GET("/mochines/:id", GetMochineInfo)

			user.POST("/mochines/:id/use", UseMochine)
			user.GET("/mochines/:id/use", CheckUse)

			user.POST("/mochines/:id/borrow", Borrow)
			user.POST("/mochines/:id/return", Return)
			user.GET("/borrows", BorrowSearch)
		}

		admin := api.Group("/")
		admin.Use(middleware.AdminOnlyMiddle())
		{
			// Person
			admin.POST("/people", AddPerson)
			admin.POST("/people/:id", UpdatePersonInfo)
			admin.POST("/people/:id/image", UpdateImage)
			admin.DELETE("/people/:id", DeletePerson)

			// Mochine
			admin.POST("/mochines/:id", UpdateMochine)
			admin.DELETE("/mochines/:id", DeleteMochine)
		}
	}

	if gin.IsDebugging() {
		router.Run(":8081")
	} else {
		log.Println("Serve on http://127.0.0.1:8080")
		router.Run(":8080")
	}

}
