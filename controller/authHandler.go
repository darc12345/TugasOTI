package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ControllerDB) PostLoginHandler(c *gin.Context) {
	err := h.serviceDB.PostLoginService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully logged in")

}
func (h *ControllerDB) PostRegisterHandler(c *gin.Context) {
	err := h.serviceDB.PostRegisterService(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully registered")

}
