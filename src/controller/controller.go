
package controller

import (
	"net/http"
	"tracker/entity"
	"tracker/service"
	"github.com/gin-gonic/gin"
)

// Controller handles transaction-related HTTP requests
type Controller struct {
	TransactionService service.TransactionService
}

// Config holds controller dependencies
type Config struct {
	R                 *gin.Engine
	TransactionService service.TransactionService
}

// NewController initializes routes and returns a Controller instance
func NewController(c *Config) *Controller {
	controller := &Controller{
		TransactionService: c.TransactionService,
	}
	apiRoutes := c.R.Group("/api")
	{
		apiRoutes.GET("/txn/:id", controller.FindAllTransactionsById)
		apiRoutes.GET("/txn", controller.FindAllTransactions)
		apiRoutes.POST("/txn/add", controller.AddTransaction)
		apiRoutes.POST("/txn/edit", controller.EditTransaction)
		apiRoutes.DELETE("/txn/delete", controller.DeleteTransaction)
	}
	return controller
}

// FindAllTransactions retrieves all transactions
func (c *Controller) FindAllTransactionsById(ctx *gin.Context) {
	transactions, err := c.TransactionService.FindById(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Transaksi": transactions})
}

// FindAllTransactions retrieves all transactions
func (c *Controller) FindAllTransactions(ctx *gin.Context) {
	transactions, err := c.TransactionService.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Transaksi": transactions})
}

// AddTransaction creates a new transaction
func (c *Controller) AddTransaction(ctx *gin.Context) {
	var txn entity.Transaction
	if err := ctx.ShouldBindJSON(&txn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.TransactionService.Add(ctx.Request.Context(), txn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"Pesan": "berhasil ditambahkan!"})
}

// EditTransaction updates an existing transaction
func (c *Controller) EditTransaction(ctx *gin.Context) {
	var txn entity.Transaction
	if err := ctx.ShouldBindJSON(&txn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.TransactionService.Edit(ctx.Request.Context(), txn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"Pesan": "berhasil diubah!"})
}

// DeleteTransaction removes a transaction by ID
func (c *Controller) DeleteTransaction(ctx *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.TransactionService.Delete(ctx.Request.Context(), request.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"Pesan": "berhasil dihapus!"})
}
