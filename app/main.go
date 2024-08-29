package main

import (
	"github.com/gin-gonic/gin"

	"productivityapp/models"
	"productivityapp/controllers"
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

	models.ConnectDatabase()
	r.GET("/contacts", controllers.FindContacts)
	r.POST("contacts", controllers.CreateContact)

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
