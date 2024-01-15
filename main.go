package main

import (
	"log"
	"os"
	"tracking-app/app/config"
	"tracking-app/app/controllers"
	"tracking-app/app/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"

)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init DB
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Buat objek dbConfig dengan nilai-nilai dari environment variables
	dbConfig := config.DatabaseConfig{
		Host:     dbHost,
		Port:     dbPort,
		Username: dbUser,
		Password: dbPassword,
		DBName:   dbName,
	}

	dsn := dbConfig.GetDSN()
	config.InitDatabase(dsn)

	// Get the value of the BASE_URL environment variable
	baseUrl := os.Getenv("BASE_URL")

	// Initialize standard Go html template engine
	engine := html.New("./app/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//Middleware logger
	app.Use(logger.New())

	//Session Declaration
	store := session.New()

	// Serve Static
	app.Static("/", "./static")

	// Wire Controllers
	dashboardController := controllers.NewDashboardController(baseUrl, store)
	shipmentController := controllers.NewShipmentController(baseUrl, store)
	driverController := controllers.NewDriverController(baseUrl, store)
	scannerController := controllers.NewScannerController(baseUrl, store)
	authController := controllers.NewAuthController(baseUrl, store)
	customerController := controllers.NewCustomerController(baseUrl, store)

	// Middleware auth
	authMiddleware := middlewares.AuthMiddleware(store)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("welcome", fiber.Map{}, "layouts/main")
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		message := c.Query("message")
		return c.Render("login", fiber.Map{"message": message})
	})
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", fiber.Map{})
	})

	// Routes with Controller
	// Dashboard
	app.Get("/dashboard", authMiddleware, dashboardController.Index)
	// Shipment
	app.Get("/shipment", authMiddleware, shipmentController.Index)
	app.Post("/create-barcode", authMiddleware, shipmentController.Create)
	app.Get("/search-po", authMiddleware, shipmentController.Search)
	app.Get("/shipments", authMiddleware, shipmentController.GetAllShipments)
	// Driver
	app.Get("/driver", authMiddleware, driverController.Index)
	app.Post("/drivers", authMiddleware, driverController.Create)
	app.Post("/drivers-update", authMiddleware, driverController.Update)
	app.Post("/drivers-delete", authMiddleware, driverController.Delete)
	//CUSTOMER
	app.Get("/customer", authMiddleware, customerController.Index)
	app.Post("/customers", authMiddleware, customerController.Create)
	app.Post("/customers-update", authMiddleware, customerController.Update)
	app.Post("/customers-delete", authMiddleware, customerController.Delete)
	// Scanner
	app.Get("/scanner", authMiddleware, scannerController.Index)
	app.Get("/detail", authMiddleware, scannerController.Detail)
	app.Post("/update-status-shipment", authMiddleware, scannerController.UpdateShipment)

	// Auth
	app.Post("/handle-register", authController.HandleRegister)
	app.Post("/handle-login", authController.HandleLogin)
	app.Get("/logout", authController.Logout)

	// Start server
	// err = app.Listen("localhost:" + "8000")
	// jika ingin diakses dari ip address dari hp
	err = app.Listen(":8000")
	if err != nil {
		return
	}

}
