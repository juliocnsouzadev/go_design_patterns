package cs

import (
	"fmt"
	"math/rand"
)

func CounterStrike(numberPlayers int) {
	weapons := []string{"AK-47", "Maverick", "Gut Knife", "Desert Eagle"}
	var bots []*Bot
	for i := 0; i < numberPlayers; i++ {
		rnd := rand.Intn(len(weapons))
		bot := NewRandomPlayer(&weapons[rnd])
		bots = append(bots, bot)
	}

	fmt.Println("** Terrorists **")
	for _, player := range bots {
		if terrorist != player.pType {
			continue
		}
		fmt.Println(player.mission())
	}

	fmt.Println("\n** Counter Terrorists **")
	for _, player := range bots {
		if counterTerrorist != player.pType {
			continue
		}
		fmt.Println(player.mission())
	}
}
