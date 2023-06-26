package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gitub.com/joho.godotenv"
	"gorm.io/gorm"
)


type Book struct{

Author string   `json: "author"`
Title string    `json:"title"`
Publisher string `json:"publisher"`
}


type Repository struct{
	DB *gorm.DB
}
func (r *Repository)SetupRoutes(app *fiber.App)
	api := app.Group("/api")
	api.Post("/create_books",r.Createbook)
	api.Delete("/delet_books/:id",r.Deletebook)
	api.Get("get_books/:id",r.GetBookID)
	api.Get("/books",r.Getbooks)
	err := godotenv.Load(".env")
	if err != nil {

		log.Fatal(err)
	}

func main() {
	r := Repository{// own datatypes created

		DB: db,
	}
	db ,err := storage.newConnection(config)
	if err != nil{
		log.Fatal("could not load the database")
	}
app := fiber.new()
r.SetupRoutes(app) //struct method
app.Listen(":8080")


}


