package workflow

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khankhulgun/workflow/controllers"
	"github.com/khankhulgun/workflow/database/migrations"
	"github.com/khankhulgun/workflow/database/seeds"
	"github.com/khankhulgun/workflow/resolver"
	"github.com/lambda-platform/lambda/agent/agentMW"
	"github.com/lambda-platform/lambda/config"
)

func Set(app *fiber.App, empResolver ...resolver.EmployeeResolver) {
	if len(empResolver) > 0 {
		controllers.SetResolver(empResolver[0])
	}
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

	w.Get("/workflows/:category_id", agentMW.IsLoggedIn(), controllers.GetWorkflowsByCategory)

}
