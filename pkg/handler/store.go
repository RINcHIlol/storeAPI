package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"storeApi"
	"strconv"
)

func (h *Handler) addProduct(c *gin.Context) {
	var input storeApi.Product

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	isAdd, err := h.services.Store.AddNewProduct(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isAdd": isAdd,
	})
}

func (h *Handler) getProducts(c *gin.Context) {
	products, err := h.services.Store.GetProducts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"products": products,
	})
}

func (h *Handler) getProductById(c *gin.Context) {

	productId, err := strconv.Atoi(c.Param("id"))

	product, err := h.services.Store.GetProductById(productId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"product": product,
	})
}

func (h *Handler) buyProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	isBought, err := h.services.Store.BuyProduct(productId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isBought": isBought,
	})
}

func (h *Handler) updateProduct(c *gin.Context) {
	var input storeApi.Product

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	productId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	isUpdate, err := h.services.Store.UpdateProductById(productId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isUpdate": isUpdate,
	})
}
