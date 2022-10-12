package main

import (
  "github.com/gin-gonic/gin"
  "myretail.com/retailer_service/models"
  "myretail.com/retailer_service/controllers"
)

func main() {
  r := gin.Default()
  r.SetTrustedProxies(nil)

  r.POST("/product", controllers.CreateProduct)
  r.PATCH("/product/:id", controllers.UpdateProduct)
  r.GET("/products", controllers.ListProducts)
  r.GET("/product/:id", controllers.ShowProduct)

  models.InitDB()

  // TODO: Handle graceful shutdown
  r.Run(":8080")
}
