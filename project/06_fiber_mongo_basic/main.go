package main

import (
	"context"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Get("/employee", func(c *fiber.Ctx) {
		query := bson.D{{}}
		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			c.Status(500).SendString(err.Error())
		}
		var employees []Employee
		if err := cursor.All(c.Context(), &employees); err != nil {
			c.Status(500).SendString(err.Error())
		}
		err = c.JSON(employees)
		if err != nil {
			c.Status(500).SendString(err.Error())
		}
	})

	app.Post("/employee", func(c *fiber.Ctx) {

	})

	app.Put("/employee/:id", func(c *fiber.Ctx) {

	})

	app.Delete("/employee/:id", func(c *fiber.Ctx) {

	})

}

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-mongo"
const mongoURI = "mongodb://localhost:27001/" + dbName

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}

func Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	db := client.Database(dbName)
	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}
