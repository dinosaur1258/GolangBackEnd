package api

import (
	"net/http"
	"strconv"

	"github.com/dinosaur1258/GolangBackEnd/db"
	sqlc "github.com/dinosaur1258/GolangBackEnd/db/sqlc"
	"github.com/gin-gonic/gin"
)

// GetAllCompaniesHandler 獲取所有公司
func GetAllCompaniesHandler(c *gin.Context) {
	// 初始化資料庫連線
	db, err := db.DatabaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}
	defer db.Close()

	// 創建 Queries 實例
	queries := sqlc.New(db)
	ctx := c.Request.Context()

	// 使用 GetAll 函數獲取所有公司
	companies, err := queries.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching companies"})
		return
	}

	// 返回 JSON 響應
	c.JSON(http.StatusOK, companies)
}

// GetCompanyByIDHandler 根據 ID 獲取單個公司
func GetCompanyByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	// 初始化資料庫連線
	db, err := db.DatabaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}
	defer db.Close()

	// 創建 Queries 實例
	queries := sqlc.New(db)
	ctx := c.Request.Context()

	// 使用 GetByID 函數獲取公司
	company, err := queries.GetByID(ctx, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching company"})
		return
	}

	// 返回 JSON 響應
	c.JSON(http.StatusOK, company)
}

// CreateCompanyHandler 創建一個新的公司
func CreateCompanyHandler(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
		Info string `json:"info" binding:"required"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 初始化資料庫連線
	db, err := db.DatabaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}
	defer db.Close()

	// 創建 Queries 實例
	queries := sqlc.New(db)
	ctx := c.Request.Context()

	// 使用 Insert 函數創建新公司
	err = queries.Insert(ctx, sqlc.InsertParams{
		ID:   0, // 假設 ID 是自動生成的
		Name: input.Name,
		Info: input.Info,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting company"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "Company created"})
}

// UpdateCompanyHandler 更新指定 ID 的公司
func UpdateCompanyHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	var input struct {
		Name string `json:"name" binding:"required"`
		Info string `json:"info" binding:"required"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 初始化資料庫連線
	db, err := db.DatabaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}
	defer db.Close()

	// 創建 Queries 實例
	queries := sqlc.New(db)
	ctx := c.Request.Context()

	// 使用 Update 函數更新公司
	_, err = queries.Update(ctx, sqlc.UpdateParams{
		ID:   int32(id),
		Name: input.Name,
		Info: input.Info,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Company updated"})
}

// DeleteCompanyHandler 刪除指定 ID 的公司
func DeleteCompanyHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	// 初始化資料庫連線
	db, err := db.DatabaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}
	defer db.Close()

	// 創建 Queries 實例
	queries := sqlc.New(db)
	ctx := c.Request.Context()

	// 使用 Delete 函數刪除公司
	err = queries.Delete(ctx, int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting company"})
		return
	}

	c.JSON(http.StatusNoContent, nil) // No content for successful deletion
}
