package controller

import (
	"database/sql"
	"log"
	"tm/database"
	"tm/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// GetAndAddTicket işlemini yapar: en son ticket'ı alır ve yeni bir ticket ekler
func GetAndAddTicket(c *fiber.Ctx) error {
	// Query parametresini al
	typeParam := c.Query("type")

	var lastTicket models.Ticket

	// En son ticket'ı al
	row := database.DB.QueryRow("SELECT id, ticket, type FROM ticket ORDER BY id DESC LIMIT 1")
	err := row.Scan(&lastTicket.ID, &lastTicket.Ticket, &lastTicket.Type)
	if err != nil && err != sql.ErrNoRows {
		return c.Status(500).SendString(err.Error())
	}

	// Yeni ticket verisini ekle (son ticket değerine 1 ekleyerek)
	newTicketNumber := lastTicket.Ticket + 1
	// Test amaçlı statik bir değer belirle
	staticType := typeParam // Kullanıcıdan alınan type değeri
	_, err = database.DB.Exec("INSERT INTO ticket (ticket, type) VALUES (?, ?)", newTicketNumber, staticType)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Yeni ticket'ı almak için sorgu yap
	row = database.DB.QueryRow("SELECT id, ticket, type FROM ticket WHERE ticket = ?", newTicketNumber)
	var newTicket models.Ticket
	err = row.Scan(&newTicket.ID, &newTicket.Ticket, &newTicket.Type)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	log.Println(newTicket.Ticket, newTicket.Type)
	// JSON formatında yeni ticket'ı döndür
	return c.JSON(fiber.Map{
		"newTicket": newTicket,
	})
}

// GetTicketsByType işlemini yapar: Her bir type için ayrı ayrı diziler döndürür
func GetTicketsByType(c *fiber.Ctx) error {
	// Sorgular
	queries := map[string]string{
		"wheelchair": "SELECT id, ticket, type FROM ticket WHERE type = 'wheelchair' ORDER BY ticket ASC",
		"child":      "SELECT id, ticket, type FROM ticket WHERE type = 'child' ORDER BY ticket ASC",
		"person":     "SELECT id, ticket, type FROM ticket WHERE type = 'person' ORDER BY ticket ASC",
	}

	// Sonuçları tutacak diziler
	results := make(map[string][]models.Ticket)

	// Her bir type için sorgu yap
	for key, query := range queries {
		rows, err := database.DB.Query(query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer rows.Close()

		var tickets []models.Ticket
		for rows.Next() {
			var ticket models.Ticket
			if err := rows.Scan(&ticket.ID, &ticket.Ticket, &ticket.Type); err != nil {
				return c.Status(500).SendString(err.Error())
			}
			tickets = append(tickets, ticket)
		}

		if err := rows.Err(); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		results[key] = tickets
	}

	// JSON formatında tüm ticket'ları döndür
	return c.JSON(fiber.Map{
		"wheelchair": results["wheelchair"],
		"child":      results["child"],
		"person":     results["person"],
	})
}

// HandleWebSocket WebSocket bağlantısını yönetir
func HandleWebSocket(c *websocket.Conn) {
	queries := map[string]string{
		"wheelchair": "SELECT id, ticket, type FROM ticket WHERE type = 'wheelchair' ORDER BY ticket ASC",
		"child":      "SELECT id, ticket, type FROM ticket WHERE type = 'child' ORDER BY ticket ASC",
		"person":     "SELECT id, ticket, type FROM ticket WHERE type = 'person' ORDER BY ticket ASC",
	}

	results := make(map[string][]models.Ticket)

	for key, query := range queries {
		rows, err := database.DB.Query(query)
		if err != nil {
			log.Println(err)
			return
		}
		defer rows.Close()

		var tickets []models.Ticket
		for rows.Next() {
			var ticket models.Ticket
			if err := rows.Scan(&ticket.ID, &ticket.Ticket, &ticket.Type); err != nil {
				log.Println(err)
				return
			}
			tickets = append(tickets, ticket)
		}

		if err := rows.Err(); err != nil {
			log.Println(err)
			return
		}

		results[key] = tickets
	}

	data := fiber.Map{
		"wheelchair": results["wheelchair"],
		"child":      results["child"],
		"person":     results["person"],
	}

	err := c.WriteJSON(data)
	if err != nil {
		log.Println("Error sending WebSocket data:", err)
	}
}
func DeleteTicket(c *fiber.Ctx) error {
	id := c.Params("id")

	_, err := database.DB.Exec("DELETE FROM ticket WHERE id=$1", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(200).JSON("Oki")
}
