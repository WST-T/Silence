package game

import (
	"fmt"
)

// describeCurrentLocation shows details about the player's current location
func (g *Game) describeCurrentLocation() {
	location := g.Locations[g.PlayerLocation]

	fmt.Println("\n===", g.PlayerLocation, "===")
	fmt.Println(location.Description)

	// List characters in the room
	if len(location.Characters) == 0 {
		fmt.Println("There is no one else in this room.")
	} else {
		fmt.Println("\nIn this room:")
		for _, name := range location.Characters {
			if g.Characters[name].IsAlive {
				fmt.Println("  -", name)
			}
		}
	}
}

// showCommands lists all available player commands
func (g *Game) showCommands() {
	fmt.Println("\nWhat would you like to do?")
	fmt.Println("  [look] - Look around the room again")
	fmt.Println("  [go <location>] - Go to another room")
	fmt.Println("  [talk <person>] - Talk to someone in the room")
	fmt.Println("  [accuse <person>] - Accuse someone of being the killer")
	fmt.Println("  [clues] - Review collected clues and information")
	fmt.Println("  [locations] - See all locations in the manor")
	fmt.Println("  [status] - Check the status of all guests")
	fmt.Println("  [help] - Show this help menu")
	fmt.Println("  [quit] - Exit the game")
}

// showLocations lists all locations in the manor
func (g *Game) showLocations() {
	fmt.Println("\nLocations in the manor:")
	for name := range g.Locations {
		fmt.Println("  -", name)
	}
}

// showStatus displays the status of all guests (alive or dead)
func (g *Game) showStatus() {
	fmt.Println("\nGuest Status:")
	for name, character := range g.Characters {
		status := "Alive"
		if !character.IsAlive {
			status = "Dead"
		}
		fmt.Printf("  - %s: %s (Location: %s)\n", name, status, character.Location)
	}

	// Show murder count
	fmt.Printf("\nMurder Events: %d\n", g.MurderCount)
	fmt.Printf("People Murdered: %d\n", g.MurderCount*2)
	fmt.Printf("People Remaining Alive: %d\n", len(g.Characters)-g.MurderCount*2)
}

// showIntroduction displays the game's introduction text
func showIntroduction() {
	fmt.Println("======================================")
	fmt.Println("       		SILENCE                ")
	fmt.Println("======================================")
	fmt.Println("\nIt's been 15 years since graduation. Your old classmates have gathered at")
	fmt.Println("Blackwood Manor for a reunion dinner. The mood is tense as conversations")
	fmt.Println("turn to Sarah Williams, a classmate who died under mysterious circumstances")
	fmt.Println("shortly after graduation.")
	fmt.Println("\nAs a detective, you notice the strange undercurrents among the guests...")
	fmt.Println("Then the lights go out, and the killing begins.")
	fmt.Println("\nCan you find the killer before everyone is dead?")
}
