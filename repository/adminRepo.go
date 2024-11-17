package repository

import (
	"fmt"
	"log"
	"main/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *DBrepo) PostProductRepo(c *gin.Context) error {
	db := h.db
	prod := model.Product{}

	// Extract product data from the request
	err := c.BindJSON(&prod)
	if err != nil {
		log.Printf("Failed to bind JSON to product struct. Original Error: %s", err)
		return fmt.Errorf("error extracting the product data")
	}

	// Assign a new product ID
	prod.Productid = uuid.NewString()

	// Insert the product into the database
	_, err = db.Exec("INSERT INTO Products(PRODUCTID, PRODUCTDESC, PRODUCTNAME) VALUES(?,?,?);", prod.Productid, prod.Productdesc, prod.Productname)
	if err != nil {
		log.Printf("Failed to insert product into the database. Product ID: %s. Original Error: %s", prod.Productid, err)
		return fmt.Errorf("failed to insert the product into the database")
	}

	return nil
}

func (h *DBrepo) DeleteProductRepo(c *gin.Context) error {
	db := h.db
	id := c.Param("id")

	// Delete the product from the database
	_, err := db.Exec("DELETE FROM Products WHERE PRODUCTID=?;", id)
	if err != nil {
		log.Printf("Failed to delete product from the database. Product ID: %s. Original Error: %s", id, err)
		return fmt.Errorf("failed to delete the product from the database")
	}

	return nil
}

func (h *DBrepo) PutProductRepo(c *gin.Context) error {
	prod := model.Product{}

	// Extract product data from the request
	err := c.BindJSON(&prod)
	if err != nil {
		log.Printf("Failed to bind JSON to product struct. Original Error: %s", err)
		return fmt.Errorf("error extracting the product data")
	}

	// Use product ID from the URL parameter
	prod.Productid = c.Param("id")
	fmt.Println(prod.Productid)

	// Update the product in the database
	db := h.db
	_, err = db.Exec("UPDATE Products SET PRODUCTDESC=?, PRODUCTNAME=? WHERE PRODUCTID=?;", prod.Productdesc, prod.Productname, prod.Productid)
	if err != nil {
		log.Printf("Failed to update product in the database. Product ID: %s. Original Error: %s", prod.Productid, err)
		return fmt.Errorf("failed to update the product in the database")
	}

	return nil
}
