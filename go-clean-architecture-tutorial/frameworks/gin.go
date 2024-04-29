package frameworks

import (
	"go-clean-architecture-tutorial/core/entity"

	"github.com/gin-gonic/gin"
)

var productArray ProductDatabaseArray = ProductDatabaseArray{
	products: []entity.Product{
		{ID: 1, Name: "Product 1", Price: 100},
		{ID: 2, Name: "Product 2", Price: 200},
		{ID: 3, Name: "Product 3", Price: 300},
	},
}

func InitializeGin() {
	// Gin initialization code here
	router := gin.Default()
	router.GET("/products", getProducts)
	router.POST("/products", addProducts)
	router.Run("localhost:8080")
}

func getProducts(c *gin.Context) {
	// Get products logic here
	products, _ := productArray.GetAll()

	c.JSON(200, products)
}

func addProducts(c *gin.Context) {
	// Add products logic here
	product := entity.Product{}
	c.BindJSON(&product)

	productArray.Save(&product)

	c.JSON(200, product)
}
