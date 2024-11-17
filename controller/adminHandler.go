package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ControllerDB) PostProductHandler(c *gin.Context) {
	err := h.serviceDB.PostProductService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, "Data inserted")
}
func (h *ControllerDB) PutProductHandler(c *gin.Context) {
	err := h.serviceDB.PutProductService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, "Data Updated")
}
func (h *ControllerDB) DeleteProductHandler(c *gin.Context) {
	err := h.serviceDB.DeleteProductService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, "Data Deleted")
}
