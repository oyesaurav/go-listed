package controllers

import (
	"log"
	_ "os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/oyesaurav/go-todo/app/models"
	"github.com/oyesaurav/go-todo/pkg/utils"
	"github.com/oyesaurav/go-todo/platform/database"
)

func GetTasks(c *fiber.Ctx) error {
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	tasks, err := db.GetTasks()

	if err != nil {
		// Return status 500 and database query error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(tasks),
		"books": tasks,
	})
}

func CreateTasks(c *fiber.Ctx) error {
	// Create new Book struct
	task := &models.Task{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(task); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()
	task.ID = uuid.New()
	log.Print(uuid.New().String())
	task.Status = "TODO"
	task.CreatedAt = time.Now()
	if err := validate.Struct(task); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	if err := db.CreateTask(task); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"task":  task,
	})
}
