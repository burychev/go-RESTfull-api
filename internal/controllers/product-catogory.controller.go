package controllers

import (
	"products/internal/common"
	"products/internal/contracts"
	"products/internal/models"
	"products/internal/services"
	"products/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductCategoryController struct {
	service services.IProductCategoryService
}

func NewProductCategoryController(api fiber.Router, s services.IProductCategoryService) *ProductCategoryController {
	controller := &ProductCategoryController{service: s}

	api.Post("/categories", controller.CreateCategory)
	api.Put("/categories/:id", controller.UpdateCategory)
	api.Delete("/categories/:id", controller.DeleteCategory)
	api.Get("/categories/:id", controller.GetCategoryById)
	api.Get("/categories", controller.GetAll)
	return controller
}

func (c *ProductCategoryController) CreateCategory(ctx *fiber.Ctx) error {
	var req contracts.CreateProductCategoryRequest

	if err := utils.ReadRequestFromBody(ctx, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BaseResponse{Message: "Invalid request", Data: nil})
	}

	entity := models.ProductCategory{Name: req.Name, Description: req.Description}
	createdEntity, creationErr := c.service.CreateCategory(entity)

	if creationErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.BaseResponse{Message: "Creation error", Data: nil})
	}

	res := &contracts.CreateProductCategoryResponse{
		Id:          createdEntity.Id,
		Name:        createdEntity.Name,
		Description: createdEntity.Description,
	}
	return ctx.JSON(common.BaseResponse{Message: "Created successfully", Data: res})
}

func (c *ProductCategoryController) UpdateCategory(ctx *fiber.Ctx) error {
	var req contracts.UpdateProductCategoryRequest
	id, _ := strconv.Atoi(ctx.Params("id"))

	if err := utils.ReadRequestFromBody(ctx, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BaseResponse{Message: "Invalid request", Data: nil})
	}

	entity := models.ProductCategory{Id: id, Name: req.Name, Description: req.Description}

	updated, updationErr := c.service.UpdateCategory(entity)

	if updationErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.BaseResponse{Message: "Updation error", Data: nil})
	}

	res := &contracts.CreateProductCategoryResponse{
		Id:          updated.Id,
		Name:        updated.Name,
		Description: updated.Description,
	}
	return ctx.JSON(common.BaseResponse{Message: "Updated successfully", Data: res})
}

func (c *ProductCategoryController) DeleteCategory(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BaseResponse{Message: "Invalid request", Data: nil})
	}

	delErr := c.service.DeleteCategory(id)

	if delErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.BaseResponse{Message: "Deletion error", Data: nil})
	}

	return ctx.JSON(common.BaseResponse{Message: "Deleted successfully", Data: nil})
}

func (c *ProductCategoryController) GetCategoryById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	req, _ := strconv.Atoi(id)
	taken, takenErr := c.service.GetCategoryById(req)

	if takenErr != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(common.BaseResponse{Message: "Not found", Data: nil})
	}

	res := &contracts.CreateProductCategoryResponse{
		Id:          taken.Id,
		Name:        taken.Name,
		Description: taken.Description,
	}
	return ctx.JSON(common.BaseResponse{Message: "Finded successfully", Data: res})
}

func (c *ProductCategoryController) GetAll(ctx *fiber.Ctx) error {
	res := c.service.GetAllCategories()
	return ctx.JSON(common.BaseResponse{Message: "All categories", Data: res})
}
