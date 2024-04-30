package ginUtil

import (
	"github.com/gin-gonic/gin"
	"go-clean-architecture-tutorial/core/entity"
	"go-clean-architecture-tutorial/core/usecase"
	"go-clean-architecture-tutorial/frameworks/database"
	"go-clean-architecture-tutorial/interface_adapters/controllers"
)

func Initialization(router *gin.Engine) {
	// Gin initialization code here
	router.GET("/products", getAllProducts)
	router.POST("/products", addProducts)
	router.Run("localhost:8080")
}

func getAllProducts(c *gin.Context) {
	// Get products logic here
	ProductController := controllers.GetAllProductsController{
		UseCase: usecase.GetAllProductsUsecase{
			ProductRepository: database.ProductArray.ProductRepository,
		},
	}

	result, err := ProductController.Handle(c.Request)
	if err != nil {
		c.JSON(200, result)
	}
}

func addProducts(c *gin.Context) {
	// Add products logic here
	product := entity.Product{}
	c.BindJSON(&product)

	database.ProductArray.Save(&product)

	c.JSON(200, product)
}
