package service

import (
	"main/model"

	"github.com/gin-gonic/gin"
)

func (h *ServiceDB) GetProductService(c *gin.Context) ([]model.Product, error) {
	return h.RepoDB.GetProductRepo(c)
}
func (h *ServiceDB) GetProductbyIDService(c *gin.Context) (model.Product, error) {
	return h.RepoDB.GetProductbyIDRepo(c)
}
