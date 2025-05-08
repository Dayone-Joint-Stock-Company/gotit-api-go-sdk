package main

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"gotit_api_go_sdk"
)

// Config holds application configuration
type Config struct {
	BaseURL          string
	GIAuthorization  string
	Port             string
	Environment      string
}

// App holds application dependencies
type App struct {
	config *Config
	client *resty.Client
	sdk    *gotit_api_go_sdk.APIClient
}

const (
	defaultPort      = "5004"
	baseURL         = "api-biz-stg.gotit.vn"
)

// newConfig creates a new Config instance
func newConfig() (*Config, error) {
	giAuth := os.Getenv("X_GI_AUTHORIZATION")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	env := os.Getenv("GOTIT_API_ENV")
	if env == "" {
		env = "development"
	}

	return &Config{
		BaseURL:         baseURL,
		GIAuthorization: giAuth,
		Port:            port,
		Environment:     env,
	}, nil
}

// newApp creates a new App instance
func newApp(config *Config) *App {
	// Initialize REST client
	client := resty.New()
	client.SetBaseURL("https://" + config.BaseURL)
	client.SetHeader("X-GI-Authorization", config.GIAuthorization)

	// Initialize SDK client
	sdkConfig := gotit_api_go_sdk.NewConfiguration()
	sdkConfig.Host = config.BaseURL
	sdkConfig.Scheme = "https"
	sdkConfig.AddDefaultHeader("X-GI-Authorization", config.GIAuthorization)
	sdk := gotit_api_go_sdk.NewAPIClient(sdkConfig)

	return &App{
		config: config,
		client: client,
		sdk:    sdk,
	}
}

func loadPrivateKeyFromString(privateKeyStr string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	return priv, nil
}

// Move signData to be a standalone function
func signData(message string, privateKey *rsa.PrivateKey) (string, error) {
	hashed := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", fmt.Errorf("failed to sign data: %v", err)
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func generateTransactionRef() string {
	randomBytes := make([]byte, 4)
	if _, err := rand.Read(randomBytes); err != nil {
		log.Printf("Warning: Failed to generate random bytes: %v", err)
		// Fallback to timestamp-based reference
		return fmt.Sprintf("000742_%s", time.Now().Format("20060102150405"))
	}
	
	timestamp := time.Now().Format("20060102")
	randomHex := hex.EncodeToString(randomBytes)
	
	return fmt.Sprintf("000742_%s-%s", timestamp, randomHex)
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	
	if _, err := rand.Read(result); err != nil {
		log.Printf("Warning: Failed to generate random string: %v", err)
		// Fallback to timestamp-based string
		return time.Now().Format("20060102150405")
	}
	
	for i := range result {
		result[i] = charset[int(result[i])%len(charset)]
	}
	
	return string(result)
}

func setupRouter(app *App) *gin.Engine {
	if app.config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	
	r := gin.Default()

	// Middleware
	r.Use(corsMiddleware())
	r.Use(loggingMiddleware())

	// API Routes
	api := r.Group("/")
	{
		// Brands routes
		brands := api.Group("/brands")
		{
			brands.GET("", app.getBrands)
			brands.GET("/:id", app.getBrandDetail)
			brands.GET("/brand_product", app.getBrandProducts)
		}

		// Categories routes
		categories := api.Group("/categories")
		{
			categories.GET("", app.getCategories)
			categories.GET("/category_product", app.getCategoryProducts)
		}

		// Products routes
		products := api.Group("/products")
		{
			products.GET("", app.getProducts)
			products.GET("/:id", app.getProductDetail)
			products.GET("/:id/stores", app.getProductStores)
		}

		// Vouchers routes
		vouchers := api.Group("/vouchers")
		{
			vouchers.GET("/e", app.generateVouchersE)
			vouchers.GET("/v", app.generateVouchersV)
			vouchers.GET("/g", app.generateVouchersG)
			vouchers.GET("/send_sms", app.sendVoucherSMS)
			vouchers.GET("/send_email", app.sendVoucherEmail)
			vouchers.GET("/send_zns", app.sendVoucherZNS)
			vouchers.GET("/check_zns", app.checkStatusZNS)
			vouchers.GET("/check_status_voucher", app.checkStatusVoucher)
		}
	}

	return r
}

// Middleware functions
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func loggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		
		c.Next()
		
		latency := time.Since(start)
		status := c.Writer.Status()
		
		log.Printf("[%d] %s %s - %v", status, c.Request.Method, path, latency)
	}
}

func main() {
	// Initialize configuration
	config, err := newConfig()
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	// Initialize application
	app := newApp(config)

	// Setup router
	router := setupRouter(app)

	// Start server
	addr := ":" + config.Port
	log.Printf("Server starting on %s...", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// API Handlers
func (app *App) getBrands(c *gin.Context) {
	brands, _, err := app.sdk.BrandsAPI.GetListOfBrands(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		Execute()
	if err != nil {
		log.Printf("Error in getBrands: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, brands)
}

func (app *App) getBrandDetail(c *gin.Context) {
	brandID := c.Param("id")
	id, err := strconv.ParseInt(brandID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand ID"})
		return
	}
	
	brand, _, err := app.sdk.BrandsAPI.GetDetailOfBrand(context.Background(), id).
		XGIAuthorization(app.config.GIAuthorization).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, brand)
}

func (app *App) getBrandProducts(c *gin.Context) {
	products, _, err := app.sdk.BrandsAPI.BrandsByProducts(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, products)
}

func (app *App) getCategories(c *gin.Context) {
	categories, _, err := app.sdk.CategoriesAPI.GetListOfCategories(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, categories)
}

func (app *App) getCategoryProducts(c *gin.Context) {
	products, _, err := app.sdk.CategoriesAPI.GetCategoryByProduct(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, products)
}

func (app *App) getProducts(c *gin.Context) {
	products, _, err := app.sdk.ProductsAPI.GetListOfProducts(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		Page(1).
		PageSize(200).
		MinPrice(1000).
		MaxPrice(10000000).
		IsExcludeStoreListInfo(false).
		StoreListPage(1).
		StoreListPageSize(50).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, products)
}

func (app *App) getProductDetail(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, _, err := app.sdk.ProductsAPI.GetProductDetail(context.Background(), id).
		XGIAuthorization(app.config.GIAuthorization).
		IsExcludeStoreListInfo(false).
		StoreListPage(1).
		StoreListPageSize(50).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, product)
}

func (app *App) getProductStores(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	
	stores, _, err := app.sdk.ProductsAPI.GetStoresOfProduct(context.Background(), id).
		XGIAuthorization(app.config.GIAuthorization).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, stores)
}

func (app *App) generateVouchersE(c *gin.Context) {
	privateKeyStr := os.Getenv("PRIVATE_KEY_STR")
	if privateKeyStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PRIVATE_KEY_STR environment variable is required for E vouchers"})
		return
	}

	privateKey, err := loadPrivateKeyFromString(privateKeyStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to load private key: %v", err)})
		return
	}

	transactionRef := generateTransactionRef()
	message := fmt.Sprintf("%s|%s", app.config.GIAuthorization, transactionRef)
	
	signature, err := signData(message, privateKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	voucher, _, err := app.sdk.VoucherAPI.CreateVoucherLinkE(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		REQUESTCREATEVOUCHERLINKE(gotit_api_go_sdk.REQUESTCREATEVOUCHERLINKE{
			OrderName:        "Got It Promotion",
			TransactionRefId: transactionRef,
			Amount:          100000,
			ExpiryDate:      "2025-11-15",
			ProductId:       gotit_api_go_sdk.PtrInt64(1541),
			UseOtp:          1,
			OtpType:         gotit_api_go_sdk.PtrInt64(1),
			Password:        gotit_api_go_sdk.PtrString("123456"),
			ReceiverName:    gotit_api_go_sdk.PtrString("Quang Huynh"),
			Phone:           gotit_api_go_sdk.PtrString("0394162063"),
			Email:           gotit_api_go_sdk.PtrString("quang.huynh@gotit.vn"),
			Signature:       signature,
		}).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, voucher)
}

func (app *App) generateVouchersV(c *gin.Context) {
	transactionRef := generateTransactionRef()
	orderName := generateRandomString(10)

	voucher, _, err := app.sdk.VoucherAPI.CreateVoucherLinkV(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		REQUESTCREATEVOUCHERLINKV(gotit_api_go_sdk.REQUESTCREATEVOUCHERLINKV{
			ProductId:             1541,
			ProductPriceId:        2991,
			Quantity:             1,
			OrderName:            orderName,
			ExpiryDate:           "2025-11-15",
			TransactionRefId:     transactionRef,
			UseOtp:               gotit_api_go_sdk.PtrInt64(1),
			OtpType:              gotit_api_go_sdk.PtrInt64(1),
			Password:             gotit_api_go_sdk.PtrString("123456"),
			ReceiverName:         gotit_api_go_sdk.PtrString("Quang Huynh"),
			Phone:                gotit_api_go_sdk.PtrString("0394162063"),
			Email:                gotit_api_go_sdk.PtrString("quang.huynh@gotit.vn"),
			ActiveDate:           gotit_api_go_sdk.PtrString(""),
			IsConvertToCoverLink: gotit_api_go_sdk.PtrInt64(1),
		}).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, voucher)
}

func (app *App) generateVouchersG(c *gin.Context) {
	transactionRef := generateTransactionRef()
	orderName := generateRandomString(10)

	voucher, _, err := app.sdk.VoucherAPI.CreateVoucherLinkG(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		REQUESTCREATEVOUCHERLINKG(gotit_api_go_sdk.REQUESTCREATEVOUCHERLINKG{
			ProductList: []gotit_api_go_sdk.PRODUCTDEFAULTLINKG{
				{
					ProductId:      1541,
					ProductPriceId: 2991,
					Quantity:       1,
				},
			},
			OrderName:        orderName,
			ExpiryDate:       "2025-11-15",
			TransactionRefId: transactionRef,
			UseOtp:          gotit_api_go_sdk.PtrInt64(1),
			OtpType:         gotit_api_go_sdk.PtrInt64(1),
			Password:        gotit_api_go_sdk.PtrString("123456"),
			ReceiverName:    gotit_api_go_sdk.PtrString("Quang Huynh"),
			Phone:           gotit_api_go_sdk.PtrString("0394162063"),
			Email:           gotit_api_go_sdk.PtrString("quang.huynh@gotit.vn"),
		}).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, voucher)
}

func (app *App) sendVoucherSMS(c *gin.Context) {
	response, _, err := app.sdk.VoucherSendMethodAPI.SendVoucherBySMS(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		REQUESTSENDVOUCHERBYSMS(gotit_api_go_sdk.REQUESTSENDVOUCHERBYSMS{
			VoucherLinkCode: "gLsZaFRN",
			PhoneNo:         "0394162063",
			ReceiverNm:      gotit_api_go_sdk.PtrString("Got It"),
			SenderNm:        gotit_api_go_sdk.PtrString("Got It"),
		}).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, response)
}

func (app *App) sendVoucherEmail(c *gin.Context) {
	response, _, err := app.sdk.VoucherSendMethodAPI.SendVoucherByEmail(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		REQUESTSENDVOUCHERBYEMAIL(gotit_api_go_sdk.REQUESTSENDVOUCHERBYEMAIL{
			VoucherLinkCode: "gLsZaFRN",
			Email:           "quang.huynh@gotit.vn",
			ReceiverNm:      gotit_api_go_sdk.PtrString("Got It"),
			SenderNm:        gotit_api_go_sdk.PtrString("Got It"),
			Message:         gotit_api_go_sdk.PtrString("Have a good day"),
		}).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, response)
}

func (app *App) sendVoucherZNS(c *gin.Context) {
	transactionID := generateRandomString(15)

	response, _, err := app.sdk.VoucherSendMethodAPI.SendVoucherByZNS(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		REQUESTSENDVOUCHERBYZNS(gotit_api_go_sdk.REQUESTSENDVOUCHERBYZNS{
			PhoneNo:        "0394162063",
			ReceiverNm:     gotit_api_go_sdk.PtrString("Got It"),
			VoucherSerials: []string{"E2V2RML6F52"},
			TransactionId:  transactionID,
		}).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, response)
}

func (app *App) checkStatusZNS(c *gin.Context) {
	response, _, err := app.sdk.VoucherSendMethodAPI.CheckStatusZNS(context.Background()).
		XGIAuthorization(app.config.GIAuthorization).
		REQUESTCHECKSTATUSZNS(gotit_api_go_sdk.REQUESTCHECKSTATUSZNS{
			TransactionId: gotit_api_go_sdk.PtrString("aNGqaAljTaLc5B3"),
		}).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, response)
}

func (app *App) checkStatusVoucher(c *gin.Context) {
	response, _, err := app.sdk.VoucherStatusAPI.CheckVoucher(context.Background(), "1111_1aa11").
		XGIAuthorization(app.config.GIAuthorization).
		IsExcludeStoreListInfo(false).
		StoreListPage(1).
		StoreListPageSize(50).
		Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, response)
} 