package handlers

import (
    "strconv"
    
    "github.com/gofiber/fiber/v2"
    "github.com/mrajnansky/go-test/database"
    "github.com/mrajnansky/go-test/models"
)

func ListFacts(c *fiber.Ctx) error {
    var facts []models.Fact
    
    // Parse query parameters
    page, err := strconv.Atoi(c.Query("page", "1"))
    if err != nil || page < 1 {
        page = 1
    }
    
    pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
    if err != nil || pageSize < 1 {
        pageSize = 10
    }
    
    // Calculate offset
    offset := (page - 1) * pageSize
    
    // Get total count
    var totalRecords int64
    database.DB.Db.Model(&models.Fact{}).Count(&totalRecords)
    
    // Get facts with pagination
    result := database.DB.Db.Limit(pageSize).Offset(offset).Find(&facts)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Could not retrieve facts",
            "error":   result.Error.Error(),
        })
    }
    
    // Create response
    response := fiber.Map{
        "facts":        facts,
        "page":         page,
        "pageSize":     pageSize,
        "totalRecords": totalRecords,
        "totalPages":   (totalRecords + int64(pageSize) - 1) / int64(pageSize),
    }
    
    return c.Status(fiber.StatusOK).JSON(response)
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
