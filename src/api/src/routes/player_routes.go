package routes

import (
	"example.com/m/v2/src/db"
	"example.com/m/v2/src/models"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
)

func GetPlayers(c *fiber.Ctx) error {
	var players []models.Player
	iter := db.Session.Query("SELECT id, name, score FROM players").Iter()

	var id gocql.UUID
	var name string
	var score int
	for iter.Scan(&id, &name, &score) {
		players = append(players, models.Player{ID: id, Name: name, Score: score})
	}

	if err := iter.Close(); err != nil {
		return c.Status(500).SendString("Error receiving players")
	}
	return c.JSON(players)
}

func AddPlayer(c *fiber.Ctx) error {
	player := new(models.Player)
	if err := c.BodyParser(player); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	player.ID = gocql.TimeUUID()
	if err := db.Session.Query("INSERT INTO players (id, name, score) VALUES (?, ?, ?)",
		player.ID, player.Name, player.Score).Exec(); err != nil {
		return c.Status(500).SendString("Error adding player")
	}

	return c.JSON(player)
}

// Register player routes
func SetupPlayerRoutes(app *fiber.App) {
	app.Get("/players", GetPlayers)
	app.Post("/players", AddPlayer)
}
