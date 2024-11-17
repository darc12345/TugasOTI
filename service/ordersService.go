package service

import (
	"main/model"

	"github.com/gin-gonic/gin"
)

func (h ServiceDB) PostOrderService(c *gin.Context) error {
	err := h.RepoDB.PostOrderRepo(c)
	return err
}
func (h ServiceDB) GetOrderByIDService(c *gin.Context) (model.OrdersDetail, error) {
	result, err := h.RepoDB.GetOrderByIDRepo(c)
	return result, err
}
