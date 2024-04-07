package handlers

import (

	"github.com/PatriciaDelCastillo/Rest-Api-Goland/database"
	"github.com/PatriciaDelCastillo/Rest-Api-Goland/models"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hola Mundo")

}
// listar 
func GetTasks(c *fiber.Ctx) error {
	db := database.DB
	var tasks []models.Task

	db.Find(&tasks)
	return c.JSON(tasks)

}

// crear 
func CreateTasks(c *fiber.Ctx) error {
	db := database.DB
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).JSON(err)
	}

	db.Create(&task)
	return c.JSON(task)

}

//  modificar o editar por id
func UpdateTasks(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var task models.Task

	if err := db.First(&task, id).Error; err != nil {
		return c.Status(404).JSON(err)
	}

	updateTasks := new(models.Task)
	if err := c.BodyParser(updateTasks); err != nil {
		return c.Status(400).JSON(err)
	}
	task.Body = updateTasks.Body
	db.Save(&task)
	return c.JSON(task)

}

// listar por id

func CreateTasksId(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var task models.Task

	db.Find(&task,id)
	return c.JSON(task)
}

//eliminar por id
func DeleteTask(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id")
	var task models.Task
	if err:= db.First(&task,id).Error;err!=nil{
		return c.Status(404).JSON(err)
	}
	db.Delete(&task)
	return c.JSON(fiber.Map{"status": "success","message":"Task Deleted","id":task.ID})
}
