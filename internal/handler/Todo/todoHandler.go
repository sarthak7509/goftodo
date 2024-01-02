package Todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sarthak7509/goftodo/database"
	"github.com/sarthak7509/goftodo/internal/models"
)

func CreateTodo(c *fiber.Ctx, user models.User) error {
	db := database.DB
	payload := new(models.TodoCreatePayload)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	newTodo := models.Todo{
		ID:             uuid.New(),
		Title:          payload.Title,
		SubTitle:       payload.SubTitle,
		EstimateSprint: payload.EstimateSprint,
		IsDone:         payload.IsDone,
		Priority:       payload.Priority,
		UserId:         user.Id,
		User:           user,
	}

	errs := db.Create(&newTodo).Error
	if errs != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": models.TodoGetResponse(newTodo)})
}

func GetTodoList(c *fiber.Ctx, user models.User) error {
	db := database.DB
	var todos []models.Todo
	err := db.Find(&todos, "user_id=?", user.Id).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	return c.Status(200).JSON(fiber.Map{"data": models.ListOfTodos(todos)})
}

func GetTodo(c *fiber.Ctx, user models.User) error {
	db := database.DB
	id := c.Params("todoId")
	var todo models.Todo
	err := db.Find(&todo, "id=?", id).Error

	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Todo present", "data": nil})
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	return c.Status(200).JSON(fiber.Map{"data": models.TodoGetResponse(todo)})
}

func UpdateTodo(c *fiber.Ctx, user models.User) error {
	db := database.DB

	payload := new(models.TodoCreatePayload)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	id := c.Params("todoId")
	var todo models.Todo
	err = db.Find(&todo, "id=?", id).Error

	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No Todo present", "data": nil})
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	todo.Title = payload.Title
	todo.SubTitle = payload.SubTitle
	todo.EstimateSprint = payload.EstimateSprint
	todo.Priority = payload.Priority
	todo.IsDone = payload.IsDone

	db.Save(&todo)

	return c.Status(200).JSON(fiber.Map{"data": models.TodoGetResponse(todo)})
}

func DeleteTodo(c *fiber.Ctx, user models.User) error {
	db := database.DB
	id := c.Params("todoId")
	var todo models.Todo
	err := db.Find(&todo, "id=?", id).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	err = db.Delete(&todo).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return c.SendStatus(200)
}
