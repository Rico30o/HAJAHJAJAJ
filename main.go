package main

import (
	"fmt"
	"instapay/db"
	"instapay/middleware"
	"instapay/routes"
	"log"

	models "instapay/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	//----------------------------------- Database ---------------------------------
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	err = db.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Auto Migrate
	db.DB.AutoMigrate(&models.User{})

	// Set up Fiber app
	app := fiber.New()

	//------------------------------------logs ------------------------------------
	// app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                           // means all user can access the API or you can just specify the user that can use the system (example: facebook.com) only
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", // header that we only accept``
	}),
		logger.New(logger.Config{ //Modified Logs
			Format:     "${cyan}${time} ${white}| ${green}${status} ${white}| ${ip} | ${host} | ${method} | ${magenta}${path} ${white}| ${red}${latency} ${white}\n",
			TimeFormat: "01/02/2006 3:04 PM",
			// TimeZone:   envRouting.PostgresTimeZone,
		}))

	//=================logfile ====================
	// err = os.MkdirAll("./logs", os.ModePerm)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Generate a dynamic filename for the log file, based on the current date.
	// logFilename := fmt.Sprintf("./logs/Alert %s.log", time.Now().Format("2006-01-02"))

	// // Open the log file in append mode.
	// logfile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer logfile.Close()

	// // Set the output of the log to the file.
	// log.SetOutput(logfile)

	//----------------------------------------------------------------------------

	// Register user routes not found 404
	routes.SetupUserRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			if e, ok := err.(*fiber.Error); ok && e.Code == fiber.StatusNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"Errors": fiber.Map{
						"Error": []fiber.Map{
							{
								"Source":      "Gateway",
								"ReasonCode":  "NOT_FOUND",
								"Description": "URL not found",
								"Recoverable": false,
								"Details":     nil,
							},
						},
					},
				})
			}
			return err
		}
		return nil
	})

	//---------------------------- localhost ----------------------------
	// Start server
	// err = app.Listen(":8080")
	// if err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }

	// address := fmt.Sprintf("192.168.0.122:%s", middleware.GetEnv("PORT"))

	// if err := app.Listen(address); err != nil {
	// 	log.Fatal(err.Error())
	// }

	if middleware.GetEnv("SSL") == "enabled" {
		log.Fatal(app.ListenTLS(
			fmt.Sprintf(":%s", middleware.GetEnv("PORT")),
			middleware.GetEnv("SSL_CERTIFICATE"),
			middleware.GetEnv("SSL_KEY"),
		))
	} else {
		err := app.Listen(fmt.Sprintf("192.168.0.123:%s", middleware.GetEnv("PORT")))
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	//-----------------------------------------------------------------
}
