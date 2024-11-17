package repository

import (
	"fmt"
	"log"
	"main/model"

	"github.com/gin-gonic/gin"
)

func (h *DBrepo) GetProductRepo(c *gin.Context) ([]model.Product, error) {
	db := h.db

	// Query products from the database
	rows, err := db.Query("SELECT PRODUCTID, PRODUCTDESC, PRODUCTNAME FROM Products")
	if err != nil {
		log.Printf("Failed to retrieve products from database. Original Error: %s", err)
		return nil, fmt.Errorf("error fetching product list")
	}
	defer rows.Close()

	var products []model.Product

	// Iterate through the result set
	for rows.Next() {
		var prod model.Product
		err := rows.Scan(&prod.Productid, &prod.Productdesc, &prod.Productname)
		if err != nil {
			log.Printf("Failed to scan product data. Original Error: %s", err)
			return nil, fmt.Errorf("error processing product data")
		}
		products = append(products, prod)
	}

	// Check for errors during rows iteration
	if err = rows.Err(); err != nil {
		log.Printf("Error during rows iteration for products. Original Error: %s", err)
		return nil, fmt.Errorf("error processing product data")
	}

	return products, nil
}

func (h *DBrepo) GetProductbyIDRepo(c *gin.Context) (model.Product, error) {
	db := h.db
	id := c.Param("id")
	var prod model.Product

	// Query product by ID
	err := db.QueryRow("SELECT PRODUCTID, PRODUCTDESC, PRODUCTNAME FROM Products WHERE PRODUCTID = ?", id).
		Scan(&prod.Productid, &prod.Productdesc, &prod.Productname)
	if err != nil {
		log.Printf("Failed to retrieve product by ID from database. Product ID: %s, Original Error: %s", id, err)
		return model.Product{}, fmt.Errorf("error fetching product details")
	}

	return prod, nil
}
