package api

import (
	"github.com/gin-gonic/gin"
)

// StartServer 設置和啟動 Gin 路由
func StartServer() {
	router := gin.Default()

	// 定義路由
	router.GET("/companies", GetAllCompaniesHandler)
	router.GET("/companies/:id", GetCompanyByIDHandler)
	router.POST("/companies", CreateCompanyHandler)
	router.PUT("/companies/:id", UpdateCompanyHandler)
	router.DELETE("/companies/:id", DeleteCompanyHandler)

	// 啟動伺服器
	router.Run(":8080")
}
