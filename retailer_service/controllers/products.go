package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "myretail.com/retailer_service/models"
)

// TODO: refactor some of the boiler-plate code like
// fetching by Id, validing the input to a func in resp packages

type createProductInput struct {
  Name string `json:"product_name" binding:"required"`
  Price float64 `json:"price" binding:"required"`
  Quantity uint `json:"quantity" binding:"required"`
}

type updateProductInput struct {
  Price float64 `json:"price" binding:"required"`
  Quantity uint `json:"quantity" binding:"required"`
}

// POST /product
func CreateProduct(c *gin.Context) {
  var input createProductInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  product := models.Product{Name: input.Name, Price: input.Price, Quantity: input.Quantity}
  result := models.DB.Create(&product)
  if result.Error != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": result.Error})
  }
  c.JSON(http.StatusOK, product)
}

// PATCH /product/:id
func UpdateProduct(c *gin.Context) {
  var product models.Product
  if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  var input updateProductInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  result := models.DB.Model(&product).Updates(&models.Product{Price: input.Price, Quantity: input.Quantity})
  if result.Error != nil {
    c.JSON(http.StatusUnprocessableEntity, gin.H{"error": result.Error})
  }
  c.JSON(http.StatusOK, product)
}

// GET /products
func ListProducts(c *gin.Context) {
  var prods []models.Product
  models.DB.Where("quantity > ?", 0).Find(&prods)
  c.JSON(http.StatusOK, gin.H{"products": prods})
}

// GET /product/:id
func ShowProduct(c *gin.Context) {
  var product models.Product

  if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, product)
}
