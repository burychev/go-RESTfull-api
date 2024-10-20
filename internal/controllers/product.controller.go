package controllers

import (
	"fmt"
	"products/internal/common"
	"products/internal/contracts"
	"products/internal/models"
	"products/internal/services"
	"products/pkg/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	service services.IProductService
}

func NewProductController(api fiber.Router, s services.IProductService) *ProductController {
	controller := &ProductController{service: s}

	api.Post("/products", controller.CreateProduct)
	api.Put("/products/:id", controller.UpdateProduct)
	api.Delete("/products/:id", controller.DeleteProduct)
	api.Get("/products/:id", controller.GetProduct)
	api.Get("/products", controller.GetAll)
	return controller
}

func (c *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	var req contracts.CreateProductRequest

	if err := utils.ReadRequestFromBody(ctx, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BaseResponse{Message: "Invalid request", Data: nil})
	}

	entity := models.Product{Name: req.Name, Description: req.Description, Price: req.Price, CategoryId: req.CategoryId}
	createdEntity, creationErr := c.service.CreateProduct(entity)

	if creationErr != nil {
		fmt.Println(creationErr)
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.BaseResponse{Message: "Creation error", Data: nil})
	}

	res := &contracts.CreateProductResponse{
		Id:          createdEntity.Id,
		Name:        createdEntity.Name,
		Description: createdEntity.Description,
		Price:       createdEntity.Price,
		CategoryId:  createdEntity.CategoryId,
	}
	return ctx.JSON(common.BaseResponse{Message: "Created successfully", Data: res})
}

func (c *ProductController) UpdateProduct(ctx *fiber.Ctx) error {
	var req contracts.UpdateProductRequest

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BaseResponse{Message: "Invalid request", Data: nil})
	}

	if err := utils.ReadRequestFromBody(ctx, &req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BaseResponse{Message: "Invalid request", Data: nil})
	}

	entity := models.Product{Id: id, Name: req.Name, Description: req.Description, Price: req.Price, CategoryId: req.CategoryId}

	updated, updationErr := c.service.UpdateProduct(entity)

	if updationErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.BaseResponse{Message: "Updation error", Data: nil})
	}

	res := &contracts.CreateProductResponse{
		Id:          updated.Id,
		Name:        updated.Name,
		Description: updated.Description,
		Price:       updated.Price,
		CategoryId:  updated.CategoryId,
	}
	return ctx.JSON(common.BaseResponse{Message: "Updated successfully", Data: res})
}

func (c *ProductController) DeleteProduct(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BaseResponse{Message: "Invalid request", Data: nil})
	}

	delErr := c.service.DeleteProduct(id)

	if delErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.BaseResponse{Message: "Deletion error", Data: nil})
	}

	return ctx.JSON(common.BaseResponse{Message: "Deleted successfully", Data: nil})
}

func (c *ProductController) GetProduct(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.BaseResponse{Message: "Invalid request", Data: nil})
	}

	taken, takenErr := c.service.GetProductById(id)

	if takenErr != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(common.BaseResponse{Message: "Not found", Data: nil})
	}

	res := &contracts.CreateProductResponse{
		Id:          taken.Id,
		Name:        taken.Name,
		Description: taken.Description,
		Price:       taken.Price,
		CategoryId:  taken.CategoryId,
	}
	return ctx.JSON(common.BaseResponse{Message: "Finded product", Data: res})
}

func (c *ProductController) GetAll(ctx *fiber.Ctx) error {

	res := c.service.GetAllProducts()

	return ctx.JSON(common.BaseResponse{Message: "All products", Data: res})
}
