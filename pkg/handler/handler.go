package handler

import (
	"github.com/gin-gonic/gin"
	"storeApi/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/store")
	{
		auth.POST("/add", h.addProduct)
		auth.GET("/get", h.getProducts)
		auth.GET("/get/:id", h.getProductById)
		auth.POST("/buy/", h.buyProduct)
		auth.PUT("/update/:id", h.updateProduct)
		auth.PUT("/update/count", h.addCountProduct)
	}

	return router
}
