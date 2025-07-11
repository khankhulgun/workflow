package workflow

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khankhulgun/workflow/controllers"
	"github.com/khankhulgun/workflow/database/migrations"
	"github.com/khankhulgun/workflow/database/seeds"
	"github.com/lambda-platform/lambda/agent/agentMW"
	"github.com/lambda-platform/lambda/config"
)

func Set(app *fiber.App) {
	if config.Config.App.Migrate == "true" {
		migrations.Migrate()
	}
	if config.Config.App.Seed == "true" {
		seeds.Seed()
	}

	p := app.Group("/process")
	p.Get("/history/:id", controllers.History)
	p.Get("/history-with-user/:id", controllers.HistoryWithUser)
	p.Post("/notification", agentMW.IsLoggedIn(), controllers.SendNotification)

	w := app.Group("/workflow")
	w.Get("/steps", controllers.Steps)

}
