package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/mrajnansky/go-test/database"
    "github.com/mrajnansky/go-test/models"
    "github.com/mrajnansky/go-test/utils"
)

func ListFacts(c *fiber.Ctx) error {
    var facts []models.Fact
    return utils.HandlePagination(c, database.DB.Db, &models.Fact{}, &facts)
}

func CreateFact(c *fiber.Ctx) error {
    fact := new(models.Fact)
    if err := c.BodyParser(fact); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    database.DB.Db.Create(&fact)

    return c.Status(200).JSON(fact)
}

func GetRandomFact(c *fiber.Ctx) error {
    var fact models.Fact
    
    result := database.DB.Db.Order("RANDOM()").First(&fact)
    
    if result.Error != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "message": "No facts found",
            "error":   result.Error.Error(),
        })
    }

    return c.Status(200).JSON(fact)
}
