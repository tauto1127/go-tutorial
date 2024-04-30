package frameworks

import (
	"go-clean-architecture-tutorial/frameworks/ginUtil"

	"github.com/gin-gonic/gin"
)

func InitializeGin() {
	// Gin initialization code here
	router := gin.Default()

	ginUtil.Initialization(router)
}
