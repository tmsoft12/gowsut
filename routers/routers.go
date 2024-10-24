package routers

import (
	"tm/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func InitRouters(app *fiber.App) {
	app.Get("/", controller.GetAndAddTicket)
	app.Get("/ticket", controller.GetTicketsByType)
	app.Get("/ws", websocket.New(controller.HandleWebSocket))
	app.Delete("/del/:id", controller.DeleteTicket)
}
