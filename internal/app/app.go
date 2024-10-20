package app

import (
	"database/sql"
	"products/internal/config"
	"products/internal/controllers"
	"products/internal/repositories"
	"products/internal/services"
	"products/pkg/database/postgresql"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	conf := config.GetConfig()

	db := initDB(&conf.PostgresConfig)
	defer db.Close()

	app := fiber.New()

	categoryRepository := repositories.NewProductCategoryRepository(db)
	categoryService := services.NewProductCategoryService(categoryRepository)
	controllers.NewProductCategoryController(app, categoryService)

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	controllers.NewProductController(app, productService)
	app.Listen(conf.AppConfig.Host + ":" + conf.AppConfig.Port)
}

func initDB(conf *config.PostgresConfig) *sql.DB {
	db, err := postgresql.ConnectToDB(conf)

	if err != nil {
		panic(err)
	}

	return db
}
