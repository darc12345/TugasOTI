package repository

import (
	"fmt"
	"log"
	"main/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h DBrepo) PostOrderRepo(c *gin.Context) error {
	db := h.db
	var orders model.OrdersDetail

	// Bind JSON to OrdersDetail struct
	err := c.ShouldBindJSON(&orders)
	if err != nil {
		log.Printf("Failed to bind JSON to OrdersDetail struct. Original Error: %s", err)
		return fmt.Errorf("error extracting order details")
	}

	// Generate a new OrderID
	orders.OrderID = uuid.NewString()

	// Insert order into Orders table
	_, err = db.Exec("INSERT INTO Orders(USERID, ORDERID) VALUES(?,?);", orders.UserID, orders.OrderID)
	if err != nil {
		log.Printf("Failed to insert order into Orders table. Order ID: %s, Original Error: %s", orders.OrderID, err)
		return fmt.Errorf("error saving the order")
	}

	// Insert products into OrderDetails table
	productLength := len(orders.PurchasedProducts)
	for i := 0; i < productLength; i++ {
		prod := orders.PurchasedProducts[i]
		_, err = db.Exec("INSERT INTO OrderDetails(ORDERID, PRODUCTID, QUANTITY, PRICE, TOTAL) VALUES(?,?,?,?,?);",
			orders.OrderID, prod.ProductID, prod.Quantity, prod.Price, prod.Total)
		if err != nil {
			log.Printf("Failed to insert product into OrderDetails table. Order ID: %s, Product ID: %s, Original Error: %s", orders.OrderID, prod.ProductID, err)
			return fmt.Errorf("error saving order details")
		}
	}

	return nil
}

func (h DBrepo) GetOrderByIDRepo(c *gin.Context) (model.OrdersDetail, error) {
	db := h.db
	var orders model.OrdersDetail

	// Extract OrderID from URL parameter
	orders.OrderID = c.Param("id")

	// Query user ID associated with the order
	err := db.QueryRow("SELECT USERID FROM Orders WHERE ORDERID=?", orders.OrderID).Scan(&orders.UserID)
	if err != nil {
		log.Printf("Failed to retrieve order by ID from Orders table. Order ID: %s, Original Error: %s", orders.OrderID, err)
		return model.OrdersDetail{}, fmt.Errorf("error fetching the order details")
	}

	// Query purchased products for the order
	rows, err := db.Query("SELECT PRODUCTID, QUANTITY, PRICE, TOTAL FROM OrderDetails WHERE ORDERID=?", orders.OrderID)
	if err != nil {
		log.Printf("Failed to retrieve order details from OrderDetails table. Order ID: %s, Original Error: %s", orders.OrderID, err)
		return model.OrdersDetail{}, fmt.Errorf("error fetching the order products")
	}
	defer rows.Close()

	// Scan each row and append to PurchasedProducts
	for rows.Next() {
		var ord struct {
			ProductID string
			Quantity  uint
			Price     uint
			Total     uint
		}
		err := rows.Scan(&ord.ProductID, &ord.Quantity, &ord.Price, &ord.Total)
		if err != nil {
			log.Printf("Failed to scan row in OrderDetails table. Order ID: %s, Original Error: %s", orders.OrderID, err)
			return model.OrdersDetail{}, fmt.Errorf("error processing order details")
		}
		orders.PurchasedProducts = append(orders.PurchasedProducts, ord)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration in OrderDetails table. Order ID: %s, Original Error: %s", orders.OrderID, err)
		return model.OrdersDetail{}, fmt.Errorf("error processing order details")
	}

	return orders, nil
}
