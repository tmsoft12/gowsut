package controller

import (
	"database/sql"
	"log"
	"strconv"
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

// DeleteTicket işlemini yapar: Belirli bir ticket'ı siler
func DeleteTicket(c *fiber.Ctx) error {
	// URL'den ID parametresini al
	ticketID := c.Params("id")

	// ID'yi integer'a çevir
	id, err := strconv.Atoi(ticketID)
	if err != nil {
		return c.Status(400).SendString("Geçersiz ticket ID")
	}

	// Veritabanında ticket'ı sil
	res, err := database.DB.Exec("DELETE FROM ticket WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Etkilenen satır sayısını kontrol et
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Eğer silinen satır yoksa, ticket bulunamamıştır
	if rowsAffected == 0 {
		return c.Status(404).SendString("Ticket bulunamadı")
	}

	// Başarılı mesajı döndür
	return c.SendString("Ticket başarıyla silindi")
}
