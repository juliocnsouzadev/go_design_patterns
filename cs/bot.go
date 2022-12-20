package cs

import (
	"fmt"
	"math/rand"
)

const (
	terrorist        = "terrorist"
	counterTerrorist = "counter_terrorist"
)

var (
	plantBomb   = "PLANT A BOMB"
	diffuseBomb = "DIFFUSE BOMB"
	tasks       = map[string]*string{
		terrorist:        &plantBomb,
		counterTerrorist: &diffuseBomb,
	}
)

type Bot struct {
	pType  string
	task   *string
	weapon *string
}

func (p *Bot) mission() string {
	return fmt.Sprintf("%s carrying a %s has the task to %s", p.pType, *p.weapon, *p.task)
}

func NewRandomPlayer(weapon *string) *Bot {
	i := rand.Intn(100)
	if i%2 == 0 {
		return newTerrorist(weapon)
	}
	return newCounterTerrorist(weapon)
}

func newTerrorist(weapon *string) *Bot {
	return &Bot{pType: terrorist, task: tasks[terrorist], weapon: weapon}
}

func newCounterTerrorist(weapon *string) *Bot {
	return &Bot{pType: counterTerrorist, task: tasks[counterTerrorist], weapon: weapon}
}
