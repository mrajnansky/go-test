package utils

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type PaginationParams struct {
    Page     int `json:"page"`
    PageSize int `json:"pageSize"`
    Offset   int `json:"-"`
}

type PaginationResponse struct {
    Records      interface{} `json:"records"`
    Page         int         `json:"page"`
    PageSize     int         `json:"pageSize"`
    TotalRecords int64       `json:"totalRecords"`
    TotalPages   int64       `json:"totalPages"`
}

func ParsePaginationParams(c *fiber.Ctx) PaginationParams {
    page, err := strconv.Atoi(c.Query("page", "1"))
    if err != nil || page < 1 {
        page = 1
    }

    pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
    if err != nil || pageSize < 1 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize

    return PaginationParams{
        Page:     page,
        PageSize: pageSize,
        Offset:   offset,
    }
}

func GetTotalCount(db *gorm.DB, model interface{}) (int64, error) {
    var totalRecords int64
    err := db.Model(model).Count(&totalRecords).Error
    return totalRecords, err
}

func PaginateQuery(db *gorm.DB, params PaginationParams, result interface{}) error {
    return db.Limit(params.PageSize).Offset(params.Offset).Find(result).Error
}

func BuildPaginationResponse(records interface{}, params PaginationParams, totalRecords int64) PaginationResponse {
    totalPages := (totalRecords + int64(params.PageSize) - 1) / int64(params.PageSize)
    
    return PaginationResponse{
        Records:      records,
        Page:         params.Page,
        PageSize:     params.PageSize,
        TotalRecords: totalRecords,
        TotalPages:   totalPages,
    }
}

func HandlePagination(c *fiber.Ctx, db *gorm.DB, model interface{}, result interface{}) error {
    params := ParsePaginationParams(c)

    totalRecords, err := GetTotalCount(db, model)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Could not count records",
            "error":   err.Error(),
        })
    }

    err = PaginateQuery(db, params, result)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Could not retrieve records",
            "error":   err.Error(),
        })
    }

    response := BuildPaginationResponse(result, params, totalRecords)
    
    return c.Status(fiber.StatusOK).JSON(response)
}
