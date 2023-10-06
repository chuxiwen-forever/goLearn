package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/liu/fiber-crm-basic/database"
	"github.com/liu/fiber-crm-basic/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("", lead.GetLeads)
	app.Get("", lead.GetLead)
	app.Post("", lead.NewLead)
	app.Delete("", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
	defer func(DBConn *gorm.DB) {
		err := DBConn.Close()
		if err != nil {
		}
	}(database.DBConn)
}
