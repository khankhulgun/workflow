package workflow

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khankhulgun/workflow/database/migrations"
	"github.com/khankhulgun/workflow/database/seeds"
	"github.com/lambda-platform/lambda/config"
)

func Set(app *fiber.App) {
	if config.Config.App.Migrate == "true" {
		migrations.Migrate()
	}
	if config.Config.App.Seed == "true" {
		seeds.Seed()
	}
}
