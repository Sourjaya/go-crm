package main

import (
	"github.com/Sourjaya/go-crm/database"
	"github.com/Sourjaya/go-crm/lead"
	l "github.com/Sourjaya/go-crm/logger"
	"github.com/Sourjaya/go-crm/router"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func init() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		l.Error("Failed to connect to database")
	}
	l.Info("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	l.Info("Database Migrated")
}
func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
