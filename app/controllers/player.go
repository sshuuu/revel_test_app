package controllers

import (
	"math/rand"
	"strconv"
	"test_app/app/models"
	"test_app/app/routes"
	"time"

	"github.com/revel/revel"
)

// Player type
type Player struct {
	*revel.Controller
}

// Index take all players
func (c Player) Index() revel.Result {
	players := []models.Player{}
	Gdb.Find(&players)
	players[0].Age()
	c.ViewArgs["players"] = players
	return c.Render()
}

// Show show one player detail
func (c Player) Show() revel.Result {
	playerIDStr := c.Params.Get("id")
	playerID, _ := strconv.Atoi(playerIDStr)
	player := models.Player{}
	Gdb.Where("id = ?", playerID).Find(&player)
	return c.Render(player)
}

// Create create one player
func (c Player) Create() revel.Result {
	createPlayerRamdomly()
	return c.Redirect(routes.Player.Index())
}

// Delete delete one player
func (c Player) Delete() revel.Result {
	playerIDStr := c.Params.Get("id")
	playerID, _ := strconv.Atoi(playerIDStr)
	deletePlayer(playerID)
	return c.Redirect(routes.Player.Index())
}

func createPlayerRamdomly() {
	name := getNameRamdomly()
	age := getAgeRandomly()
	player := models.Player{Name: name, Birthday: time.Now().AddDate(-age, 0, 0)}
	Gdb.NewRecord(player)
	Gdb.Create(&player)
	Gdb.NewRecord(player)
}

func getNameRamdomly() string {
	names := []string{"ichiro", "jiro", "saburo", "shiro", "goro", "jon"}
	return choice(names)
}

func getAgeRandomly() int {
	return rand.Intn(100)
}

func choice(s []string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(s))
	return s[i]
}

func deletePlayer(playerID int) {
	player := models.Player{}
	Gdb.Where("id = ?", playerID).Delete(&player)
}
