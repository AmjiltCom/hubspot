package HubSpot

import (
	"github.com/AmjiltCom/hubspot/routes"
	"github.com/gofiber/fiber/v2"
)

func SetHubSpot(e *fiber.App) {
	routes.HubSpotApi(e)
}
