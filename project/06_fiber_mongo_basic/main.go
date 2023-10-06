package main

import (
	"context"
	"errors"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		collection := mg.Db.Collection("employees")
		employee := new(Employee)
		if err := c.BodyParser(employee); err != nil {
			c.Status(500).SendString(err.Error())
		}
		employee.ID = ""
		insertResult, err := collection.InsertOne(c.Context(), employee)
		if err != nil {
			c.Status(500).SendString(err.Error())
		}
		filer := bson.D{{Key: "_id", Value: insertResult.InsertedID}}
		createdRecord := collection.FindOne(c.Context(), filer)
		createEmployee := &Employee{}
		err = createdRecord.Decode(createEmployee)
		if err != nil {
			c.Status(500).SendString(err.Error())
		}
		err = c.Status(200).JSON(createEmployee)
		if err != nil {
			c.Status(500).SendString(err.Error())
		}
	})

	app.Put("/employee/:id", func(c *fiber.Ctx) {
		id := c.Params("id")
		employeeId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.SendStatus(400)
		}
		employee := new(Employee)
		if err := c.BodyParser(employee); err != nil {
			c.Status(400).SendString(err.Error())
		}
		query := bson.D{{Key: "_id", Value: employeeId}}
		update := bson.D{{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		}}
		err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				c.SendStatus(400)
			}
			c.SendStatus(500)
		}
		employee.ID = id
		err = c.Status(200).JSON(employee)
		if err != nil {
			c.Status(500).SendString(err.Error())
		}
	})

	app.Delete("/employee/:id", func(c *fiber.Ctx) {
		employeeId, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			c.Status(400).SendString(err.Error())
		}
		query := bson.D{{Key: "_id", Value: employeeId}}
		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), &query)
		if err != nil {
			c.SendStatus(500)
		}
		if result.DeletedCount < 1 {
			c.SendStatus(400)
		}
		err = c.Status(200).JSON("record deleted")
		if err != nil {
			c.Status(500).SendString(err.Error())
		}
	})

	log.Fatal(app.Listen("localhost:3000"))
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
