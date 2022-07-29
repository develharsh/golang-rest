package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// "github.com/develharsh/golang-fiber-postgresql-rest-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func GetBookByID(context *fiber.Ctx) error {
	ID, err := strconv.Atoi(context.Params("id"))
	if err != nil {
		fmt.Println(err)
	}
	// context.Status(http.StatusBadRequest).JSON(
	// 	&fiber.Map{"message": ID, "success": false})
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": ID,
		"success": true})
	return nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Use(logger.New())
	api := app.Group("/api")
	api.Get("/get_book/:id", GetBookByID)
	app.Listen(os.Getenv("BASEURL"))
}
