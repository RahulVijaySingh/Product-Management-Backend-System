package handlers

import (
	"encoding/json"
	"net/http"
	"log"
	"product-management-system/cache"
	"product-management-system/database"

	"github.com/gin-gonic/gin"
)

// CreateProduct handles creating a new product
func CreateProduct(c *gin.Context) {
	var product struct {
		UserID             int      `json:"user_id"`
		ProductName        string   `json:"product_name"`
		ProductDescription string   `json:"product_description"`
		ProductImages      []string `json:"product_images"`
		ProductPrice       float64  `json:"product_price"`
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	_, err := database.DB.Exec(
		`INSERT INTO products (user_id, product_name, product_description, product_images, product_price)
		VALUES ($1, $2, $3, $4, $5)`,
		product.UserID, product.ProductName, product.ProductDescription, product.ProductImages, product.ProductPrice,
	)

	if err != nil {
		// Log the error for debugging
		log.Printf("Failed to create product: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}




func GetProductByID(c *gin.Context) {
	productID := c.Param("id")

	// Log the product ID being requested
	log.Printf("Fetching product with ID: %s", productID)

	// Check Redis cache first
	cachedData, err := cache.GetCache("product_" + productID)
	if err == nil && cachedData != "" {
		log.Println("Product found in cache")
		var product map[string]interface{}
		if err := json.Unmarshal([]byte(cachedData), &product); err == nil {
			c.JSON(http.StatusOK, product)
			return
		}
		log.Println("Failed to parse cached product data")
	}

	// Fetch from database if not in cache
	log.Println("Fetching product from database")
	row := database.DB.QueryRow(`
		SELECT id, user_id, product_name, product_description, product_images, compressed_product_images, product_price
		FROM products WHERE id = $1`, productID)

	var product struct {
		ID                 int      `json:"id"`
		UserID             int      `json:"user_id"`
		ProductName        string   `json:"product_name"`
		ProductDescription string   `json:"product_description"`
		ProductImages      []string `json:"product_images"`
		CompressedImages   []string `json:"compressed_product_images"`
		ProductPrice       float64  `json:"product_price"`
	}

	if err := row.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription, &product.ProductImages, &product.CompressedImages, &product.ProductPrice); err != nil {
		log.Printf("Error fetching product from database: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Cache the fetched product
	log.Println("Caching product data")
	productJSON, _ := json.Marshal(product)
	_ = cache.SetCache("product_"+productID, string(productJSON))

	// Return product details
	c.JSON(http.StatusOK, product)
}
// GetProducts handles retrieving all products for a user with optional filters
func GetProducts(c *gin.Context) {
	userID := c.Query("user_id")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")
	productName := c.Query("product_name")

	query := "SELECT id, product_name, product_price FROM products WHERE user_id=$1"
	args := []interface{}{userID}

	if minPrice != "" {
		query += " AND product_price >= $2"
		args = append(args, minPrice)
	}
	if maxPrice != "" {
		query += " AND product_price <= $3"
		args = append(args, maxPrice)
	}
	if productName != "" {
		query += " AND product_name ILIKE $4"
		args = append(args, "%"+productName+"%")
	}

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products"})
		return
	}
	defer rows.Close()

	var products []struct {
		ID          int     `json:"id"`
		ProductName string  `json:"product_name"`
		ProductPrice float64 `json:"product_price"`
	}

	for rows.Next() {
		var product struct {
			ID          int     `json:"id"`
			ProductName string  `json:"product_name"`
			ProductPrice float64 `json:"product_price"`
		}
		if err := rows.Scan(&product.ID, &product.ProductName, &product.ProductPrice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning products"})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}
