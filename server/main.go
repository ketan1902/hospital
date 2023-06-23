package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ketan1902/golang-react-appointment/routes"
)

func main() {
	// port := os.Getenv("PORT")

	// if port == "" {
	// 	port = "8000"
	// }

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id", routes.GetEntryById)
	router.GET("/entry/disease", routes.GetEntriesByDisease)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/disease/update/:id", routes.UpdateDisease)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)
	router.Run("localhost:8008")

}
