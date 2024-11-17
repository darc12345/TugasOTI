package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h ControllerDB) PostOrderController(c *gin.Context) {
	err := h.serviceDB.PostOrderService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())

		return
	}
	c.JSON(http.StatusOK, "Order successfully posted")
}
func (h ControllerDB) GetOrderByIDcontroller(c *gin.Context) {
	result, err := h.serviceDB.GetOrderByIDService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}
