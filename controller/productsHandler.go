package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ControllerDB) GetProductHandler(c *gin.Context) {
	prod, err := h.serviceDB.GetProductService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, prod)
}
func (h *ControllerDB) GetProductbyIDHandler(c *gin.Context) {
	prod, err := h.serviceDB.GetProductbyIDService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, prod)
}
