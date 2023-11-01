package main

 
import (
	"github.com/harkuur/go-restapi-gin/models"
	"github.com/harkuur/go-restapi-gin/controllers/productcontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()

	r.GET("api/products", productcontroller.Index)
	r.GET("api/products/:id", productcontroller.Show)
	r.POST("api/product", productcontroller.Create)
	r.PUT("api/products:id", productcontroller.Update)
	r.DELETE("api/products", productcontroller.Delete)

	r.Run();
}