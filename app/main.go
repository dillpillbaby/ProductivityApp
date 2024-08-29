package main

import (
	"productivityapp/controllers"
	"productivityapp/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "1234ASDFjkl;"
// 	dbname   = "ProductivityApp"
// )

// type test struct {
// 	ID   uint
// 	Name string
// }

func main() {
	r := gin.Default()

	// Set up CORS middleware
	config := cors.Config{
		AllowOrigins: []string{"http://localhost:5173"}, // Replace with your frontend URL
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}

	r.Use(cors.New(config))

	models.ConnectDatabase()
	r.GET("/contacts", controllers.FindContacts)
	r.GET("/colors", controllers.FindColors)
	r.GET("/journalentries", controllers.FindJournalEntries)
	r.GET("/moods", controllers.FindMoods)
	r.POST("contacts", controllers.CreateContact)
	r.POST("/colors", controllers.CreateColor)
	r.POST("/journalentries", controllers.CreateJournalEntry)
	r.POST("/moods", controllers.CreateMood)
	r.DELETE("/moods/:id", controllers.DeleteMood)

	r.Run()
	// dsn := "host=localhost user=postgres password=1234ASDFjkl; dbname=ProductivityApp port=5432"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// })
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// readTest := &test{}
	// db.Unscoped().First(&readTest)
	// fmt.Println("ID: %d, Name: %s\n", readTest.ID, readTest.Name)
}
