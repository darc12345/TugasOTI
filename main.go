package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/controller"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	file, err := os.OpenFile("log/myLOG.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	cfg := mysql.Config{
		User:   "bob",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "projectdb",
	}
	// postgresURI := "mysql://root:password@(127.0.0.1:3306)/projectdb"
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Print("Koneksi ke database gagal")
		fmt.Print(err)
		os.Exit(0)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("ping db gagal")
		fmt.Print(err)
		os.Exit(0)
	}
	if err != nil {
		fmt.Println("Error dropping Products table:", err)
	}

	r := gin.Default()
	server := &http.Server{
		Addr:              ":8080",
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		Handler:           r,
	}
	ctrlHandler := controller.NewControllerDB(db)

	r.GET("/api/v1/products", ctrlHandler.GetProductHandler)
	r.GET("/api/v1/products/:id", ctrlHandler.GetProductbyIDHandler)
	r.POST("/api/v1/orders/buy", ctrlHandler.PostOrderController)
	r.GET("api/v1/orders/:id", ctrlHandler.GetOrderByIDcontroller)
	r.POST("/admin/product", ctrlHandler.PostProductHandler)
	r.DELETE("/admin/products/:id", ctrlHandler.DeleteProductHandler)
	r.PUT("/admin/products/:id", ctrlHandler.PutProductHandler)
	r.POST("/api/v1/login", ctrlHandler.PostLoginHandler)
	r.POST("/api/v1/register", ctrlHandler.PostRegisterHandler)
	log.Fatal(server.ListenAndServe())
}
